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
	// Read in content from migration file to a buffer.
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Extract version from migration file name.
	v, err := ExtractVersionFromFile(path)
	if err != nil {
		return err
	}

	// Set fields.
	m.SetSQL(string(b))
	m.SetVersion(v)

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
