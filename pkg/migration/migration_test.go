package migration

import (
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

// Unit test Migration.Name
func TestMigrationName(t *testing.T) {
	m := new(Migration)
	s := "CreateTableUsers"
	exp := "create_table_users"

	m.SetName(s)    // Set
	act := m.Name() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.LeadingComment
func TestMigrationLeadingComment(t *testing.T) {
	m := new(Migration)
	s := "This is a comment"
	exp := "-- This is a comment"

	m.SetLeadingComment(s)    // Set
	act := m.LeadingComment() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.UpSQL
func TestMigrationUpSQL(t *testing.T) {
	m := new(Migration)
	s := "CREATE TABLE users;"

	exp := s
	m.SetUpSQL(s)    // Set
	act := m.UpSQL() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.UpSQL
func TestMigrationDownSQL(t *testing.T) {
	m := new(Migration)
	s := "DROP TABLE users;"

	exp := s
	m.SetDownSQL(s)    // Set
	act := m.DownSQL() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.SQL
func TestMigrationSQL(t *testing.T) {
	m := new(Migration)
	up := "CREATE TABLE users;"
	down := "DROP TABLE users;"

	exp := `CREATE TABLE users;

-- MIGRATION DELIMITER (DO NOT DELETE THIS COMMENT) --

DROP TABLE users;`

	m.SetUpSQL(up)
	m.SetDownSQL(down)

	act := m.SQL() // Get
	if exp != act {
		t.Errorf("want %q\n; got %q\n", exp, act)
	}
}

// Unit test Migration.Version
func TestMigrationVersion(t *testing.T) {
	m := new(Migration)
	var i int64 = 1234567890

	exp := i
	m.SetVersion(i)    // Set
	act := m.Version() // Get
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Migration.WriteToFile()
func TestMigrationWriteToFile(t *testing.T) {
	m := new(Migration)
	m.SetName("CreateTableUsers")
	m.SetUpSQL("CREATE TABLE users;")
	m.SetDownSQL("DROP TABLE users;")
	m.SetVersion(time.Now().UnixNano())

	fn, err := m.WriteToFile(tmpTestMigrationsDir)
	if err != nil {
		t.Fatal(err)
	}

	exp := m.SQL()
	act, err := fileutil.ReadFileAsString(fn)
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %q\n; got %q\n", exp, act)
	}
}

// Unit test Migration.ReadFromFile()
func TestMigrationReadFromFile(t *testing.T) {
	// Allocate a migration and write it to a file.
	m := new(Migration)
	m.SetName("CreateTableUsers")
	m.SetUpSQL("CREATE TABLE users;")
	m.SetDownSQL("DROP TABLE users;")
	m.SetVersion(time.Now().UnixNano())

	fn, err := m.WriteToFile(tmpTestMigrationsDir)
	if err != nil {
		t.Fatal(err)
	}

	// Allocate a new migration and read in from file.
	rm := new(Migration)
	err = rm.ReadFromFile(fn)
	if err != nil {
		t.Fatal(err)
	}

	// Verify name
	exp := m.Name()
	act := rm.Name()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}

	// Verify upSQL
	exp = m.UpSQL()
	act = rm.UpSQL()
	if exp != act {
		t.Errorf("want %q\n; got %q\n", exp, act)
	}

	// Verify downSQL
	exp = m.DownSQL()
	act = rm.DownSQL()
	if exp != act {
		t.Errorf("want %q\n; got %q\n", exp, act)
	}

	expv := m.Version()
	actv := rm.Version()
	if expv != actv {
		t.Errorf("want %d; got %d", expv, actv)
	}
}
