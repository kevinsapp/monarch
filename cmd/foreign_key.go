package cmd

import (
	"errors"

	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addForeignKeyCmd)
	dropCmd.AddCommand(dropForeignKeyCmd)
}

// addForeignKeyCmd generates a migration file to add a foreign key column and constraint to a table.
var addForeignKeyCmd = &cobra.Command{
	Aliases: []string{"fk"},
	Use:     "foreignkey [childTableName] [parentTableName]",
	Short:   "Generate a migration file to add a foreign key to a table.",
	RunE:    addForeignKeyMigration,
}

// dropForeignKeyCmd generates a migration file to drop a foreign key column and constraint from a table.
var dropForeignKeyCmd = &cobra.Command{
	Aliases: []string{"fk"},
	Use:     "foreignkey [childTableName] [parentTableName]",
	Short:   "Generate a migration file to drop a foreign key from a table.",
	RunE:    dropForeignKeyMigration,
}

// addForeignKeyMigration creates a migration file to add a foreign key to a table.
func addForeignKeyMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a referencing table name (child table) as the first argument and a
	// referenced table name (parent table) as the second argument.
	if len(args) < 2 {
		return errors.New("requires childTableName and parentTableName arguments")
	}

	// Set foreign key data.
	fk := new(sqlt.ForeignKey)
	fk.SetReferencingTableName(args[0])
	fk.SetReferencedTableName(args[1])

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(fk, sqlt.AddForeignKeyTmpl)
	if err != nil {
		return err
	}

	// Process SQL template for "down" migration.
	downSQL, err := sqlt.ProcessTmpl(fk, sqlt.DropForeignKeyTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("AddForeignKeyTo_"+fk.ReferencingTableName(), upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}

// dropForeignKeyMigration creates a migration file to add a foreign key to a table.
func dropForeignKeyMigration(cmd *cobra.Command, args []string) error {
	// Caller should supply a referencing table name (child table) as the first argument and a
	// referenced table name (parent table) as the second argument.
	if len(args) < 2 {
		return errors.New("requires childTableName and parentTableName arguments")
	}

	// Set foreign key data.
	fk := new(sqlt.ForeignKey)
	fk.SetReferencingTableName(args[0])
	fk.SetReferencedTableName(args[1])

	// Process SQL template for "up" migration.
	upSQL, err := sqlt.ProcessTmpl(fk, sqlt.DropForeignKeyTmpl)
	if err != nil {
		return err
	}

	// Process SQL template for "down" migration.
	downSQL, err := sqlt.ProcessTmpl(fk, sqlt.AddForeignKeyTmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = createMigration("DropForeignKeyFrom_"+fk.ReferencingTableName(), upSQL, downSQL)
	if err != nil {
		return err
	}

	return err
}
