package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	migrationCmd.AddCommand(renameCmd)
}

// renameCmd ...
var renameCmd = &cobra.Command{
	Use: "rename",
}
