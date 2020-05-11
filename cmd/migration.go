package cmd

import (
	"fmt"
	"os"

	"github.com/kevinsapp/monarch/pkg/sql"
	"github.com/spf13/cobra"
)

const migrationsDir string = "migrations"

func init() {
	generateCmd.AddCommand(migrationCmd)
}

// migrationCmd ...
var migrationCmd = &cobra.Command{
	Use:              "migration",
	Aliases:          []string{"m"},
	PersistentPreRun: mkdirMigrations,
}

// createMigration creates a migration file based on arguments.
func createMigration(fname, sqlt string, data interface{}) (*os.File, error) {
	// Create a migration file.
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Process SQL template
	sql, err := sql.ProcessTmpl(data, sqlt)
	if err != nil {
		return f, err
	}

	// Write SQL to the file.
	_, err = f.WriteString(sql)
	if err != nil {
		return f, err
	}

	return f, err
}

// mkdirMigrations creates a directory called `migrations` in the current
// working directory. If the `migrations` directory already exists,
// mkdirMigations does nothing.
func mkdirMigrations(cmd *cobra.Command, args []string) {
	const (
		dn             = "migrations" // directory name
		fm os.FileMode = 0755         // 0755 Unix file permissions
	)
	err := os.MkdirAll(dn, fm)
	if err != nil {
		fmt.Printf("Error creating %s directory: %s\n", dn, err)
	}
}
