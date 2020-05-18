package cmd

import (
	"fmt"
	"log"
	"time"

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
// If ping fails, log and return an error.
func ping(cmd *cobra.Command, args []string) error {
	now := time.Now()
	err := db.Ping()
	since := time.Since(now)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return err
	}

	fmt.Printf("Database connection OK. Server replied in %s.\n", since)

	return nil
}
