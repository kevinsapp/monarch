package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/spf13/cobra"
)

// Unit test createTableMigrations()
func TestCreateTableMigrations(t *testing.T) {
	// Create a migrations directory.
	var cmd = &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll("migrations") // Do cleanup

	// Run createTableMigrations()
	args = append(args, "users")
	err := createTableMigrations(cmd, args)
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

	// Check that the up migration file was created.
	matched, _ := regexp.MatchString(`_create_table_users_up.sql`, files[1].Name())
	if !matched {
		t.Error("up migration file not found")
	}

	// Check that the down migration file was created.
	matched, _ = regexp.MatchString(`_create_table_users_down.sql`, files[0].Name())
	if !matched {
		t.Error("down migration file not found")
	}
}

// Unit test dropTableMigrations()
func TestDropTableMigrations(t *testing.T) {
	// Create a migrations directory.
	var cmd = &cobra.Command{}
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

	// Check that the up migration file was created.
	matched, _ := regexp.MatchString(`_drop_table_users_up.sql`, files[0].Name())
	if !matched {
		t.Error("up migration file not found")
	}
}

// Unit test createTableMigrations()
func TestRenameTableMigrations(t *testing.T) {
	// Create a migrations directory.
	var cmd = &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll("migrations") // Do cleanup

	// Run createTableMigrations()
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

	// Check that the up migration file was created.
	matched, _ := regexp.MatchString(`_rename_table_users_up.sql`, files[1].Name())
	if !matched {
		t.Error("up migration file not found")
	}

	// Check that the down migration file was created.
	matched, _ = regexp.MatchString(`_rename_table_users_down.sql`, files[0].Name())
	if !matched {
		t.Error("down migration file not found")
	}
}
