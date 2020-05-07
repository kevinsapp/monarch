package generate

import (
	"github.com/spf13/cobra"
)

func init() {
	CreateCmd.AddCommand(TableCmd)
}

// CreateCmd ...
var CreateCmd = &cobra.Command{
	Use: "create",
}
