package fileutil

import (
	"os"
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

// MkdirP creates a directory at path relative to the current working directory.
// If the directory already exists, Mkdir does nothing. Makes parents as needed.
func MkdirP(path string) error {
	const mode os.FileMode = 0755 // 0755 Unix file permissions
	err := os.MkdirAll(path, mode)
	if err != nil {
		return err
	}

	return err
}
