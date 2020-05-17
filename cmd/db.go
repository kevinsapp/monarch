package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	// pg
	_ "github.com/lib/pq"
)

// Global ...
var db *sql.DB

func init() {
	rootCmd.AddCommand(dbCmd)
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use: "db",
	// Run: openDB,
	PersistentPreRun: openDB,
}

// openDB ...
func openDB(cmd *cobra.Command, args []string) {
	host := "localhost"
	port := 5432
	user := "postgres"
	pw := "postgres"
	dbname := "monarch_development"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pw, dbname)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
}
