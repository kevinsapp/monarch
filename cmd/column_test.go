package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kevinsapp/monarch/pkg/migration"
	"github.com/spf13/cobra"
)

// Test data - expected SQL
const (
	testAddColumnsSQL string = `ALTER TABLE users
ADD COLUMN given_name varchar
ADD COLUMN family_name varchar;`

	testDropColumnsSQL string = `ALTER TABLE users
DROP COLUMN given_name
DROP COLUMN family_name;`

	testRenameColumnsUpSQL string = `ALTER TABLE users
RENAME COLUMN given_name TO first_name
RENAME COLUMN family_name TO last_name;`

	testRenameColumnsDownSQL string = `ALTER TABLE users
RENAME COLUMN first_name TO given_name
RENAME COLUMN last_name TO family_name;`
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

	// Verify that exactly one file was created.
	if l := len(files); l != 1 {
		t.Errorf("wrong number of files created: want 1; got %d", l)
	}

	// Verify that the file can be read in to a migration object.
	path := fmt.Sprintf("%s/%s", migrationsDir, files[0].Name())
	m := new(migration.Migration)
	err = m.ReadFromFile(path)
	if err != nil {
		t.Error(err)
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

	// Verify that the file can be read in to a migration object.
	path := fmt.Sprintf("%s/%s", migrationsDir, files[0].Name())
	m := new(migration.Migration)
	err = m.ReadFromFile(path)
	if err != nil {
		t.Error(err)
	}
}

// Unit test addColumnMigrations()
func TestRenameColumnMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Run addColumnMigrations()
	args = append(args, "users")               // table name
	args = append(args, "givenName:firstName") // first column argument
	args = append(args, "familyName:lastName") // second column argument
	err := renameColumnMigrations(cmd, args)   // run command
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

	// Verify that the file can be read in to a migration object.
	path := fmt.Sprintf("%s/%s", migrationsDir, files[0].Name())
	m := new(migration.Migration)
	err = m.ReadFromFile(path)
	if err != nil {
		t.Error(err)
	}
}
