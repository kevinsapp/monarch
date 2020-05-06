package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

// Generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a migration file",
	Long:  "Generate a migration file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating migration file...")
		fmt.Printf("File: %s\n", args[0])
	},
}
