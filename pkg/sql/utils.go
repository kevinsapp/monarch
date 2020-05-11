package sql

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

// ProcessTmpl applies a data structure to a SQL template and returns a string.
func ProcessTmpl(data interface{}, sqlt string) (string, error) {
	// Initialize a template.
	t := template.New("sql")

	// Parse the template.
	t, err := t.Parse(sqlt)
	if err != nil {
		return "", err
	}

	// Apply the data structure to the template and write the result to a buffer.
	var tbuf bytes.Buffer
	err = t.Execute(&tbuf, data)
	if err != nil {
		return "", err
	}

	// Get contents of the buffer as a string.
	sql := tbuf.String()

	return sql, err
}

// FileAsString reads in the contents of a SQL file and returns a string.
func FileAsString(fn string) (string, error) {
	// Open file with name fn
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close() // Do cleanup

	// Read file contents to buffer
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	// Return file contents as a string
	return string(b), nil
}
