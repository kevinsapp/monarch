package cmd

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

const (
	testCreateTableSQL string = `-- Table: users

CREATE TABLE users (
	id uuid DEFAULT gen_random_uuid() NOT NULL,

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
`
	testDropTableSQL string = `-- Table: users

DROP TABLE users;
`

	testRenameTableUpSQL string = `-- Table: users

ALTER TABLE users RENAME TO people;
`

	testRenameTableDownSQL string = `-- Table: people

ALTER TABLE people RENAME TO users;
`
)

// Unit test createMigration()
func TestCreateMigration(t *testing.T) {
	// Create a migrations directory.
	var cmd = &cobra.Command{}
	var args []string
	mkdirMigrations(cmd, args)
	defer os.RemoveAll("migrations") // Do cleanup

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := table{}
	td.Name = "users"

	// Create a migration file.
	fn := fmt.Sprintf("migrations/%d_create_table_%s_up.sql", timestamp, td.Name)
	f, err := createMigration(fn, sqltCreateTable, td)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the migration file has the expected name.
	exp := fn
	act := f.Name()
	if exp != act {
		t.Errorf("want %s; got %s", exp, act)
	}

	// Check that the migration file contains the expected content.
	exp = testCreateTableSQL
	act, err = fileAsString(fn)
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s; got %s", exp, act)
	}

}

// Unit test mkdirMigrations()
func TestMkdirMigrations(t *testing.T) {
	// Ensure that a migrations directory does not already exist.
	dir := "migrations"
	os.RemoveAll(dir)

	// Run mkdirMigrations()
	var cmd = &cobra.Command{}
	var args []string
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(dir) // Do cleanup

	// Check that the migrations directory exists.
	stat, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Check that migrations is a directory and not some other type of file
	if !stat.IsDir() {
		t.Errorf("want directory true; got directory false")
	}

	// Check that the migrations directory has the correct Unix permissions
	exp := "-rwxr-xr-x"                // expected Unix permissions
	act := stat.Mode().Perm().String() // actual Unix permissions
	if exp != act {
		t.Errorf("want %s; got %s", exp, act)
	}
}

// Unit test templateAsSQL
func TestTemplateAsSQL(t *testing.T) {
	// Set template data.
	td := table{}
	td.Name = "users"

	// Run templateAsSQL()
	act, err := templateAsSQL(td, sqltCreateTable)
	if err != nil {
		t.Fatal(err)
	}

	// Check that processed SQL matches the expected SQL.
	exp := testCreateTableSQL
	if exp != act {
		t.Errorf("want %s; got %s", exp, act)
	}
}
