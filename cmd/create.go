package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	migrationCmd.AddCommand(createCmd)
}

// createCmd ...
var createCmd = &cobra.Command{
	Use: "create",
}
