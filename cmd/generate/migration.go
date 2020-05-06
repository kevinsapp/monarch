package generate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// MigrationCmd generates a migration file.
var MigrationCmd = &cobra.Command{
	Use:   "migration [name]",
	Short: "Generate a migration file.",
	Long:  "Generate a migration file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating migration file...")

		// Make a "migrations" directory in the current working directory,
		// unless it already exists.
		dir := "migrations"       // directory name
		var fm os.FileMode = 0755 // 0755 Unix file permissions
		err := os.MkdirAll(dir, fm)
		if err != nil {
			fmt.Printf("Error creating %s directory: %s\n", dir, err)
			return err
		}

		return err
	},
}
