package cmd

import (
	"errors"
	"fmt"

	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(createIndexCmd)
}

// createIndexCmd generates a migration file to create an index on a table column.
var createIndexCmd = &cobra.Command{
	Use:   "index [tablename] [columnname]",
	Short: "Generate a migration file to create an index on a table column.",
	RunE:  createIndexMigration,
}

// createIndexMigration creates a migration file to create an index on a table column.
func createIndexMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a table name as the first argument and a column name
	// as the second argument.
	if len(args) < 2 {
		return errors.New("requires a tablename argument followed by a columnname argument")
	}

	// Set index data.
	idx := new(sqlt.Index)
	idx.SetTableName(args[0])
	idx.SetColumnName(args[1])

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(idx, sqlt.CreateDefaultIndexTmpl)
	if err != nil {
		return err
	}

	// Process SQL template for "down" migration.
	downSQL, err := sqlt.ProcessTmpl(idx, sqlt.DropIndexTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	migrationName := fmt.Sprintf("CreateIndexOn_%s_%s", idx.TableName(), idx.ColumnName())
	err = createMigration(migrationName, upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}
