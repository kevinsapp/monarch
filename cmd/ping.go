package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	dbCmd.AddCommand(pingCmd)
}

// pingCmd ...
var pingCmd = &cobra.Command{
	Use:  "ping",
	RunE: ping,
}

// Ping the database to verify that the server is accessible.
// If ping fails, exit the application with an error.
func ping(cmd *cobra.Command, args []string) error {
	err := db.Ping()
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
		return err
	}

	fmt.Println("Successfully connected to database.")

	return nil
}
