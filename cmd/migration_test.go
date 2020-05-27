package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kevinsapp/monarch/pkg/migration"
	"github.com/spf13/cobra"
)

func TestCreateMigration(t *testing.T) {
	// Create a migrations directory
	var cmd = &cobra.Command{}
	var args []string
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	name := "CreateTableUsers"
	upSQL := "CREATE TABLE users;"
	downSQL := "DROP TABLE users;"

	err := createMigration(name, upSQL, downSQL)
	if err != nil {
		t.Error(err)
	}

	// Get the list of files in the directory specificed by path.
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		t.Fatal(err)
	}

	// Allocate a new migration and read in from file.
	path := fmt.Sprintf("%s/%s", migrationsDir, files[0].Name())
	m := new(migration.Migration)
	err = m.ReadFromFile(path)
	if err != nil {
		t.Fatal(err)
	}

	// Verify upSQL
	exp := upSQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("want %q;\ngot%q\n", exp, act)
	}

	// Verify downSQL
	exp = downSQL
	act = m.DownSQL()
	if exp != act {
		t.Errorf("want %q;\ngot%q\n", exp, act)
	}
}

// Unit test mkdirMigrations()
func TestMkdirMigrations(t *testing.T) {
	// Ensure that a migrations directory does not already exist.
	os.RemoveAll(migrationsDir)

	// Run mkdirMigrations()
	var cmd = &cobra.Command{}
	var args []string
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Check that the migrations directory exists.
	stat, err := os.Stat(migrationsDir)
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
