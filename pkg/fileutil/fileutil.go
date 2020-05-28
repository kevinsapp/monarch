package fileutil

import (
	"io/ioutil"
	"os"
)

// CreateAndWriteString creates the named file with mode 0666 (before umask).
// If the file already exists the function truncates it. Then the function
// writes the content of str to the file.
func CreateAndWriteString(path, str string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	return err
}

// MkdirP creates a directory at "path" relative to the current working directory
// and makes parents as necessary. If the directory already exists, MkdirP does nothing.
func MkdirP(path string) error {
	const mode os.FileMode = 0755 // 0755 Unix file permissions
	err := os.MkdirAll(path, mode)
	if err != nil {
		return err
	}

	return err
}

// ReadFileAsString reads in the contents of a text file and returns a string.
func ReadFileAsString(fn string) (string, error) {
	// Read file contents to buffer
	var s string
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return s, err
	}

	// Return file contents as a string
	s = string(b)
	return s, err
}
