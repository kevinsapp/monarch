package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/iancoleman/strcase"
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
	td.Name = strcase.ToSnake(args[0])

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_create_table_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sql.CreateTableTmpl, td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_create_table_%s_down.sql", timestamp, td.Name)
	_, err = createMigration(fn, sql.DropTableTmpl, td)
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
	td.Name = strcase.ToSnake(args[0])

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_drop_table_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sql.DropTableTmpl, td)
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
	name := strcase.ToSnake(args[0])
	newName := strcase.ToSnake(args[1])
	td := sql.Table{}

	// Create an "up" migration file.
	td.Name = name       // current name of table
	td.NewName = newName // new name of table
	fn := fmt.Sprintf("migrations/%d_rename_table_%s_up.sql", timestamp, name)
	_, err := createMigration(fn, sql.RenameTableTmpl, td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	td.Name = newName // current name of table after "up" migration
	td.NewName = name // old name of table
	fn = fmt.Sprintf("migrations/%d_rename_table_%s_down.sql", timestamp, name)
	_, err = createMigration(fn, sql.RenameTableTmpl, td)
	if err != nil {
		return err
	}

	return err
}
