package migration

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// Migration ...
type Migration struct {
	sql     string
	version int64
}

// SQL returns SQL.
func (m *Migration) SQL() string {
	return m.sql
}

// SetSQL sets SQL.
func (m *Migration) SetSQL(sql string) {
	m.sql = sql
}

// Version returns version.
func (m *Migration) Version() int64 {
	return m.version
}

// SetVersion sets version.
func (m *Migration) SetVersion(ver int64) {
	m.version = ver
}

// SetFromFile sets fields from a migration file.
func (m *Migration) SetFromFile(path string) error {
	// Extract version from migration file name.
	version, err := ExtractVersionFromFile(path)
	if err != nil {
		return err
	}

	// Read in SQL content from migration file.
	sql, err := FileAsString(path)
	if err != nil {
		return err
	}

	// Set fields.
	m.SetSQL(sql)
	m.SetVersion(version)

	return err
}

// ExtractVersionFromFile extracts version from a migration file name.
func ExtractVersionFromFile(path string) (int64, error) {
	fn := filepath.Base(path)
	fnParts := strings.Split(fn, "_")
	version, err := strconv.ParseInt(fnParts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return version, err
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
