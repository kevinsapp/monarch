package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v4"
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

	// Migrate the schema.
	err = migrateSchema(conn)
	if err != nil {
		return err
	}

	// Timestamp command end.
	duration := time.Since(start)

	fmt.Printf("Database %q migrated. Command completed in %s.\n", srv.dbName, duration)

	return nil
}

func migrateSchema(conn *pgx.Conn) error {
	// sql := `SET search_path TO public;`
	// _, err = tx.Exec(ctx, sql)
	// if err != nil {
	// 	return err
	// }

	type migration struct {
		sql     string
		version int64
	}

	// Create the schema_migrations table if it does not exist.
	ctx := context.Background()
	sql := `CREATE TABLE IF NOT EXISTS schema_versions (
		version bigint NOT NULL,
		created_at timestamp(6) without time zone NOT NULL,
		CONSTRAINT schema_migrations_pkey PRIMARY KEY (version)
	);`
	_, err := conn.Exec(ctx, sql)
	if err != nil {
		return err
	}

	// Determine latest schema version based on schema_versions table.
	var latestSchemaVersion pgtype.Int8
	row := conn.QueryRow(ctx, "SELECT max(version) FROM schema_versions")
	if err != nil {
		return err
	}
	err = row.Scan(&latestSchemaVersion)
	if err != nil {
		return err
	}

	// Get the list of files in the migrations directory.
	mfs, err := ioutil.ReadDir("migrations")
	if err != nil {
		return err
	}

	// Select only the migration files with:
	// a) a suffix of "up.sql", and
	// b) a version greater than the latest version in schema_migration
	migrations := make([]migration, 0)
	var m migration
	for _, f := range mfs {
		n := f.Name()
		if strings.HasSuffix(n, "up.sql") {
			// Extract migration version from filename.
			fnParts := strings.Split(n, "_")
			ver, err := strconv.ParseInt(fnParts[0], 10, 64)
			if err != nil {
				return err
			}

			if ver > latestSchemaVersion.Int {
				sql, err := sqlt.FileAsString("migrations/" + n)
				if err != nil {
					return err
				}
				m.version = ver
				m.sql = sql
				migrations = append(migrations, m)
			}
		}
	}

	// Begin a database transaction.
	tx, err := conn.Begin(ctx)
	defer tx.Rollback(ctx)

	// Read in and execute the SQL from each migration file.
	for _, m := range migrations {
		// Execute SQL statement from migration.
		_, err = tx.Exec(ctx, m.sql)
		if err != nil {
			return err
		}

		// Insert new migration version into schema_version table
		_, err = tx.Exec(ctx, "INSERT INTO schema_versions (version, created_at) VALUES ($1, now());", m.version)
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
