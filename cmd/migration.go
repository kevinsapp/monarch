package cmd

import (
	"fmt"

	"github.com/kevinsapp/monarch/pkg/fileutil"
	"github.com/kevinsapp/monarch/pkg/sqlt"
	"github.com/spf13/cobra"
)

const migrationsDir string = "migrations"

func init() {
	generateCmd.AddCommand(migrationCmd)
	migrationCmd.AddCommand(addCmd)
	migrationCmd.AddCommand(createCmd)
	migrationCmd.AddCommand(dropCmd)
	migrationCmd.AddCommand(renameCmd)
}

// migrationCmd ...
var migrationCmd = &cobra.Command{
	Use:              "migration",
	Aliases:          []string{"m"},
	PersistentPreRun: mkdirMigrations,
}

// addCmd ...
var addCmd = &cobra.Command{
	Use: "add",
}

// createCmd ...
var createCmd = &cobra.Command{
	Use: "create",
}

// dropCmd ...
var dropCmd = &cobra.Command{
	Use: "drop",
}

// renameCmd ...
var renameCmd = &cobra.Command{
	Use: "rename",
}

// createMigration creates a migration file based on arguments.
func createMigration(path, tmpl string, data interface{}) error {
	// Process SQL template
	sql, err := sqlt.ProcessTmpl(data, tmpl)
	if err != nil {
		return err
	}

	// Create migration file.
	err = fileutil.CreateAndWriteString(path, sql)
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s.\n", path)

	return err
}

// mkdirMigrations creates a directory called `migrations` in the current
// working directory. If the `migrations` directory already exists,
// mkdirMigations does nothing.
func mkdirMigrations(cmd *cobra.Command, args []string) {
	path := "migrations"
	err := fileutil.MkdirP(path)
	if err != nil {
		fmt.Printf("Error creating directory %q: %s\n", path, err)
	}
}
