package sqlt

import (
	"testing"
)

// Unit test Table.Name()
func TestTableName(t *testing.T) {
	tbl := Table{}
	tbl.name = "test_table"

	exp := tbl.name
	act := tbl.Name() // Run Table.Name()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Table.SetName()
func TestTableSetName(t *testing.T) {
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
		tbl := Table{}
		tbl.SetName(c[0]) // Run Table.SetName()
		exp := c[1]
		act := tbl.name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Table.NewName()
func TestTableNewName(t *testing.T) {
	tbl := Table{}
	tbl.newName = "test_table"

	exp := tbl.newName
	act := tbl.NewName() // Run Table.NewName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Table.SetNewName()
func TestTableSetNewName(t *testing.T) {
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
		tbl := Table{}
		tbl.SetNewName(c[0]) // Run Table.SetNewName()
		exp := c[1]
		act := tbl.newName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Table.Columns()
func TestTableColumns(t *testing.T) {
	expCols := []Column{
		{name: "id", colType: "uuid"},
		{name: "givenName", colType: "varchar"},
		{name: "familyName", colType: "character varying"},
		{name: "created_at", colType: "timestamp(6) without time zone"},
		{name: "updated_at", colType: "timestamp(6) without time zone"},
	}

	// Assign expCols to Table.columns
	tbl := Table{name: "users"}
	tbl.columns = expCols

	actCols := tbl.Columns() // Run Table.Columns()

	for i, v := range expCols {
		// Compare column names
		exp := v.name
		act := actCols[i].name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}

		// Compare column types
		exp = v.colType
		act = actCols[i].colType
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Table.SetColumns()
func TestTableSetColumns(t *testing.T) {
	expCols := []Column{
		{name: "id", colType: "uuid"},
		{name: "givenName", colType: "varchar"},
		{name: "familyName", colType: "character varying"},
		{name: "created_at", colType: "timestamp(6) without time zone"},
		{name: "updated_at", colType: "timestamp(6) without time zone"},
	}

	// Assign expCols to Table.columns
	tbl := Table{name: "users"}
	tbl.SetColumns(expCols) // Run Table.SetColumns()

	actCols := tbl.columns

	for i, v := range expCols {
		// Compare column names
		exp := v.name
		act := actCols[i].name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}

		// Compare column types
		exp = v.colType
		act = actCols[i].colType
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Table.AddColumn()
func TestTableAddColumn(t *testing.T) {
	expCols := []Column{
		{name: "id", colType: "uuid"},
		{name: "givenName", colType: "varchar"},
		{name: "familyName", colType: "character varying"},
		{name: "created_at", colType: "timestamp(6) without time zone"},
		{name: "updated_at", colType: "timestamp(6) without time zone"},
	}

	// Add expCols to table.
	tbl := Table{name: "users"}
	for _, v := range expCols {
		tbl.AddColumn(v) // Run Table.AddColumn()
	}

	actCols := tbl.columns

	for i, v := range expCols {
		// Compare column names
		exp := v.name
		act := actCols[i].name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}

		// Compare column types
		exp = v.colType
		act = actCols[i].colType
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}
