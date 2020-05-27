package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/kevinsapp/monarch/pkg/migration"
	"github.com/spf13/cobra"
)

func init() {
	dbCmd.AddCommand(migrateDBCmd)
}

// migrateCmd ...
var migrateDBCmd = &cobra.Command{
	Use:   "migrate",
	Short: `Migrate a database.`,
	RunE:  migrateDB,
}

func migrateDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Connect to the database server.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	// Timestamp command start.
	start := time.Now()

	// Up migrate the schema.
	err = upMigrateSchema(ctx, conn)
	if err != nil {
		return err
	}

	// Timestamp command end.
	duration := time.Since(start)

	fmt.Printf("Database %q migrated. Command completed in %s.\n", srv.dbName, duration)

	return nil
}

func upMigrateSchema(ctx context.Context, conn *pgx.Conn) error {
	// Create the schema_migrations table if it does not exist.
	err := createSchemaMigrationsTable(ctx, conn)
	if err != nil {
		return err
	}

	// Fetch latest schema version from schema_versions table.
	ver, err := fetchSchemaVersion(ctx, conn)
	if err != nil {
		return err
	}

	// Stage the "up" migrations later than schema version
	ms, err := migration.LoadAllLaterThan(ver, migrationsDir)
	if err != nil {
		return err
	}

	// Execute migrations.
	err = execUpMigrations(ctx, conn, ms)

	return err
}

func createSchemaMigrationsTable(ctx context.Context, conn *pgx.Conn) error {
	sql := `CREATE TABLE IF NOT EXISTS schema_versions (
		version bigint NOT NULL,
		created_at timestamp(6) without time zone NOT NULL,
		CONSTRAINT schema_migrations_pkey PRIMARY KEY (version)
	);`

	_, err := conn.Exec(ctx, sql)
	if err != nil {
		return err
	}

	return err
}

// fetchSchemaVersion fetches latest schema version from schema_versions table.
func fetchSchemaVersion(ctx context.Context, conn *pgx.Conn) (int64, error) {
	r := conn.QueryRow(ctx, "SELECT max(version) FROM schema_versions")

	var v pgtype.Int8
	err := r.Scan(&v)
	if err != nil {
		return v.Int, err
	}

	fmt.Printf("Current schema version is: %d\n", v.Int)

	return v.Int, err
}

func execUpMigrations(ctx context.Context, conn *pgx.Conn, ms []migration.Migration) error {
	// Begin a database transaction.
	tx, err := conn.Begin(ctx)
	defer tx.Rollback(ctx)

	// sql := `SET search_path TO public;`
	// _, err = tx.Exec(ctx, sql)
	// if err != nil {
	// 	return err
	// }

	// Migrate schema.
	for _, m := range ms {
		// Execute SQL statement from migration.
		_, err = tx.Exec(ctx, m.UpSQL())
		if err != nil {
			return err
		}

		// Insert migration version into schema_version table
		_, err = tx.Exec(ctx, "INSERT INTO schema_versions (version, created_at) VALUES ($1, now());", m.Version())
		if err != nil {
			return err
		}
	}

	// All statements must have executed ok, so commit the tranaction.
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return err
}
