package cmd

import (
	"fmt"
	"log"
	"time"

	sqlt "github.com/kevinsapp/monarch/pkg/sql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// pg
	_ "github.com/lib/pq"
)

// Global ...
var db *sqlt.DB
var dsn string

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(createDBCmd)
	dbCmd.AddCommand(dropDBCmd)

	host := viper.Get("development.host")
	port := viper.Get("development.port")
	user := viper.Get("development.user")
	pw := viper.Get("development.password")
	dbname := viper.Get("development.database")
	sslmode := viper.Get("development.sslmode")
	dsnFormat := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dsn = fmt.Sprintf(dsnFormat, host, port, user, pw, dbname, sslmode)
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
	var err error
	db, err = sqlt.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
}

// createDB ...
func createDB(cmd *cobra.Command, args []string) error {
	// Get config
	host := viper.GetString("development.host")
	port := viper.GetInt("development.port")
	user := viper.GetString("development.user")
	pw := viper.GetString("development.password")
	ssl := viper.GetString("development.sslmode")

	// Configure database object params for SQL template.
	dbname := viper.GetString("development.database")
	database := sqlt.Database{}
	database.SetName(dbname)
	database.SetOwner(user)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.CreateDBTmpl)
	if err != nil {
		log.Fatalf("Error: createDB: %s\n", err)
	}

	// Format data source name (dsn)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s", host, port, user, pw, ssl)

	// Open a DB connection pool
	db, err := sqlt.Open("postgres", dsn)
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
	// Get config
	host := viper.GetString("development.host")
	port := viper.GetInt("development.port")
	user := viper.GetString("development.user")
	pw := viper.GetString("development.password")
	ssl := viper.GetString("development.sslmode")

	// Configure database object params for SQL template.
	dbname := viper.GetString("development.database")
	database := sqlt.Database{}
	database.SetName(dbname)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.DropDBTmpl)
	if err != nil {
		log.Fatalf("Error: createDB: %s\n", err)
	}

	// Format data source name (dsn)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s", host, port, user, pw, ssl)

	// Open a DB connection pool
	db, err := sqlt.Open("postgres", dsn)
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
