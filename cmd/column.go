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
	addColumnCmd.Flags().String("table", "TABLE_NAME", "table to add column to")
}

// addColumnCmd generates an "up" migration file to add a column to
// a table and a "down" migration file to remove that column.
var addColumnCmd = &cobra.Command{
	Use:   "column [ [name:type] ... ]",
	Short: "Generate migration files to add a column named [name] with type [type].",
	Long: `Generate an "up" migration file to add a column named [name] with type
[type] and a companion "down" migration file to remove that column.`,
	RunE: addColumnMigrations,
}

// dropColumnCmd generates an "up" migration file to remove a column
// from a table.
var dropColumnCmd = &cobra.Command{
	Use:   "column [ [name] ... ]",
	Short: "Generate migration files to add a column named [name].",
	Long:  `Generate an "up" migration file to add a column named [name].`,
	RunE:  addColumnMigrations,
}

// addColumnMigrations creates an "up" migration file to add a column to
// a table and a "down" migration file to remove that column.
func addColumnMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 1 {
		return errors.New("requires a name:type argument")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := sql.Table{}
	td.Name, _ = cmd.Flags().GetString("table")

	// Add columns to table object
	for _, v := range args {
		nameType := strings.Split(v, ":")

		col := sql.Column{}
		col.Name = strcase.ToSnake(nameType[0])
		col.Type = strcase.ToSnake(nameType[1]) // TODO: we we need snake_case for this string?
		col.Type = strings.ToUpper(col.Type)

		td.Columns = append(td.Columns, col)
	}

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_add_columns_to_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sql.AddColumnTmpl, td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	fn = fmt.Sprintf("migrations/%d_add_columns_to_%s_down.sql", timestamp, td.Name)
	_, err = createMigration(fn, sql.DropColumnTmpl, td)
	if err != nil {
		return err
	}

	return err
}

// dropColumnMigrations creates an "up" migration file to drop a column
// from a table.
func dropColumnMigrations(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument.
	if len(args) < 1 {
		return errors.New("requires a name argument")
	}

	// Set timestamp and table data.
	timestamp := time.Now().UnixNano()
	td := sql.Table{}
	td.Name, _ = cmd.Flags().GetString("table")

	// Drop columns from table object
	for _, v := range args {
		col := sql.Column{}
		col.Name = strcase.ToSnake(v)

		td.Columns = append(td.Columns, col)
	}

	// Create an "up" migration file.
	fn := fmt.Sprintf("migrations/%d_drop_columns_from_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sql.DropColumnTmpl, td)
	if err != nil {
		return err
	}

	return err
}
