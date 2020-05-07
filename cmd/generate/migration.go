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
		var (
			dn string      = "migrations" // directory name
			fm os.FileMode = 0755         // 0755 Unix file permissions
		)
		err := os.MkdirAll(dn, fm)
		if err != nil {
			fmt.Printf("Error creating %s directory: %s\n", dn, err)
			return err
		}

		// Create a new migration file in the migrations directory.
		// fname := "test.sql"
		// f, err := os.Create(fname)
		// if err != nil {
		// 	fmt.Printf("Error creating %s file: %s\n", fname, err)
		// 	return err
		// }

		return err
	},
}
