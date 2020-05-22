package sqlt

import (
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

// SetFromFile sets attributes from a migration file.
func (m *Migration) SetFromFile(filename string) error {
	fn := filename

	// Extract migration version from filename.
	fnParts := strings.Split(fn, "_")
	ver, err := strconv.ParseInt(fnParts[0], 10, 64)
	if err != nil {
		return err
	}

	// Read in SQL content from migration file.
	sql, err := FileAsString(fn)
	if err != nil {
		return err
	}

	// Set attributes.
	m.version = ver
	m.sql = sql

	return err
}
