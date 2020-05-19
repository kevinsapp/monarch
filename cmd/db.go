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
var dsn string

// dbServer
type dbServer struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

func (s *dbServer) getDSN() string {
	// If dbName is not set, format a data source name without a dbname and return it.
	if s.dbName == "" {
		format := "host=%s port=%d user=%s password=%s sslmode=%s"
		dsn := fmt.Sprintf(format, s.host, s.port, s.user, s.password, s.sslMode)
		return dsn
	}

	// Format at data source name.
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dsn := fmt.Sprintf(format, s.host, s.port, s.user, s.password, s.dbName, s.sslMode)

	return dsn
}

func (s *dbServer) initFromConfig() {
	// Read in config.
	s.host = viper.GetString("development.host")
	s.port = viper.GetInt("development.port")
	s.user = viper.GetString("development.user")
	s.password = viper.GetString("development.password")
	s.dbName = viper.GetString("development.database")
	s.sslMode = viper.GetString("development.sslmode")
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(createDBCmd)
	dbCmd.AddCommand(dropDBCmd)
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use:              "db",
	PersistentPreRun: openDB,
}

// createCmd ...
var createDBCmd = &cobra.Command{
	Use:  "create",
	RunE: createDB,
}

// dropCmd ...
var dropDBCmd = &cobra.Command{
	Use:  "drop",
	RunE: dropDB,
}

// openDB ...
func openDB(cmd *cobra.Command, args []string) {
	var srv dbServer
	srv.initFromConfig()
	dsn = srv.getDSN()

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
}

// createDB ...
func createDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Configure database object params for SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)
	database.SetOwner(srv.user)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.CreateDBTmpl)
	if err != nil {
		log.Fatalf("Error: createDB: %s\n", err)
	}

	// Open a DB connection pool
	srv.dbName = "" // dbName should be blank before getting dsn.
	db, err := sql.Open("postgres", srv.getDSN())
	if err != nil {
		log.Fatalf("ERROR: createDB: %s\n", err)
	}
	defer db.Close()

	// Execute query to create database.
	start := time.Now()
	_, err = db.Exec(query)
	duration := time.Since(start)
	if err != nil {
		log.Fatalf("Error: createDB: %s\n", err)
	}

	fmt.Printf("Database %q created. Server replied in %s.\n", database.Name(), duration)

	return nil
}

// dropDB ...
func dropDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Configure database object params for SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.DropDBTmpl)
	if err != nil {
		log.Fatalf("Error: createDB: %s\n", err)
	}

	// Open a DB connection pool
	srv.dbName = "" // dbName should be blank before getting dsn.
	db, err := sql.Open("postgres", srv.getDSN())
	if err != nil {
		log.Fatalf("ERROR: dropDB: %s\n", err)
	}
	defer db.Close()

	// Execute query to create database.
	start := time.Now()
	_, err = db.Exec(query)
	duration := time.Since(start)
	if err != nil {
		log.Fatalf("Error: createDB: %s\n", err)
	}

	fmt.Printf("Database %q dropped. Server replied in %s.\n", database.Name(), duration)

	return nil
}
