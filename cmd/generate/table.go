package generate

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

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

	// Set timestamp and table name.
	timestamp := time.Now().UnixNano()
	name := args[0]

	// Create an "up" migration file.
	_, err := upTableMigration(timestamp, name)
	if err != nil {
		return err
	}

	// Create a "down" migration file.
	_, err = dnTableMigration(timestamp, name)
	if err != nil {
		return err
	}

	return err
}

// upTableMigration creates an "up" migration file to create a table.
func upTableMigration(timestamp int64, name string) (*os.File, error) {
	// Format a filename for the "up" migration.
	fn := fmt.Sprintf("migrations/%d_create_table_%s_up.sql", timestamp, name)

	// Create a file for the "up" migration.
	f, err := os.Create(fn)
	if err != nil {
		return nil, err
	}

	// Generate SQL to create a table with name [name].
	sql, err := createTableSQL(name)
	if err != nil {
		return nil, err
	}

	// Write SQL to file.
	_, err = f.WriteString(sql)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// dnTableMigration creates a "downb" migration file to drop a table.
func dnTableMigration(timestamp int64, name string) (*os.File, error) {
	// Format a filename for the "down" migration.
	fn := fmt.Sprintf("migrations/%d_create_table_%s_down.sql", timestamp, name)

	// Create a file for the "down" migration.
	f, err := os.Create(fn)
	if err != nil {
		return nil, err
	}

	// Write SQL to the file.
	sql := dropTableSQL(name)
	_, err = f.WriteString(sql)
	if err != nil {
		return nil, err
	}

	return f, err

}

// createTableSQL returns a string of SQL to create a table.
func createTableSQL(name string) (string, error) {
	// Define a SQL template.
	sql := `--Up migration for {{.Name}} table

CREATE TABLE {{.Name}} (
	id uuid DEFAULT gen_random_uuid() NOT NULL,

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT {{.Name}}_pkey PRIMARY KEY (id)
);
`
	// Define a data structure to apply to the SQL template.
	data := struct {
		Name string
	}{
		Name: name,
	}

	sql, err := templateAsSQL(data, sql)

	return sql, err
}

// dropTableSQL returns a string of SQL to drop a table.
func dropTableSQL(name string) string {
	sql := `-- Down migration for %s table

DROP TABLE %s;
`

	sql = fmt.Sprintf(sql, name, name)

	return sql
}
