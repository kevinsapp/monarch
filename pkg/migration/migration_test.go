package migration

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// Unit test Migration.SQL
func TestMigrationSQL(t *testing.T) {
	m := Migration{}
	s := "CREATE TABLE users;"

	exp := s
	m.SetSQL(s)    // Set
	act := m.SQL() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.SQL
func TestMigrationVersion(t *testing.T) {
	m := Migration{}
	var i int64 = 1234567890

	exp := i
	m.SetVersion(i)    // Set
	act := m.Version() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.SQL
func TestMigrationSetFromFile(t *testing.T) {
	m := Migration{}
	ts := time.Now().UnixNano()
	n := fmt.Sprintf("%d_create_table_users_up.sql", ts)

	// Create migration file.
	f, err := os.Create(n)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(n) // Do cleanup.

	// Write SQL to the file (any string would do)
	s := `CREATE TABLE users (
		id bigint NOT NULL,
		created_at timestamp(6) without time zone NOT NULL,
		updated_at timestamp(6) without time zone NOT NULL
	);`
	_, err = f.WriteString(s)
	if err != nil {
		t.Fatal(err)
	}

	// Run SetFromFile method.
	err = m.SetFromFile(n)
	if err != nil {
		t.Fatal(err)
	}

	// Verify version.
	expv := ts
	actv := m.Version() // Get
	if expv != actv {
		t.Errorf("want %d; got %d", expv, actv)
	}

	// Verify SQL.
	exps := s
	acts := m.SQL() // Get
	if exps != acts {
		t.Errorf("want %q; got %q", exps, acts)
	}
}
