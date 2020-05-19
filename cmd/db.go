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
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use:               "db",
	PersistentPreRunE: openDB,
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

// pingCmd ...
var pingDBCmd = &cobra.Command{
	Use:  "ping",
	RunE: pingDB,
}

// openDB ...
func openDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()
	dsn := srv.getDSN()

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("ERROR: openDB: %s\n", err)
		return err
	}

	return err
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
		log.Printf("ERROR: createDB: %s\n", err)
		return err
	}

	// Open a DB connection pool
	srv.dbName = "" // dbName should be blank before getting dsn.
	db, err := sql.Open("postgres", srv.getDSN())
	if err != nil {
		log.Printf("ERROR: createDB: %s\n", err)
		return err
	}
	defer db.Close()

	// Execute query to create database.
	start := time.Now()
	_, err = db.Exec(query)
	duration := time.Since(start)
	if err != nil {
		log.Printf("ERROR: createDB: %s\n", err)
		return err
	}

	fmt.Printf("Database %q created. Server replied in %s.\n", database.Name(), duration)

	return err
}

// dropDB ...
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

	// Open a DB connection pool
	srv.dbName = "" // dbName should be blank before getting dsn.
	db, err := sql.Open("postgres", srv.getDSN())
	if err != nil {
		log.Printf("ERROR: dropDB: %s\n", err)
		return err
	}
	defer db.Close()

	// Execute query to drop database.
	start := time.Now()
	_, err = db.Exec(query)
	duration := time.Since(start)
	if err != nil {
		log.Printf("ERROR: dropDB: %s\n", err)
		return err
	}

	fmt.Printf("Database %q dropped. Server replied in %s.\n", database.Name(), duration)

	return err
}

// Ping the database to verify that the server is accessible.
// If ping fails, log and return an error.
func pingDB(cmd *cobra.Command, args []string) error {
	now := time.Now()
	err := db.Ping()
	since := time.Since(now)
	if err != nil {
		log.Printf("ERROR: pingDB: %s\n", err)
		return err
	}

	fmt.Printf("Database connection OK. Server replied in %s.\n", since)

	return err
}

// resetDB
func resetDB(cmd *cobra.Command, args []string) error {
	// Drop DB
	err := dropDB(cmd, args)
	if err != nil {
		log.Printf("ERROR: resetDB: %s\n", err)
		return err
	}

	// Create DB
	err = createDB(cmd, args)
	if err != nil {
		log.Printf("ERROR: resetDB: %s\n", err)
	}

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

func (s *dbServer) getDSN() string {
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

func (s *dbServer) initFromConfig() {
	// Read in config.
	s.host = viper.GetString("development.host")
	s.port = viper.GetInt("development.port")
	s.user = viper.GetString("development.user")
	s.password = viper.GetString("development.password")
	s.dbName = viper.GetString("development.database")
	s.sslMode = viper.GetString("development.sslmode")
}
