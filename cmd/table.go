package cmd

import (
	"errors"
	"strings"

	"github.com/kevinsapp/monarch/pkg/sqlt"
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
	Short: "Generate a migration file to create a table named [name].",
	RunE:  createTableMigration,
}

// dropTableCmd generates an "up" migration file to drop a table.
var dropTableCmd = &cobra.Command{
	Use:   "table [name]",
	Short: "Generate a migration file to drop a table named [name].",
	Long: `Generate a migration file to drop a table named [name].
	WARNING: This migration is irreversible and any data in the table will be lost
	when the migration is run and the table has been dropped.`,
	RunE: dropTableMigration,
}

// renameTableCmd generates  an "up" migration file to rename a table from [name]
// to [newname] and a companion "down" migration file to rename the table from
// [newname] to [name].
var renameTableCmd = &cobra.Command{
	Use:   "table [name] [newname]",
	Short: "Generate a migration file to rename a table from [name] to [newname].",
	RunE:  renameTableMigration,
}

// createTableMigration creates a migration file to create a table.
func createTableMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 1 {
		return errors.New("requires a name argument")
	}

	// Set table data.
	tableName := args[0]
	t := new(sqlt.Table)
	t.SetName(tableName)

	// If column args are present, parse args and add columns to table.
	if len(args) > 1 {
		// Add columns to table object
		for _, v := range args[1:] {
			nameType := strings.Split(v, ":")

			col := sqlt.Column{}
			col.SetName(nameType[0])
			col.SetType(nameType[1])

			t.AddColumn(col)
		}
	}

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.CreateTableTmpl)
	if err != nil {
		return err
	}

	// Process SQL template for "down" migration.
	downSQL, err := sqlt.ProcessTmpl(t, sqlt.DropTableTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("CreateTable_"+tableName, upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}

// dropTableMigration creates an "up" migration file to drop a table.
func dropTableMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 1 {
		return errors.New("requires a name argument")
	}

	// Set table data.
	tableName := args[0]
	t := new(sqlt.Table)
	t.SetName(tableName)

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.DropTableTmpl)
	if err != nil {
		return err
	}

	// // Process SQL template for "down" migration.
	// downSQL, err := sqlt.ProcessTmpl(t, sqlt.DropTableTmpl)
	// if err != nil {
	// 	return err
	// }

	// Create migration file.
	err = createMigration("DropTable_"+tableName, upSQL, "")
	if err != nil {
		return err
	}

	return err
}

// renameTableMigration generates a migration file to rename a table from
// [name] to [newname].
func renameTableMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply name of an existing table as the first argument,
	// and a new name for that table as the second argument.
	if len(args) < 2 {
		return errors.New("requires two arguments: name and newname")
	}

	// Set timestamp and table data.
	tableName := args[0]
	newName := args[1]
	t := new(sqlt.Table)
	t.SetName(tableName)
	t.SetNewName(newName)

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.RenameTableTmpl)
	if err != nil {
		return err
	}

	// Process SQL template for "down" migration.
	t.SetName(newName)      // swap name and newname
	t.SetNewName(tableName) // swap name and newname
	downSQL, err := sqlt.ProcessTmpl(t, sqlt.RenameTableTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("RenameTable_"+tableName, upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}
