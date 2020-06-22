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
	testAddForeignKeySQL string = `ALTER TABLE cars
	ADD COLUMN people_id bigint;

ALTER TABLE cars
	ADD CONSTRAINT cars_people_mnrk_fkc FOREIGN KEY (people_id)
	REFERENCES people (id);`

	testDropForeignKeySQL string = `ALTER TABLE cars
	DROP CONSTRAINT IF EXISTS cars_people_mnrk_fkc;
	
ALTER TABLE cars
	DROP COLUMN IF EXISTS people_id;`
)

// Unit test createTableMigrations()
func TestAddForeignKeyMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Run addForeignKeyMigration()
	args = append(args, "cars")   // child table
	args = append(args, "people") // parent table
	err := addForeignKeyMigration(cmd, args)
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

	// Verify that the upSQL is as expected
	exp := testAddForeignKeySQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify that the downSQL is as expected
	exp = testDropForeignKeySQL
	act = m.DownSQL()
	if exp != act {
		t.Errorf("\nwant %q;\ngot %q\n", exp, act)
	}
}
