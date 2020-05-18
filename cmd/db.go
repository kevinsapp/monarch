package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

// Global ...
var db *sql.DB

func init() {
	rootCmd.AddCommand(dbCmd)
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use:              "db",
	PersistentPreRun: openDB,
}

// openDB ...
func openDB(cmd *cobra.Command, args []string) {
	// h := "localhost"
	// p := 5432
	// u := "postgres"
	// pw := "postgres"
	// dn := "monarch_development"
	// ssl := "disable"

	h := viper.Get("development.host")
	p := viper.Get("development.port")
	u := viper.Get("development.user")
	pw := viper.Get("development.password")
	dn := viper.Get("development.database")
	ssl := viper.Get("development.sslmode")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", h, p, u, pw, dn, ssl)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
}
