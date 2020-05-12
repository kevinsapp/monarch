package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	migrationCmd.AddCommand(addCmd)
}

// addCmd ...
var addCmd = &cobra.Command{
	Use: "add",
}
