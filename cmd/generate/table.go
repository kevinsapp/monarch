package generate

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// table ...
type table struct {
	Name string
}

// TableCmd generates an "up" migration file to create a table and a "down" migration
// file to drop that table.
var TableCmd = &cobra.Command{
	Use:   "table [name]",
	Short: "Generate migration files to create a table named [name].",
	Long: `Generate an "up" migration file to create a table named [name]
and a companion "down" migration file to drop that table.`,
	RunE: createTableMigrations,
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

	// Create an "up" migration file.
	// _, err := upTableMigration(timestamp, name)
	fn := fmt.Sprintf("migrations/%d_create_table_%s_up.sql", timestamp, td.Name)
	_, err := createMigration(fn, sqltCreateTable, td)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	// _, err = dnTableMigration(timestamp, name)
	fn = fmt.Sprintf("migrations/%d_create_table_%s_down.sql", timestamp, td.Name)
	_, err = createMigration(fn, sqltDropTable, td)
	if err != nil {
		return err
	}

	return err
}
