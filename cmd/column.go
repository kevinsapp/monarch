package cmd

import (
	"errors"
	"strings"

	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addColumnCmd)
	dropCmd.AddCommand(dropColumnCmd)
	recastCmd.AddCommand(recastColumnCmd)
	renameCmd.AddCommand(renameColumnCmd)
}

// addColumnCmd generates a migration file to add a column to a table.
var addColumnCmd = &cobra.Command{
	Use:   "column [tableName] [ [colName:type] ... ]",
	Short: "Generate a migration file to add a column named [colName] with type [type].",
	RunE:  addColumnMigration,
}

// dropColumnCmd generates a migration file to remove a column from a table.
var dropColumnCmd = &cobra.Command{
	Use:   "column [ [name] ... ]",
	Short: "Generate a migration file to drop a column named [colName].",
	RunE:  dropColumnMigration,
}

// recastColumnCmd generates a migration file to change a column's data type.
var recastColumnCmd = &cobra.Command{
	Use:   "column [tableName] [ [name:newType] ... ]",
	Short: "Generate a migration file to change column's data type to [newType].",
	Long: `Generate a migration file to change column's data type to [newType]. This command will
	generate a simple ALTER COLUMN statment to change the type. You may need to add a USING clause
	if you require more complex type conversions. See the PostgreSQL documentation for details.`,
	RunE: recastColumnMigration,
}

// renameColumnCmd generates a migration file to rename a column from [name] to [newname].
var renameColumnCmd = &cobra.Command{
	Use:   "column [tableName] [ [name:newName] ... ]",
	Short: "Generate a migration file to rename a column from [name] to [newName].",
	RunE:  renameColumnMigration,
}

// addColumnMigration creates a migration file to add a column to a table.
func addColumnMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument and a column
	// name:type pair as the second argument.
	if len(args) < 2 {
		return errors.New("requires tableName and colName:type arguments")
	}

	// Set table data.
	tableName := args[0]
	t := new(sqlt.Table)
	t.SetName(tableName)

	// Add columns to table object
	for _, v := range args[1:] {
		nameType := strings.Split(v, ":")

		col := sqlt.Column{}
		col.SetName(nameType[0])
		col.SetType(nameType[1])

		t.AddColumn(col)
	}

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.AddColumnTmpl)
	if err != nil {
		return err
	}

	// Process SQL template for "down" migration.
	downSQL, err := sqlt.ProcessTmpl(t, sqlt.DropColumnTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("AddColumnsTo_"+tableName, upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}

// dropColumnMigration creates a migration file to drop a column from a table.
func dropColumnMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 2 {
		return errors.New("requires tableName and colName arguments")
	}

	// Set table data.
	tableName := args[0]
	t := new(sqlt.Table)
	t.SetName(tableName)

	// Drop columns from table object
	for _, v := range args[1:] {
		col := sqlt.Column{}
		col.SetName(v)

		t.AddColumn(col)
	}

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.DropColumnTmpl)
	if err != nil {
		return err
	}

	// // Process SQL template for "down" migration.
	// downSQL, err := sqlt.ProcessTmpl(t, sqlt.DropColumnTmpl)
	// if err != nil {
	// 	return err
	// }

	// Create migration file.
	err = createMigration("DropColumnsFrom_"+tableName, upSQL, "")
	if err != nil {
		return err
	}

	return err
}

// recastColumnMigration creates a migration file to rename a column from [name] to [newName].
func recastColumnMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 2 {
		return errors.New("requires tableName abd columnName:newType arguments")
	}

	tableName := args[0]
	t := new(sqlt.Table)
	t.SetName(tableName)

	// Add columns to table object
	for _, v := range args[1:] {
		nameType := strings.Split(v, ":")

		col := sqlt.Column{}
		col.SetName(nameType[0])
		col.SetType(nameType[1])

		t.AddColumn(col)
	}

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.RecastColumnTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("RecastColumnsIn_"+tableName, upSQL, "")
	if err != nil {
		return err
	}

	return err
}

// renameColumnMigration creates a migration file to rename a column from [name] to [newName].
func renameColumnMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 2 {
		return errors.New("requires tableName and colName:newName arguments")
	}

	// Set table data.
	tableName := args[0]
	t := new(sqlt.Table)
	t.SetName(tableName)

	// Add columns to table object
	for _, v := range args[1:] {
		names := strings.Split(v, ":")

		col := sqlt.Column{}
		col.SetName(names[0])
		col.SetNewName(names[1])

		t.AddColumn(col)
	}

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(t, sqlt.RenameColumnTmpl)
	if err != nil {
		return err
	}

	// Zero columns
	t.SetColumns([]sqlt.Column{})

	// Add columns to table object, but now reverse the names.
	for _, v := range args[1:] {
		names := strings.Split(v, ":")

		col := sqlt.Column{}
		col.SetName(names[1])
		col.SetNewName(names[0])

		t.AddColumn(col)
	}

	// Process SQL template for "down" migration.
	downSQL, err := sqlt.ProcessTmpl(t, sqlt.RenameColumnTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("RenameColumnsIn_"+tableName, upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}
