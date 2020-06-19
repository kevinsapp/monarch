package sqlt

import (
	"testing"
)

// Unit test ForeignKey.Name()
func TestForeignKeyName(t *testing.T) {
	fk := ForeignKey{}
	fk.name = "mnrk_fk_child_parent"

	exp := fk.name
	act := fk.Name()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test ForeignKey.SetName()
func TestForeignKeySetName(t *testing.T) {
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
		fk := ForeignKey{}
		fk.SetName(c[0]) // Run ForeignKey.SetName()
		exp := c[1]
		act := fk.name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test ForeignKey.Name()
func TestForeignKeyReferencedTableName(t *testing.T) {
	fk := ForeignKey{}
	fk.referencedTableName = "parent_table"

	exp := fk.referencedTableName
	act := fk.ReferencedTableName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test ForeignKey.SetName()
func TestForeignKeySetReferencedTableName(t *testing.T) {
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
		fk := ForeignKey{}
		fk.SetReferencedTableName(c[0]) // Run ForeignKey.SetReferencedTableName()
		exp := c[1]
		act := fk.referencedTableName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test ForeignKey.ReferencingTableName()
func TestForeignKeyReferencingTableName(t *testing.T) {
	fk := ForeignKey{}
	fk.referencedTableName = "parent_table"

	exp := fk.referencedTableName
	act := fk.ReferencedTableName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test ForeignKey.SetReferencingTableName()
func TestForeignKeySetReferencingTableName(t *testing.T) {
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
		fk := ForeignKey{}
		fk.SetReferencingTableName(c[0]) // Run ForeignKey.SetReferencingTableName()
		exp := c[1]
		act := fk.referencingTableName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}
