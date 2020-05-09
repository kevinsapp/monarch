package cmd

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	generateCmd.AddCommand(migrationCmd)
}

// migrationCmd ...
var migrationCmd = &cobra.Command{
	Use:              "migration",
	Aliases:          []string{"m"},
	PersistentPreRun: mkdirMigrations,
}

// createMigration creates a migration file based on arguments.
func createMigration(fname, sqlt string, data interface{}) (*os.File, error) {
	// Create a migration file.
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}

	// Process SQL template
	sql, err := templateAsSQL(data, sqlt)
	if err != nil {
		return f, err
	}

	// Write SQL to the file.
	_, err = f.WriteString(sql)
	if err != nil {
		return f, err
	}

	return f, err
}

// mkdirMigrations creates a directory called `migrations` in the current
// working directory. If the `migrations` directory already exists,
// mkdirMigations does nothing.
func mkdirMigrations(cmd *cobra.Command, args []string) {
	const (
		dn             = "migrations" // directory name
		fm os.FileMode = 0755         // 0755 Unix file permissions
	)
	err := os.MkdirAll(dn, fm)
	if err != nil {
		fmt.Printf("Error creating %s directory: %s\n", dn, err)
	}
}

// templateAsSQL applies a data structure to a SQL template and returns a string.
func templateAsSQL(data interface{}, sqlt string) (string, error) {
	// Initialize a template.
	t := template.New("sql")

	// Parse the template.
	t, err := t.Parse(sqlt)
	if err != nil {
		return "", err
	}

	// Apply the data structure to the template and write to a buffer.
	var tbuf bytes.Buffer
	err = t.Execute(&tbuf, data)
	if err != nil {
		return "", err
	}

	// Get contents of the buffer as a string.
	sql := tbuf.String()

	return sql, err
}
