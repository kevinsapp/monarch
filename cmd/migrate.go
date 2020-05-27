package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
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

// migrateDB establishes a connection to the database and executes "up"
// migrations
func migrateDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Connect to the database server.
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, srv.dsn())
	if err != nil {
		return err
	}
	defer pool.Close()

	// Timestamp command start.
	start := time.Now()

	// Up migrate the schema.
	err = upMigrateSchema(ctx, pool)
	if err != nil {
		return err
	}

	// Timestamp command end.
	duration := time.Since(start)

	fmt.Printf("Database %q migrated. Command completed in %s.\n", srv.dbName, duration)

	return nil
}

// upMigrateSchema executes up migrates later than the last version in the
// schema_versions table.
func upMigrateSchema(ctx context.Context, pool *pgxpool.Pool) error {
	// Create the schema_migrations table if it does not exist.
	err := createSchemaVersionsTable(ctx, pool)
	if err != nil {
		return err
	}

	// Fetch latest schema version from schema_versions table.
	ver, err := fetchSchemaVersion(ctx, pool)
	if err != nil {
		return err
	}

	// Stage the "up" migrations later than schema version
	ms, err := migration.LoadAllLaterThan(ver, migrationsDir)
	if err != nil {
		return err
	}

	// Execute migrations.
	err = execUpMigrations(ctx, pool, ms)

	return err
}

// createSchemaVersionsTable creates a schema_versions table if it does not
// already exist.
func createSchemaVersionsTable(ctx context.Context, pool *pgxpool.Pool) error {
	sql := `CREATE TABLE IF NOT EXISTS schema_versions (
		version bigint NOT NULL,
		created_at timestamp(6) without time zone NOT NULL,
		CONSTRAINT schema_migrations_pkey PRIMARY KEY (version)
	);`

	_, err := pool.Exec(ctx, sql)
	if err != nil {
		return err
	}

	return err
}

// fetchSchemaVersion fetches latest schema version from schema_versions table.
func fetchSchemaVersion(ctx context.Context, pool *pgxpool.Pool) (int64, error) {
	r := pool.QueryRow(ctx, "SELECT max(version) FROM schema_versions;")

	var v pgtype.Int8
	err := r.Scan(&v)
	if err != nil {
		return v.Int, err
	}

	fmt.Printf("Current schema version is: %d\n", v.Int)

	return v.Int, err
}

// execUpMigrations executes all "up" migrations in contained in a
// []migration.Migration in a single database transaction. If any migration
// fails, then the transaction is rolled back and no migrations are committed.
func execUpMigrations(ctx context.Context, pool *pgxpool.Pool, ms []migration.Migration) error {
	// Begin a database transaction.
	tx, err := pool.Begin(ctx)
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
		stmt := "INSERT INTO schema_versions (version, created_at) VALUES ($1, now());"
		_, err = tx.Exec(ctx, stmt, m.Version())
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
