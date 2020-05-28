package migration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kevinsapp/monarch/pkg/fileutil"
)

const (
	// Delimiter used for parsing migration files.
	migrationDelimiter string = "-- MIGRATION DELIMITER (DO NOT DELETE THIS COMMENT) --"
)

// Migration ...
type Migration struct {
	name           string
	leadingComment string
	upSQL          string
	downSQL        string
	sql            string
	version        int64
}

// Name returns the migration name.
func (m *Migration) Name() string {
	return m.name
}

// SetName sets migration name after converting name string to snake_case.
func (m *Migration) SetName(name string) {
	m.name = strcase.ToSnake(name)
}

// LeadingComment returns the migration leading comment
func (m *Migration) LeadingComment() string {
	return m.leadingComment
}

// SetLeadingComment sets the migration leading comment after sanitizng comment
// and prefixing it with "-- ".
func (m *Migration) SetLeadingComment(comment string) {
	c := strings.TrimSpace(comment) // Trim whistespace
	c = strings.TrimPrefix(c, "--") // Remove any leading SQL comment indicator
	c = strings.TrimSpace(c)        // Trim whitespace again
	c = "-- " + c                   // Add leading SQL comment indicator

	m.leadingComment = c
}

// UpSQL returns SQL for an "up" migration.
func (m *Migration) UpSQL() string {
	return m.upSQL
}

// SetUpSQL sets SQL for an "up" migration.
func (m *Migration) SetUpSQL(sql string) {
	m.upSQL = sql
}

// DownSQL returns SQL for a "down" migration.
func (m *Migration) DownSQL() string {
	return m.downSQL
}

// SetDownSQL sets SQL for a "down" migration.
func (m *Migration) SetDownSQL(sql string) {
	m.downSQL = sql
}

// SQL returns the migration SQL including leading comment, up SQL and down SQL.
func (m *Migration) SQL() string {
	var sql string

	substr := make([]string, 0)
	substr = append(substr, m.upSQL)
	substr = append(substr, migrationDelimiter)
	substr = append(substr, m.downSQL)

	sql = strings.Join(substr, "\n\n")

	return sql
}

// Version returns version.
func (m *Migration) Version() int64 {
	return m.version
}

// SetVersion sets version.
func (m *Migration) SetVersion(ver int64) {
	m.version = ver
}

// ReadFromFile creates a migration file in the directory specified by "dir"
// and writes content to it based on this migration's fields.
func (m *Migration) ReadFromFile(path string) error {
	// Set name
	name := extractNameFromFile(path)
	m.SetName(name)

	// Set version
	version, err := extractVersionFromFile(path)
	if err != nil {
		return err
	}
	m.SetVersion(version)

	// Set upSQL and downSQL
	str, err := fileutil.ReadFileAsString(path)
	parts := strings.Split(str, migrationDelimiter)
	upSQL := strings.TrimSpace(parts[0])
	downSQL := strings.TrimSpace(parts[1])
	m.SetUpSQL(upSQL)
	m.SetDownSQL(downSQL)

	return err
}

// WriteToFile creates a migration file in the directory specified by "dir"
// and writes content to it based on this migration's fields.
func (m *Migration) WriteToFile(dirname string) (string, error) {

	// Generate migration file name.
	fn := fmt.Sprintf("%s/%d_%s.sql", dirname, m.Version(), m.Name())

	// Create migration file.
	err := fileutil.CreateAndWriteString(fn, m.SQL())
	if err != nil {
		return fn, err
	}

	return fn, err
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
		v, err := extractVersionFromFile(n)
		if err != nil {
			return migrations, err
		}

		// Select only the migration files with a version greater than
		// schemaVersion
		if v > version {
			err = m.ReadFromFile(dirname + "/" + n)
			if err != nil {
				return migrations, err
			}
			migrations = append(migrations, m)
			fmt.Printf("Staged %q migration version: %d\n", "up", m.Version())
		}
	}

	return migrations, err
}

// extractNameFromFile extracts name from a migration filename.
func extractNameFromFile(path string) string {
	fn := filepath.Base(path)
	fnParts := strings.Split(fn, "_")
	name := strings.Join(fnParts[1:], "_")
	name = strings.TrimSuffix(name, ".sql")

	return name
}

// extractVersionFromFile extracts version from a migration filename.
func extractVersionFromFile(path string) (int64, error) {
	fn := filepath.Base(path)
	fnParts := strings.Split(fn, "_")
	ver, err := strconv.ParseInt(fnParts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return ver, err
}
