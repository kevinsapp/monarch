package migration

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kevinsapp/monarch/pkg/fileutil"
)

const (
	tmpDir               string = "tmp"
	tmpTestMigrationsDir string = "tmp/test/migrations/"
)

func TestMain(m *testing.M) {
	// Setup
	fileutil.MkdirP(tmpTestMigrationsDir)

	// Execute tests.
	i := m.Run()

	// Teardown
	os.RemoveAll(tmpDir) // Do cleanup

	// Exit
	os.Exit(i)
}

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

// Unit test Migration.SetFromFile
func TestMigrationSetFromFile(t *testing.T) {
	m := Migration{}
	ts := time.Now().UnixNano()
	n := fmt.Sprintf("%d_create_table_users_up.sql", ts)
	path := tmpTestMigrationsDir + n

	// Write SQL to the file (any string would do)
	sql := `CREATE TABLE users (
		id bigint NOT NULL,
		created_at timestamp(6) without time zone NOT NULL,
		updated_at timestamp(6) without time zone NOT NULL
	);`
	err := fileutil.CreateAndWriteString(path, sql)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	// Run SetFromFile method.
	err = m.SetFromFile(path)
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
	exps := sql
	acts := m.SQL() // Get
	if exps != acts {
		t.Errorf("want %q\n; got %q\n", exps, acts)
	}
}

// Unit test ExtractVersionFromFile
func TestExtractVersionFromFile(t *testing.T) {
	ts := time.Now().UnixNano()
	n := fmt.Sprintf("%d_create_table_users_up.sql", ts)
	path := tmpTestMigrationsDir + n

	// Create migration file.
	_, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	// Run ExtractVersionFromFile
	v, err := ExtractVersionFromFile(path)
	if err != nil {
		t.Fatal(err)
	}

	// Verify version.
	exp := ts
	act := v
	if exp != act {
		t.Errorf("want %d; got %d", exp, act)
	}
}
