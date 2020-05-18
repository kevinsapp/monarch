package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

// Unit test dbPing()
func TestDBPing(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Open the DB connection pool.
	openDB(cmd, args)

	// Run ping() and verify that no errors occur.
	err := ping(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Open DB with invalid password in DSN.
	dsn := "host=localhost port=5432 user=postgres password=wrongpw dbname=monarch_development sslmode=disable"
	db, _ = sqlt.Open("postgres", dsn)

	// Run ping() and verify that the correct error is returned.
	err = ping(cmd, args)
	exp := `pq: password authentication failed for user "postgres"`
	act := err.Error()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}
