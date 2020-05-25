package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/kevinsapp/monarch/pkg/fileutil"
	"github.com/kevinsapp/monarch/pkg/migration"
	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(createDBCmd)
	dbCmd.AddCommand(dropDBCmd)
	dbCmd.AddCommand(migrateDBCmd)
	dbCmd.AddCommand(pingDBCmd)
	dbCmd.AddCommand(resetDBCmd)
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: `Provides subcommands for working with databases.`,
}

// createCmd ...
var createDBCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a database with the name specificed by the "database" attribute in the config file.`,
	RunE:  createDB,
}

// dropCmd ...
var dropDBCmd = &cobra.Command{
	Use:   "drop",
	Short: `Drop a database with the name specificed by the "database" attribute in the config file.`,
	RunE:  dropDB,
}

// migrateCmd ...
var migrateDBCmd = &cobra.Command{
	Use:   "migrate",
	Short: `Migrate a database.`,
	RunE:  migrateDB,
}

// pingDBCmd ...
var pingDBCmd = &cobra.Command{
	Use:   "ping",
	Short: `Verifies that Monarch can connect to the database specificed in the config file.`,
	RunE:  pingDB,
}

// resetDBCmd ...
var resetDBCmd = &cobra.Command{
	Use:   "reset",
	Short: `First drops and then creates a database with the name specificed by the "database" attribute in the config file.`,
	RunE:  resetDB,
}

// createDB creates a database with the name specificed by the "database"
// attribute in the viper config.
func createDB(cmd *cobra.Command, args []string) error {
	// Initialize a dbServer object.
	var srv dbServer
	srv.initFromConfig()

	// Configure a data object to apply to a SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)
	database.SetOwner(srv.user)

	// Process the SQL template.
	query, err := sqlt.ProcessTmpl(&database, sqlt.CreateDBTmpl)
	if err != nil {
		log.Fatalf("ERROR: createDB: %s\n", err)
	}

	// Connect to the database server.
	srv.dbName = "" // dbName should be blank before getting DSN.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: createDB: %s\n", err)
	}
	defer conn.Close(ctx)

	// Execute query to create database.
	start := time.Now()
	_, err = conn.Exec(ctx, query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q created. Command completed in %s.\n", database.Name(), duration)

	return err
}

// dropDB drops a database with the name specificed by the "database" attribute
// in the viper config.
func dropDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Configure a data object to apply to a SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.DropDBTmpl)
	if err != nil {
		log.Printf("ERROR: dropDB: %s\n", err)
		return err
	}

	// Connect to the database server.
	srv.dbName = "" // dbName should be blank before getting dsn.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: dropDB: %s\n", err)
	}
	defer conn.Close(ctx)

	// Execute query to drop database.
	start := time.Now()
	_, err = conn.Exec(ctx, query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q dropped. Command completed in %s.\n", database.Name(), duration)

	return err
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
	err = upMigrateSchema(conn)
	if err != nil {
		return err
	}

	// Timestamp command end.
	duration := time.Since(start)

	fmt.Printf("Database %q migrated. Command completed in %s.\n", srv.dbName, duration)

	return nil
}

func upMigrateSchema(conn *pgx.Conn) error {
	// Create the schema_migrations table if it does not exist.
	ctx := context.Background()
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
	ms, err := stageUpMigrationsLaterThan(ver)
	if err != nil {
		return err
	}

	// Execute migrations.
	err = execUpMigrations(ctx, conn, ms)

	return err
}

// ping connects to the database to verify that the server is accessible.
func pingDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Connect to the database server.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	start := time.Now()
	err = conn.Ping(ctx)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database connection OK. Command completed in %s.\n", duration)

	return err
}

// resetDB drops and creates a database, i.e. reset.
func resetDB(cmd *cobra.Command, args []string) error {
	err := dropDB(cmd, args)
	if err != nil {
		return err
	}

	err = createDB(cmd, args)

	return err
}

// dbServer
type dbServer struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

// dsn returns a Data Source Name (dsn) string based on the dbServer attributes.
func (s *dbServer) dsn() string {
	// If dbName is not set, format a data source name without a dbname and return it.
	if s.dbName == "" {
		format := "host=%s port=%d user=%s password=%s sslmode=%s"
		dsn := fmt.Sprintf(format, s.host, s.port, s.user, s.password, s.sslMode)
		return dsn
	}

	// Format a data source name.
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dsn := fmt.Sprintf(format, s.host, s.port, s.user, s.password, s.dbName, s.sslMode)

	return dsn
}

// intiFromConfig initalizes a dbServer{} from the viper config.
func (s *dbServer) initFromConfig() {
	// Read in config.
	s.host = viper.GetString("development.host")
	s.port = viper.GetInt("development.port")
	s.user = viper.GetString("development.user")
	s.password = viper.GetString("development.password")
	s.dbName = viper.GetString("development.database")
	s.sslMode = viper.GetString("development.sslmode")
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

func stageUpMigrationsLaterThan(version int64) ([]migration.Migration, error) {
	migrations := make([]migration.Migration, 0)

	// Get the list of files in the migrations directory.
	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		return migrations, err
	}

	var m migration.Migration
	for _, f := range files {
		n := f.Name()
		v, err := fileutil.ExtractVersionFromFile(n)
		if err != nil {
			return migrations, err
		}

		// Select only the migration files with:
		// a) a suffix of "up.sql", and
		// b) a version greater than schemaVersion
		if v > version && strings.HasSuffix(n, "up.sql") {
			m.SetFromFile("migrations/" + n)
			migrations = append(migrations, m)
			fmt.Printf("Staged %q migration version: %d\n", "up", m.Version())
		}
	}

	return migrations, err
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
		_, err = tx.Exec(ctx, m.SQL())
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
