package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kevinsapp/monarch/pkg/sql"
	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(createTableCmd)
	dropCmd.AddCommand(dropTableCmd)
	renameCmd.AddCommand(renameTableCmd)
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

// renameTableCmd generates  an "up" migration file to rename a table from [name]
// to [newname] and a companion "down" migration file to rename the table from
// [newname] to [name].
var renameTableCmd = &cobra.Command{
	Use:   "table [oldname] [newname]",
	Short: "Generate migration files to rename a table from [name] to [newname].",
	Long: `Generate an "up" migration file to rename a table from [name] to [newname]
and a companion "down" migration file to rename the table from [newname] to [name].`,
	RunE: renameTableMigrations,
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
	td := sql.Table{}
	td.SetName(args[0])

	// If column args are present, parse args and add columns to table.
	if len(args) > 1 {
		// Add columns to table object
		for _, v := range args[1:] {
			nameType := strings.Split(v, ":")

			col := sql.Column{}
			col.SetName(nameType[0])
			col.SetType(nameType[1])

			td.AddColumn(col)
		}
	}

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_create_table_%s_up.sql", timestamp, td.Name())
	_, err := createMigration(fn, sql.CreateTableTmpl, &td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_create_table_%s_down.sql", timestamp, td.Name())
	_, err = createMigration(fn, sql.DropTableTmpl, &td)
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
	td := sql.Table{}
	td.SetName(args[0])

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_drop_table_%s_up.sql", timestamp, td.Name())
	_, err := createMigration(fn, sql.DropTableTmpl, &td)
	if err != nil {
		return err
	}

	return err
}

// renameTableMigrations generates an "up" migration file to rename a table from
// [name] to [newname] and a companion "down" migration file to rename the table
// from [newname] to [name].
func renameTableMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply name of an existing table as the first argument, and a
	// new name for that table as the second argument.
	if len(args) < 2 {
		return errors.New("requires two arguments: name and newname")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := sql.Table{}
	td.SetName(args[0])
	td.SetNewName(args[1])

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_rename_table_%s_up.sql", timestamp, td.Name())
	_, err := createMigration(fn, sql.RenameTableTmpl, &td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_rename_table_%s_down.sql", timestamp, td.Name())
	td.SetName(args[1])    // swap name and newname
	td.SetNewName(args[0]) // swap name and newname
	_, err = createMigration(fn, sql.RenameTableTmpl, &td)
	if err != nil {
		return err
	}

	return err
}
