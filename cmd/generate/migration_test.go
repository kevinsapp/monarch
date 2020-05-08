package generate

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

// Unit test mkdirMigrations()
func TestMkdirMigrations(t *testing.T) {
	var cmd = &cobra.Command{}
	var args []string
	dir := "migrations"

	// Ensure that migrations directory does not exist before running the
	// function we're testing.
	os.RemoveAll(dir)

	// Run mkdirMigrations()
	mkdirMigrations(cmd, args)

	// Do cleanup
	defer os.RemoveAll(dir)

	// Check for the presence of the migrations directory
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
		t.Errorf("want file mode %s; got %s", exp, act)
	}
}

// Unit test templateAsSQL
func TestTemplateAsSQL(t *testing.T) {
	// Define the expected SQL string.
	exp := `--Up migration for users table

CREATE TABLE users (
	id uuid DEFAULT gen_random_uuid() NOT NULL,

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
`
	// Define a SQL template.
	sql := `--Up migration for {{.Name}} table

CREATE TABLE {{.Name}} (
	id uuid DEFAULT gen_random_uuid() NOT NULL,

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT {{.Name}}_pkey PRIMARY KEY (id)
);
`
	// Define a data structure to apply to the SQL template.
	data := struct {
		Name string
	}{
		Name: "users",
	}

	// Run templateAsSQK()
	act, err := templateAsSQL(data, sql)
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %s; got %s", exp, act)
	}
}
