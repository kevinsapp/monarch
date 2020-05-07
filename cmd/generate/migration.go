package generate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	MigrationCmd.AddCommand(CreateCmd)
}

// MigrationCmd ...
var MigrationCmd = &cobra.Command{
	Use:              "migration",
	Aliases:          []string{"m"},
	PersistentPreRun: mkdirMigrations,
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
