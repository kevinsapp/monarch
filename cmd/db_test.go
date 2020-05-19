package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestCreateDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Drop DB
	err := dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Run createDB() and verify that no errors occur.
	err = createDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDropDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Run dropDB() and verify that no errors occur.
	err := dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Run createDB(): need a database to drop.
	err = createDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Run dropDB() and verify that no errors occur.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

// Unit test dbPing()
func TestPingDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Reset DB
	err := resetDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}

	// Open the DB connection pool.
	openDB(cmd, args)

	// Run pingDB() and verify that no errors occur.
	err = pingDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	// Close the global db connection pool and drop the database to avoid conflicts with other tests.
	db.Close()
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

func TestResetDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Run ResetDB().
	err := resetDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDBServerDSN(t *testing.T) {
	var srv dbServer
	srv.host = "abcd"
	srv.port = 1234
	srv.user = "abcd"
	srv.password = "abcd"
	srv.dbName = "test_db"
	srv.sslMode = "disable"

	// Verify dsn with dbName set.
	exp := "host=abcd port=1234 user=abcd password=abcd dbname=test_db sslmode=disable"
	act := srv.dsn()
	if exp != act {
		t.Errorf("want %q;\n got %q\n", exp, act)
	}

	// Verify dsn with dbName blank.
	srv.dbName = ""
	exp = "host=abcd port=1234 user=abcd password=abcd sslmode=disable"
	act = srv.dsn()
	if exp != act {
		t.Errorf("want %q;\n got %q\n", exp, act)
	}
}

func TestDBServerInitFromConfig(t *testing.T) {
	viper.Set("development.host", "ahost")
	viper.Set("development.port", 1234)
	viper.Set("development.user", "auser")
	viper.Set("development.password", "apassword")
	viper.Set("development.database", "adb")
	viper.Set("development.sslmode", "disable")

	var srv dbServer
	srv.initFromConfig()

	exp := srv.host
	act := viper.GetString("development.host")
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}

	exp = srv.user
	act = viper.GetString("development.user")
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}

	exp = srv.password
	act = viper.GetString("development.password")
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}

	exp = srv.dbName
	act = viper.GetString("development.database")
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}

	exp = srv.sslMode
	act = viper.GetString("development.sslmode")
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}

	expPort := srv.port
	actPort := viper.GetInt("development.host")
	if exp != act {
		t.Errorf("want %q; got %q", expPort, actPort)
	}
}
