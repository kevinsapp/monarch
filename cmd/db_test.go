package cmd

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
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

func TestCopyDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Create a DB with the default name.
	err := resetDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
	defer dropDB(cmd, args) // Cleanup source db.

	// Run copyDB() and verify that no errors occur.
	args = append(args, "monarch_development")
	args = append(args, "copy_target_db")
	err = copyDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Establish a connection to the new db.
	var srv dbServer
	srv.initFromConfig()
	srv.dbName = "copy_target_db" // dbName should be set to new db.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		t.Fatal(err)
	}

	// Verify that we can ping the new db.
	err = conn.Ping(ctx)
	if err != nil {
		t.Error(err)
	}

	// Execute query to drop new database (copy target).
	err = conn.Close(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Drop the copied db.
	srv.initFromConfig()
	srv.dbName = "" // dbName should be blank before getting DSN.
	conn, err = pgx.Connect(ctx, srv.dsn())
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Exec(ctx, "DROP DATABASE copy_target_db;")
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

	// Run pingDB() and verify that no errors occur.
	err = pingDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	// Do cleanup.
	err = dropDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenameDB(t *testing.T) {
	// Set up arguments.
	cmd := &cobra.Command{}
	args := make([]string, 0)

	// Initialize configuration from config file.
	initConfig()

	// Create a DB with the default name.
	err := resetDB(cmd, args)
	if err != nil {
		t.Fatal(err)
	}
	defer dropDB(cmd, args) // Try to cleanup if renaming fails.

	// Run renameDB() and verify that no errors occur.
	args = append(args, "monarch_development")
	args = append(args, "renamed_test_db")
	err = renameDB(cmd, args)
	if err != nil {
		t.Error(err)
	}

	var srv dbServer
	srv.initFromConfig()
	srv.dbName = "" // dbName should be blank before getting DSN.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	// Execute query to drop database by (new) name.
	_, err = conn.Exec(ctx, "DROP DATABASE renamed_test_db;")
	if err != nil {
		t.Error(err)
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
