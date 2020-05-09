package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

// Generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
}
