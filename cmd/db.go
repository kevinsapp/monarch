package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// pg
	_ "github.com/lib/pq"
)

// Global ...
var db *sql.DB

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(createDBCmd)
	dbCmd.AddCommand(dropDBCmd)
	dbCmd.AddCommand(pingDBCmd)
	dbCmd.AddCommand(resetDBCmd)
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use:               "db",
	Short:             `Provides subcommands for working with databases.`,
	PersistentPreRunE: openDB,
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

// openDB opens a connection pool on the global DB object for connecting to the
// database specified in the viper config.
func openDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Open a new connection pool and assign it to the global "db" variable,
	// don't shadow it with a local variable.
	var err error
	db, err = sql.Open("postgres", srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: openDB: %s\n", err)
	}

	return err
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

	// Open a DB connection pool local to this method, don't assign a new
	// connection pool to the global variable.
	srv.dbName = ""                            // dbName should be blank before getting DSN.
	db, err := sql.Open("postgres", srv.dsn()) // shadow db; not global db
	if err != nil {
		log.Fatalf("ERROR: createDB: %s\n", err)
	}
	defer db.Close()

	// Execute query to create database.
	start := time.Now()
	_, err = db.Exec(query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q created. Server replied in %s.\n", database.Name(), duration)

	return err
}

// dropDB drops a database with the name specificed by the "database" attribute
// in the viper config.
func dropDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Configure database object for SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.DropDBTmpl)
	if err != nil {
		log.Printf("ERROR: dropDB: %s\n", err)
		return err
	}

	// Open a DB connection pool local to this method, don't assign a new
	// connection pool to the global variable.
	srv.dbName = "" // dbName should be blank before getting dsn.
	db, err := sql.Open("postgres", srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: dropDB: %s\n", err)
	}
	defer db.Close()

	// Execute query to drop database.
	start := time.Now()
	_, err = db.Exec(query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q dropped. Server replied in %s.\n", database.Name(), duration)

	return err
}

// ping connects to the database to verify that the server is accessible.
func pingDB(cmd *cobra.Command, args []string) error {
	start := time.Now()
	err := db.Ping()
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database connection OK. Server replied in %s.\n", duration)

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
