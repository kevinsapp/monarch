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
	testCreateDefaultIndexSQL string = `CREATE INDEX cars_color_mnrk_idx ON cars (color);`
	testDropIndexSQL          string = `DROP INDEX IF EXISTS cars_color_mnrk_idx;`
)

// Unit test createIndexMigration()
func TestCreateIndexMigration(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Run addColumnMigrations()
	args = append(args, "cars")            // table name
	args = append(args, "color")           // column name
	err := createIndexMigration(cmd, args) // run command
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

	// Verify that the upSQL is as expected.
	exp := testCreateDefaultIndexSQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("\nwant %q;\n got %q\n", exp, act)
	}

	// Verify that the downSQL is as expected.
	exp = testDropIndexSQL
	act = m.DownSQL()
	if exp != act {
		t.Errorf("\nwant %q;\n got %q\n", exp, act)
	}
}
