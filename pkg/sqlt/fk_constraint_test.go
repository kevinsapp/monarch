package sqlt

import (
	"testing"
)

// Unit test ForeignKeyConstraint.Name()
func TestForeignKeyConstraintName(t *testing.T) {
	fk := ForeignKeyConstraint{}
	fk.name = "mnrk_fk_constraint_child_parent"

	exp := fk.name
	act := fk.Name()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test ForeignKeyConstraint.SetName()
func TestForeignKeyConstraintSetName(t *testing.T) {
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
		fk := ForeignKeyConstraint{}
		fk.SetName(c[0]) // Run ForeignKeyConstraint.SetName()
		exp := c[1]
		act := fk.name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test ForeignKeyConstraint.Name()
func TestForeignKeyConstraintReferencedTableName(t *testing.T) {
	fk := ForeignKeyConstraint{}
	fk.referencedTableName = "parent_table"

	exp := fk.referencedTableName
	act := fk.ReferencedTableName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test ForeignKeyConstraint.SetName()
func TestForeignKeyConstraintSetReferencedTableName(t *testing.T) {
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
		fk := ForeignKeyConstraint{}
		fk.SetReferencedTableName(c[0]) // Run ForeignKeyConstraint.SetReferencedTableName()
		exp := c[1]
		act := fk.referencedTableName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test ForeignKeyConstraint.ReferencingTableName()
func TestForeignKeyConstraintReferencingTableName(t *testing.T) {
	fk := ForeignKeyConstraint{}
	fk.referencedTableName = "parent_table"

	exp := fk.referencedTableName
	act := fk.ReferencedTableName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test ForeignKeyConstraint.SetReferencingTableName()
func TestForeignKeyConstraintSetReferencingTableName(t *testing.T) {
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
		fk := ForeignKeyConstraint{}
		fk.SetReferencingTableName(c[0]) // Run ForeignKeyConstraint.SetReferencingTableName()
		exp := c[1]
		act := fk.referencingTableName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}
