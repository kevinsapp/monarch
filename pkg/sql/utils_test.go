package sql

import (
	"os"
	"testing"
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

	testRenameTableSQL string = `-- Table: users

ALTER TABLE users RENAME TO people;
`
)

// Unit test templateAsSQL
func TestProcessTmpl(t *testing.T) {
	cases := []struct {
		Table table
		Tmpl  string
		SQL   string
	}{
		{table{"users", ""}, CreateTableTmpl, testCreateTableSQL},       // Create table
		{table{"users", ""}, DropTableTmpl, testDropTableSQL},           // Drop table
		{table{"users", "people"}, RenameTableTmpl, testRenameTableSQL}, // Rename table
	}

	for _, c := range cases {
		td := c.Table
		tmpl := c.Tmpl
		exp := c.SQL

		// Run ProcessTmpl()
		act, err := ProcessTmpl(td, tmpl)
		if err != nil {
			t.Fatal(err)
		}

		// Check that processed SQL matches the expected SQL.
		if exp != act {
			t.Errorf("want %s; got %s", exp, act)
		}
	}
}

func TestFileAsString(t *testing.T) {
	fname := "../../test.sql"
	defer os.Remove(fname)

	cases := []string{
		testCreateTableSQL,
		testDropTableSQL,
		testRenameTableSQL,
	}

	for _, c := range cases {
		// Create a file.
		f, err := os.Create(fname)
		if err != nil {
			t.Fatal(err)
		}

		// Write SQL to the file.
		_, err = f.WriteString(c)
		if err != nil {
			t.Fatal(err)
		}
		f.Close()

		// Run ProcessTmpl()
		act, err := FileAsString(fname)
		if err != nil {
			t.Fatal(err)
		}

		// Check that processed SQL matches the expected SQL.
		exp := c
		if exp != act {
			t.Errorf("want %s; got %s", exp, act)
		}
	}
}
