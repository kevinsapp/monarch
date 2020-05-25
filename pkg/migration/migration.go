package migration

import (
	"fmt"
	"io/ioutil"
	"strings"

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

// LoadAllLaterThan ...
func LoadAllLaterThan(version int64, dirname string) ([]Migration, error) {
	migrations := make([]Migration, 0)

	// Get the list of files in the directory specificed by path.
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return migrations, err
	}

	var m Migration
	for _, f := range files {
		n := f.Name()
		v, err := fileutil.ExtractVersionFromFile(n)
		if err != nil {
			return migrations, err
		}

		// Select only the migration files with:
		// a) a suffix of "up.sql", and
		// b) a version greater than schemaVersion
		if v > version && strings.HasSuffix(n, "up.sql") {
			m.SetFromFile(dirname + "/" + n)
			migrations = append(migrations, m)
			fmt.Printf("Staged %q migration version: %d\n", "up", m.Version())
		}
	}

	return migrations, err
}
