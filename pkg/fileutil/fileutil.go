package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CreateAndWriteString creates or truncates the named file. If the file
// already exists, it is truncated. If the file does not exist, it is created
// with mode 0666 (before umask). Then it writes the content of str to the file.
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

// ExtractVersionFromFile extracts version from a migration file name.
func ExtractVersionFromFile(path string) (int64, error) {
	fn := filepath.Base(path)
	fnParts := strings.Split(fn, "_")
	ver, err := strconv.ParseInt(fnParts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return ver, err
}

// MkdirP creates a directory at "path" relative to the current working directory
// and makes parents as necessary. If the directory already exists, Mkdir does nothing.
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
