package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

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
