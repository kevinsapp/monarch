package migration

import (
	"github.com/kevinsapp/monarch/pkg/fileutil"
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
	s, err := fileutil.ReadFileAsString(path)
	if err != nil {
		return err
	}

	// Extract version from migration file name.
	v, err := fileutil.ExtractVersionFromFile(path)
	if err != nil {
		return err
	}

	// Set fields.
	m.SetSQL(s)
	m.SetVersion(v)

	return err
}
