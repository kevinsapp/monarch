package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// table ...
type table struct {
	Name string
}

func init() {
	createCmd.AddCommand(createTableCmd)
	dropCmd.AddCommand(dropTableCmd)
}

// createTableCmd generates an "up" migration file to create a table and a "down" migration
// file to drop that table.
var createTableCmd = &cobra.Command{
	Use:   "table [name]",
	Short: "Generate migration files to create a table named [name].",
	Long: `Generate an "up" migration file to create a table named [name]
and a companion "down" migration file to drop that table.`,
	RunE: createTableMigrations,
}

// dropTableCmd generates an "up" migration file to drop a table.
var dropTableCmd = &cobra.Command{
	Use:   "table [name]",
	Short: "Generate an migration file to drop a table named [name].",
	Long: `Generate an "up" migration file to drop a table named [name]
This migration is irreversible and any data in the table will be lost
when the migration is run and the table has been dropped.`,
	RunE: dropTableMigrations,
}

// createTableMigrations creates an "up" migration file to create a table and
// a "down" migration file to drop that table.
func createTableMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 1 {
		return errors.New("requires a name argument")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := table{args[0]}
	td.Name = strcase.ToSnake(td.Name)

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_create_table_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sqltCreateTable, td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_create_table_%s_down.sql", timestamp, td.Name)
	_, err = createMigration(fn, sqltDropTable, td)
	if err != nil {
		return err
	}

	return err
}

// dropTableMigrations creates an "up" migration file to drop a table.
func dropTableMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 1 {
		return errors.New("requires a name argument")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := table{args[0]}
	td.Name = strcase.ToSnake(td.Name)

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_drop_table_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sqltDropTable, td)
	if err != nil {
		return err
	}

	return err
}
