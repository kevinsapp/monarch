package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCreateDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Drop DB
	err := dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Run createDB() and verify that no errors occur.
	err = createDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDropDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Run dropDB() and verify that no errors occur.
	err := dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Run createDB(): need a database to drop.
	err = createDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Run dropDB() and verify that no errors occur.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

// Unit test dbPing()
func TestPingDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Reset DB
	err := resetDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Open the DB connection pool.
	openDB(cmd, args)

	// Run pingDB() and verify that no errors occur.
	err = pingDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	// Close the global db connection pool and drop the database to avoid conflicts with other tests.
	db.Close()
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

func TestResetDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Run ResetDB().
	err := resetDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}
