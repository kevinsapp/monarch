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
	testAddColumnsSQL string = `-- Table: users

ALTER TABLE users
ADD COLUMN given_name VARCHAR
ADD COLUMN family_name VARCHAR;
`

	testDropColumnsSQL string = `-- Table: users

ALTER TABLE users
DROP COLUMN given_name
DROP COLUMN family_name;
`
)

// Unit test addColumnMigrations()
func TestAddColumnMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Run addColumnMigrations()
	args = append(args, "users")              // table name
	args = append(args, "givenName:varchar")  // first column argument
	args = append(args, "familyName:varchar") // second column argument
	err := addColumnMigrations(cmd, args)     // run command
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
	exp := `_add_columns_to_users_up.sql`
	matched, _ := regexp.MatchString(exp, files[1].Name())
	if !matched {
		t.Errorf("up migration file with name %s not found", exp)
	}

	// Check that the up migration file has the expected content.
	exp = testAddColumnsSQL
	act, err := sql.FileAsString(migrationsDir + "/" + files[1].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("\nwant %s\ngot %s\n", exp, act)
	}

	// Check that the down migration file has the correct name.
	exp = `_add_columns_to_users_down.sql`
	matched, _ = regexp.MatchString(exp, files[0].Name())
	if !matched {
		t.Errorf("down migration file with name %s not found", exp)
	}

	// Check that the down migration file has the expected content.
	exp = testDropColumnsSQL
	act, err = sql.FileAsString(migrationsDir + "/" + files[0].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}
}

// Unit test dropColumnMigrations()
func TestDropColumnMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Run dropColumnMigrations()
	args = append(args, "users")           // table name
	args = append(args, "givenName")       // first column argument
	args = append(args, "familyName")      // second column argument
	err := dropColumnMigrations(cmd, args) // run command
	if err != nil {
		t.Fatal(err)
	}

	// Get the list of files in the migrations directory.
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check that exactly one file was created.
	if l := len(files); l != 1 {
		t.Errorf("wrong number of files created: want 1; got %d", l)
	}

	// Check that the up migration file has the correct name.
	exp := `_drop_columns_from_users_up.sql`
	matched, _ := regexp.MatchString(exp, files[0].Name())
	if !matched {
		t.Errorf("up migration file with name %s not found", exp)
	}

	// Check that the up migration file has the expected content.
	exp = testDropColumnsSQL
	act, err := sql.FileAsString(migrationsDir + "/" + files[0].Name())
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("\nwant %s\ngot %s\n", exp, act)
	}

	if exp != act {
		t.Errorf("want %s\n; got %s\n", exp, act)
	}
}
