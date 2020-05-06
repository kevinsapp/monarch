package cmd

import (
	"github.com/kevinsapp/monarch/cmd/generate"
	"github.com/spf13/cobra"
)

func init() {
	generateCmd.AddCommand(generate.MigrationCmd)
	rootCmd.AddCommand(generateCmd)
}

// Generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
}
