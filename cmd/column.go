package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/kevinsapp/monarch/pkg/sql"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addColumnCmd)
	dropCmd.AddCommand(dropColumnCmd)
	renameCmd.AddCommand(renameColumnCmd)
}

// addColumnCmd generates an "up" migration file to add a column to
// a table and a "down" migration file to remove that column.
var addColumnCmd = &cobra.Command{
	Use:   "column [tableName] [ [colName:type] ... ]",
	Short: "Generate migration files to add a column named [colName] with type [type].",
	Long: `Generate an "up" migration file to add a column named [colName] with type
[type] and a companion "down" migration file to remove that column.`,
	RunE: addColumnMigrations,
}

// dropColumnCmd generates an "up" migration file to remove a column
// from a table.
var dropColumnCmd = &cobra.Command{
	Use:   "column [ [name] ... ]",
	Short: "Generate migration files to drop a column named [colName].",
	Long:  `Generate an "up" migration file to drop a column named [colName].`,
	RunE:  dropColumnMigrations,
}

// renameColumnMigrations generates an "up" migration file to rename a column from
// [name] to [newname] and a companion "down" migration file to rename the column
// from [newname] to [name].
var renameColumnCmd = &cobra.Command{
	Use:   "column [tableName] [ [name:newName] ... ]",
	Short: "Generate migration files to rename a column from [name] to [newName].",
	Long: `Generate an "up" migration file to rename a column from [name] to [newName]
and a companion "down" migration file to rename the column from [newName] to [name].`,
	RunE: renameColumnMigrations,
}

// addColumnMigrations creates an "up" migration file to add a column to
// a table and a "down" migration file to remove that column.
func addColumnMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 2 {
		return errors.New("requires tableName and colName:type arguments")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := sql.Table{}
	td.SetName(args[0])

	// Add columns to table object
	for _, v := range args[1:] {
		nameType := strings.Split(v, ":")

		col := sql.Column{}
		col.Name = strcase.ToSnake(nameType[0])
		col.Type = strcase.ToSnake(nameType[1]) // TODO: we we need snake_case for this string?
		col.Type = strings.ToUpper(col.Type)

		td.AddColumn(col)
	}

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_add_columns_to_%s_up.sql", timestamp, td.Name())
	_, err := createMigration(fn, sql.AddColumnTmpl, &td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_add_columns_to_%s_down.sql", timestamp, td.Name())
	_, err = createMigration(fn, sql.DropColumnTmpl, &td)
	if err != nil {
		return err
	}

	return err
}

// dropColumnMigrations creates an "up" migration file to drop a column
// from a table.
func dropColumnMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 2 {
		return errors.New("requires tableName and colName arguments")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := sql.Table{}
	td.SetName(args[0])

	// Drop columns from table object
	for _, v := range args[1:] {
		col := sql.Column{}
		col.Name = strcase.ToSnake(v)

		td.AddColumn(col)
	}

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_drop_columns_from_%s_up.sql", timestamp, td.Name())
	_, err := createMigration(fn, sql.DropColumnTmpl, &td)
	if err != nil {
		return err
	}

	return err
}

// renameColumnMigrations creates an "up" migration file to rename a column from
// [name] to [newName] a "down" migration file to rename that column from [newName]
// to [name].
func renameColumnMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 2 {
		return errors.New("requires tableName and colName:newName arguments")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := sql.Table{}
	td.SetName(args[0])

	// Add columns to table object
	for _, v := range args[1:] {
		names := strings.Split(v, ":")

		col := sql.Column{}
		col.Name = strcase.ToSnake(names[0])
		col.NewName = strcase.ToSnake(names[1])

		td.AddColumn(col)
	}

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_rename_columns_in_%s_up.sql", timestamp, td.Name())
	_, err := createMigration(fn, sql.RenameColumnTmpl, &td)
	if err != nil {
		return err
	}

	// Add columns to table object
	td.SetColumns([]sql.Column{}) // Reinitalize columns
	for _, v := range args[1:] {
		names := strings.Split(v, ":")

		col := sql.Column{}
		col.Name = strcase.ToSnake(names[1])
		col.NewName = strcase.ToSnake(names[0])

		td.AddColumn(col)
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_rename_columns_in_%s_down.sql", timestamp, td.Name())
	_, err = createMigration(fn, sql.RenameColumnTmpl, &td)
	if err != nil {
		return err
	}

	return err
}
