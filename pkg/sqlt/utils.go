package sqlt

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

// ProcessTmpl applies a data structure to a SQL template and returns a string.
func ProcessTmpl(data interface{}, sqlt string) (string, error) {
	// Initialize a template.
	var s string
	t := template.New("t")

	// Parse the template.
	t, err := t.Parse(sqlt)
	if err != nil {
		return s, err
	}

	// Apply the data structure to the template and write the result to a buffer.
	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		return s, err
	}

	// Get contents of the byte buffer as a string.
	s = b.String()

	return s, err
}

// FileAsString reads in the contents of a SQL file and returns a string.
func FileAsString(fn string) (string, error) {
	// Read file contents to buffer
	var s string
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return s, err
	}

	// Return file contents as a string
	return string(b), err
}
