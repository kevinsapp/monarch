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
	m := Migration{}
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
	m := Migration{}
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
	m := Migration{}
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
	m := Migration{}
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
	m := Migration{}
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
	m := Migration{}
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
	m, fn := makeMigrationHelper("users", t)

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
	m, fn := makeMigrationHelper("users", t)

	// Allocate a new migration and read in from file.
	rm := new(Migration)
	err := rm.ReadFromFile(fn)
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

func TestLoadAllLaterThan(t *testing.T) {
	// Migration one
	makeMigrationHelper("one", t)

	// Generate a timestamp
	version := time.Now().UnixNano()

	// Migrations two and three; later than "version" timestamp.
	makeMigrationHelper("two", t)
	makeMigrationHelper("three", t)

	// Run LoadAllLaterThan() - later than version
	migrations, err := LoadAllLaterThan(version, tmpTestMigrationsDir)
	if err != nil {
		t.Error(err)
	}

	// Should have loaded exactly two migrations (the latter two).
	if count := len(migrations); count != 2 {
		t.Errorf("want 2; got %d", count)
	}

	// Verifiy migration 2 version is greater than "version" timestamp
	expv := version
	actv := migrations[0].Version()
	if expv >= actv {
		t.Errorf("\nversion %d not later than %d\n", expv, actv)
	}

	// Verify migration 2 name
	exp := "create_table_two"
	act := migrations[0].Name()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify migration 2 upSQL
	exp = "CREATE TABLE two;"
	act = migrations[0].UpSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify migration 2 downSQL
	exp = "DROP TABLE two;"
	act = migrations[0].DownSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verifiy migration 3 version is greater than "version" timestamp
	expv = version
	actv = migrations[1].Version()
	if expv >= actv {
		t.Errorf("\nversion %d not later than %d\n", expv, actv)
	}

	// Verify migration 3 name
	exp = "create_table_three"
	act = migrations[1].Name()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify migration 3 upSQL
	exp = "CREATE TABLE three;"
	act = migrations[1].UpSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}

	// Verify migration 3 downSQL
	exp = "DROP TABLE three;"
	act = migrations[1].DownSQL()
	if exp != act {
		t.Errorf("\nwant %q\n got %q\n", exp, act)
	}
}

// makeMigrationHelper creates a migration to create a table with name "tn"
// the writes it to a file by calling WriteToFile().
func makeMigrationHelper(tn string, t *testing.T) (Migration, string) {
	m := Migration{}
	m.SetName("CreateTable_" + tn)
	m.SetUpSQL("CREATE TABLE " + tn + ";")
	m.SetDownSQL("DROP TABLE " + tn + ";")
	m.SetVersion(time.Now().UnixNano())
	f, err := m.WriteToFile(tmpTestMigrationsDir)
	if err != nil {
		t.Fatal(err)
	}

	return m, f
}
