package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	MigrationCmd.AddCommand(CreateCmd)
}

// CreateCmd ...
var CreateCmd = &cobra.Command{
	Use: "create",
}
