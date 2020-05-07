package generate

import (
	"github.com/spf13/cobra"
)

func init() {
	CreateCmd.AddCommand(TableCmd)
}

// CreateCmd generates a migration file.
var CreateCmd = &cobra.Command{
	Use: "create",
}
