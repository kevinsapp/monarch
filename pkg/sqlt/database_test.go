package sqlt

import (
	"testing"
)

// Unit test Database.Name()
func TestDatabaseName(t *testing.T) {
	db := Database{}
	db.name = "monarch_test"

	exp := db.name
	act := db.Name() // Run Database.Name()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Database.SetName()
func TestDatabaseSetName(t *testing.T) {
	cases := [][]string{
		{"test", "test"},
		{"test_case", "test_case"},
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test", "test"},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
	}
	for _, c := range cases {
		db := Database{}
		db.SetName(c[0]) // Run Database.SetName()
		exp := c[1]
		act := db.name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Database.NewName()
func TestDatabaseNewName(t *testing.T) {
	db := Database{}
	db.newName = "test_table"

	exp := db.newName
	act := db.NewName() // Run Database.NewName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Database.SetNewName()
func TestDatabaseSetNewName(t *testing.T) {
	cases := [][]string{
		{"test", "test"},
		{"test_case", "test_case"},
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test", "test"},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
	}
	for _, c := range cases {
		db := Database{}
		db.SetNewName(c[0]) // Run Database.SetNewName()
		exp := c[1]
		act := db.newName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

func TestDatabaseOwner(t *testing.T) {
	db := Database{}
	db.owner = "postgres"

	exp := db.owner
	act := db.Owner() // Run Database.Owner()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

func TestDatabaseSetOwner(t *testing.T) {
	o := "postgres"

	db := Database{}
	db.SetOwner(o) // Run Database.SetOwner()
	exp := o
	act := db.owner
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

func TestDatabaseCopyTargetName(t *testing.T) {
	db := Database{}
	db.copyTargetName = "targetdb"

	exp := db.copyTargetName
	act := db.CopyTargetName() // Run Database.CopyTargetName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

func TestDatabaseSetCopyTargetName(t *testing.T) {
	c := "targetdb"

	db := Database{}
	db.SetCopyTargetName(c) // Run Database.SetCopyTargetName()
	exp := c
	act := db.copyTargetName
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}
