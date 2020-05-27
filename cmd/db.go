package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(createDBCmd)
	dbCmd.AddCommand(dropDBCmd)
	dbCmd.AddCommand(pingDBCmd)
	dbCmd.AddCommand(renameDBCmd)
	dbCmd.AddCommand(resetDBCmd)
}

// dbCmd ...
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: `Provides subcommands for working with databases.`,
}

// createCmd ...
var createDBCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a database with the name specificed by the "database" attribute in the config file.`,
	RunE:  createDB,
}

// dropCmd ...
var dropDBCmd = &cobra.Command{
	Use:   "drop",
	Short: `Drop a database with the name specificed by the "database" attribute in the config file.`,
	RunE:  dropDB,
}

// pingDBCmd ...
var pingDBCmd = &cobra.Command{
	Use:   "ping",
	Short: `Verifies that Monarch can connect to the database specificed in the config file.`,
	RunE:  pingDB,
}

// renameDBCmd ...
var renameDBCmd = &cobra.Command{
	Use:   "rename [oldname] [newname]",
	Short: `Rename a database from name to newname.`,
	Long: `Rename a database from name to newname. Note: if you change the name of your database,
	you should also change the database name specificed in your config file.`,
	RunE: renameDB,
}

// resetDBCmd ...
var resetDBCmd = &cobra.Command{
	Use:   "reset",
	Short: `First drops and then creates a database with the name specificed by the "database" attribute in the config file.`,
	RunE:  resetDB,
}

// createDB creates a database with the name specificed by the "database"
// attribute in the viper config.
func createDB(cmd *cobra.Command, args []string) error {
	// Initialize a dbServer object.
	var srv dbServer
	srv.initFromConfig()

	// Configure a data object to apply to a SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)
	database.SetOwner(srv.user)

	// Process the SQL template.
	query, err := sqlt.ProcessTmpl(&database, sqlt.CreateDBTmpl)
	if err != nil {
		log.Fatalf("ERROR: createDB: %s\n", err)
	}

	// Connect to the database server.
	srv.dbName = "" // dbName should be blank before getting DSN.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: createDB: %s\n", err)
	}
	defer conn.Close(ctx)

	// Execute query to create database.
	start := time.Now()
	_, err = conn.Exec(ctx, query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q created. Command completed in %s.\n", database.Name(), duration)

	return err
}

// dropDB drops a database with the name specificed by the "database" attribute
// in the viper config.
func dropDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Configure a data object to apply to a SQL template.
	database := sqlt.Database{}
	database.SetName(srv.dbName)

	// Process SQL template
	query, err := sqlt.ProcessTmpl(&database, sqlt.DropDBTmpl)
	if err != nil {
		log.Printf("ERROR: dropDB: %s\n", err)
		return err
	}

	// Connect to the database server.
	srv.dbName = "" // dbName should be blank before getting dsn.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: dropDB: %s\n", err)
	}
	defer conn.Close(ctx)

	// Execute query to drop database.
	start := time.Now()
	_, err = conn.Exec(ctx, query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q dropped. Command completed in %s.\n", database.Name(), duration)

	return err
}

// ping connects to the database to verify that the server is accessible.
func pingDB(cmd *cobra.Command, args []string) error {
	var srv dbServer
	srv.initFromConfig()

	// Connect to the database server.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	start := time.Now()
	err = conn.Ping(ctx)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database connection OK. Command completed in %s.\n", duration)

	return err
}

// renameDB renames a database from name to new name.
func renameDB(cmd *cobra.Command, args []string) error {
	// Initialize a dbServer object.
	var srv dbServer
	srv.initFromConfig()

	// Configure a data object to apply to a SQL template.
	database := sqlt.Database{}
	database.SetName(args[0])
	database.SetNewName(args[1])

	// Process the SQL template.
	query, err := sqlt.ProcessTmpl(&database, sqlt.RenameDBTmpl)
	if err != nil {
		log.Fatalf("ERROR: renameDB: %s\n", err)
	}

	// Connect to the database server.
	srv.dbName = "" // dbName should be blank before getting DSN.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, srv.dsn())
	if err != nil {
		log.Fatalf("ERROR: renameDB: %s\n", err)
	}
	defer conn.Close(ctx)

	// Execute query to create database.
	start := time.Now()
	_, err = conn.Exec(ctx, query)
	duration := time.Since(start)
	if err != nil {
		return err
	}

	fmt.Printf("Database %q renamed to %q. Command completed in %s.\n", database.Name(), database.NewName(), duration)

	return err
}

// resetDB drops and creates a database, i.e. reset.
func resetDB(cmd *cobra.Command, args []string) error {
	err := dropDB(cmd, args)
	if err != nil {
		return err
	}

	err = createDB(cmd, args)

	return err
}

// dbServer
type dbServer struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

// dsn returns a Data Source Name (dsn) string based on the dbServer attributes.
func (s *dbServer) dsn() string {
	// If dbName is not set, format a data source name without a dbname and return it.
	if s.dbName == "" {
		format := "host=%s port=%d user=%s password=%s sslmode=%s"
		dsn := fmt.Sprintf(format, s.host, s.port, s.user, s.password, s.sslMode)
		return dsn
	}

	// Format a data source name.
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dsn := fmt.Sprintf(format, s.host, s.port, s.user, s.password, s.dbName, s.sslMode)

	return dsn
}

// intiFromConfig initalizes a dbServer{} from the viper config.
func (s *dbServer) initFromConfig() {
	// Read in config.
	s.host = viper.GetString("development.host")
	s.port = viper.GetInt("development.port")
	s.user = viper.GetString("development.user")
	s.password = viper.GetString("development.password")
	s.dbName = viper.GetString("development.database")
	s.sslMode = viper.GetString("development.sslmode")
}
