package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

// generateCmd ...
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
}
