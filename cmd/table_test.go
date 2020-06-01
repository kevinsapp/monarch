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
	testCreateTableSQL string = `CREATE TABLE users (
	PRIMARY KEY (id),
	id bigserial NOT NULL,

	-- Specify additional fields here.

	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL
);`

	testCreateTableWithColsSQL string = `CREATE TABLE users (
	PRIMARY KEY (id),
	id bigserial NOT NULL,
	user_name varchar,
	given_name varchar,
	family_name varchar,
	locale varchar,
	active boolean,

	-- Specify additional fields here.

	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL
);`
	testDropTableSQL string = `DROP TABLE users;`

	testRenameTableUpSQL string = `ALTER TABLE users RENAME TO people;`

	testRenameTableDownSQL string = `ALTER TABLE people RENAME TO users;`
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
	err := createTableMigration(cmd, args)
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
	exp := testCreateTableSQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify that the downSQL is as expected
	exp = testDropTableSQL
	act = m.DownSQL()
	if exp != act {
		t.Errorf("\nwant %q;\ngot %q\n", exp, act)
	}
}

// Unit test createTableWithColsMigrations()
func TestCreateTableWithColsMigrations(t *testing.T) {
	// Create a migrations directory.
	cmd := &cobra.Command{}
	args := make([]string, 0)
	mkdirMigrations(cmd, args)
	defer os.RemoveAll(migrationsDir) // Do cleanup

	// Set up args
	args = append(args, "users")              // table
	args = append(args, "userName:varchar")   // column
	args = append(args, "givenName:varchar")  // column
	args = append(args, "familyName:varchar") // column
	args = append(args, "locale:varchar")     // column
	args = append(args, "active:boolean")     // column

	// Run createTableMigrations()
	err := createTableMigration(cmd, args)
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
	exp := testCreateTableWithColsSQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify that the downSQL is as expected
	exp = testDropTableSQL
	act = m.DownSQL()
	if exp != act {
		t.Errorf("\nwant %q;\ngot %q\n", exp, act)
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
	err := dropTableMigration(cmd, args)
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

	// Verify that the file can be read in to a migration object.
	path := fmt.Sprintf("%s/%s", migrationsDir, files[0].Name())
	m := new(migration.Migration)
	err = m.ReadFromFile(path)
	if err != nil {
		t.Error(err)
	}

	// Verify that the upSQL is as expected.
	exp := testDropTableSQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("\nwant %q;\n got %q\n", exp, act)
	}

	// Verify that the downSQL is as expected.
	exp = "" // downSQL should be blank since this migratin is not reversible.
	act = m.DownSQL()
	if exp != act {
		t.Errorf("\nwant %q;\ngot %q\n", exp, act)
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
	err := renameTableMigration(cmd, args)
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

	// Verify that the file can be read in to a migration object.
	path := fmt.Sprintf("%s/%s", migrationsDir, files[0].Name())
	m := new(migration.Migration)
	err = m.ReadFromFile(path)
	if err != nil {
		t.Error(err)
	}

	// Verify that the upSQL is as expected.
	exp := testRenameTableUpSQL
	act := m.UpSQL()
	if exp != act {
		t.Errorf("\nwant %q;\n got %q\n", exp, act)
	}

	// Verify that the downSQL is as expected.
	exp = testRenameTableDownSQL
	act = m.DownSQL()
	if exp != act {
		t.Errorf("\nwant %q;\ngot %q\n", exp, act)
	}
}
