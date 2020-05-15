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
	testAddColumnSQL string = `-- Table: users

ALTER TABLE users
ADD COLUMN given_name VARCHAR
ADD COLUMN family_name VARCHAR;
`
	testDropColumnSQL string = `-- Table: users

ALTER TABLE users
DROP COLUMN given_name;
`

	testRenameColumnSQL string = `-- Table: users

ALTER TABLE users
RENAME COLUMN given_name TO first_name
RENAME COLUMN family_name TO last_name;
`
)

// Unit test templateAsSQL
func TestProcessTmpl(t *testing.T) {
	// Test cases
	cases := []struct {
		Table Table
		Tmpl  string
		SQL   string
	}{
		{Table{"users", "", []Column{}}, CreateTableTmpl, testCreateTableSQL},       // Create table
		{Table{"users", "", []Column{}}, DropTableTmpl, testDropTableSQL},           // Drop table
		{Table{"users", "people", []Column{}}, RenameTableTmpl, testRenameTableSQL}, // Rename table
		{ // Add column to table
			Table{
				"users",
				"",
				[]Column{
					{"given_name", "", "VARCHAR"},
					{"family_name", "", "VARCHAR"},
				},
			},
			AddColumnTmpl,
			testAddColumnSQL,
		},
		{ // Drop column from table
			Table{
				"users",
				"",
				[]Column{
					{"given_name", "", ""},
				},
			},
			DropColumnTmpl,
			testDropColumnSQL,
		},
		{ // Rename column in table
			Table{
				"users",
				"",
				[]Column{
					{"given_name", "first_name", ""},
					{"family_name", "last_name", ""},
				},
			},
			RenameColumnTmpl,
			testRenameColumnSQL,
		},
	}

	// Run each test case.
	for _, c := range cases {
		td := c.Table
		tmpl := c.Tmpl
		exp := c.SQL

		// Run ProcessTmpl()
		act, err := ProcessTmpl(&td, tmpl)
		if err != nil {
			t.Fatal(err)
		}

		// Check that processed SQL matches the expected SQL.
		if exp != act {
			t.Errorf("\nwant %s\ngot %s", exp, act)
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
