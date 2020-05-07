package generate

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// TableCmd generates migration files to create (and drop) a table.
var TableCmd = &cobra.Command{
	Use:   "table [name]",
	Short: "Generate a migration file to create a table named [name].",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating migration files...")

		// Make directory named `migrations`
		err := mkdirMigations()
		if err != nil {
			return err
		}

		ts := time.Now().UnixNano()
		fn := fmt.Sprintf("migrations/%d_create_table_users_up.sql", ts)
		fmt.Println(fn)
		_, err = os.Create(fn)
		if err != nil {
			fmt.Printf("Error creating %s file: %s\n", fn, err)
			return err
		}

		return err
	},
}
