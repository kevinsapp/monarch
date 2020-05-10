package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	migrationCmd.AddCommand(dropCmd)
}

// dropCmd ...
var dropCmd = &cobra.Command{
	Use: "drop",
}
