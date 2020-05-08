package generate

import (
	"testing"
)

// Unit test dropTableSQL()
func TestDropTableSQL(t *testing.T) {
	exp := `-- Down migration for users table

DROP TABLE users;
`
	// Run dropTableSQL()
	act, err := dropTableSQL("users")
	if err != nil {
		t.Fatal(err)
	}

	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}
