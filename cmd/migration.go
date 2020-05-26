package cmd

import (
	"fmt"
	"time"

	"github.com/kevinsapp/monarch/pkg/fileutil"
	"github.com/kevinsapp/monarch/pkg/migration"
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
func createMigration(name, upSQL, downSQL string) error {
	// Configure a migration object.
	m := new(migration.Migration)
	m.SetName(name)
	m.SetUpSQL(upSQL)
	m.SetDownSQL(downSQL)
	m.SetVersion(time.Now().UnixNano())

	// Write migration file.
	_, err := m.WriteToFile(migrationsDir)
	if err != nil {
		return err
	}

	return err
}

// mkdirMigrations creates a directory called `migrations` in the current
// working directory. If the `migrations` directory already exists,
// mkdirMigations does nothing.
func mkdirMigrations(cmd *cobra.Command, args []string) {
	err := fileutil.MkdirP(migrationsDir)
	if err != nil {
		fmt.Printf("Error creating directory %q: %s\n", migrationsDir, err)
	}
}
