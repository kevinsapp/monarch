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
	Use: "migration",
}

// mkdirMigrations creates a directory called `migrations` in the current
// working directory. If the `migrations` directory already exists,
// mkdirMigations does nothing and returns nil.
func mkdirMigations() error {
	const (
		dn             = "migrations" // directory name
		fm os.FileMode = 0755         // 0755 Unix file permissions
	)
	err := os.MkdirAll(dn, fm)
	if err != nil {
		fmt.Printf("Error creating %s directory: %s\n", dn, err)
		return err
	}

	return err
}
