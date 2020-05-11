package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/kevinsapp/monarch/pkg/sql"
	"github.com/spf13/cobra"
)

// Test data - expected SQL
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

// Unit test createTableMigrations()
func TestCreateTableMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Run createTableMigrations()
	args = append(args, "users")
	err := createTableMigrations(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Get the list of files in the migrations directory.
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check that exactly two files were created.
	if l := len(files); l != 2 {
		t.Errorf("wrong number of files created: want 2; got %d", l)
	}

	// Check that the up migration file has the correct name.
	exp := `_create_table_users_up.sql`
	matched, _ := regexp.MatchString(exp, files[1].Name())
	if !matched {
		t.Errorf("up migration file with name %s not found", exp)
	}

	// Check that the up migration file has the expected content.
	exp = testCreateTableSQL
	act, err := sql.FileAsString(migrationsDir + "/" + files[1].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}

	// Check that the down migration file has the correct name.
	exp = `_create_table_users_down.sql`
	matched, _ = regexp.MatchString(`_create_table_users_down.sql`, files[0].Name())
	if !matched {
		t.Errorf("down migration file with name %s not found", exp)
	}

	// Check that the down migration file has the expected content.
	exp = testDropTableSQL
	act, err = sql.FileAsString(migrationsDir + "/" + files[0].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}
}

// Unit test dropTableMigrations()
func TestDropTableMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll("migrations") // Do cleanup

	// Run dropTableMigrations()
	args = append(args, "users")
	err := dropTableMigrations(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Get the list of files in the migrations directory.
	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		t.Fatal(err)
	}

	// Check that exactly one file was created.
	if l := len(files); l != 1 {
		t.Errorf("wrong number of files created: want 1; got %d", l)
	}

	// Check that the up migration file has to correct name.
	exp := `_drop_table_users_up.sql`
	matched, _ := regexp.MatchString(exp, files[0].Name())
	if !matched {
		t.Errorf("up migration file with name %s not found", exp)
	}

	// Check that the up migration file has the expected content.
	exp = testDropTableSQL
	act, err := sql.FileAsString(migrationsDir + "/" + files[0].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}
}

// Unit test createTableMigrations()
func TestRenameTableMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll("migrations") // Do cleanup

	// Run renameTableMigrations()
	args = append(args, "users")  // first argument
	args = append(args, "people") // second argument
	err := renameTableMigrations(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Get the list of files in the migrations directory.
	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		t.Fatal(err)
	}

	// Check that exactly two files were created.
	if l := len(files); l != 2 {
		t.Errorf("wrong number of files created: want 2; got %d", l)
	}

	// Check that the up migration has the correct name.
	exp := `_rename_table_users_up.sql`
	matched, _ := regexp.MatchString(exp, files[1].Name())
	if !matched {
		t.Errorf("up migration file with name %s not found", exp)
	}

	// Check that the up migration file has the expected content.
	exp = testRenameTableUpSQL
	act, err := sql.FileAsString(migrationsDir + "/" + files[1].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}

	// Check that the down migration file was created.
	exp = `_rename_table_users_down.sql`
	matched, _ = regexp.MatchString(exp, files[0].Name())
	if !matched {
		t.Errorf("down migration file with name %s not found", exp)
	}

	// Check that the down migration file has the expected content.
	exp = testRenameTableDownSQL
	act, err = sql.FileAsString(migrationsDir + "/" + files[0].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}
}
