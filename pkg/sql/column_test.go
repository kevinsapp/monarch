package sql

import (
	"testing"
)

// Unit test Column.Name()
func TestColumnName(t *testing.T) {
	col := Column{}
	col.name = "test_column"

	exp := col.name
	act := col.Name() // Run Column.Name()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Column.SetName()
func TestColumnSetName(t *testing.T) {
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
		col := Table{}
		col.SetName(c[0]) // Run Column.SetName()
		exp := c[1]
		act := col.name
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Column.NewName()
func TestColumnNewName(t *testing.T) {
	col := Table{}
	col.newName = "test_column"

	exp := col.newName
	act := col.NewName() // Run Column.NewName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Column.SetNewName()
func TestColumnSetNewName(t *testing.T) {
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
		col := Table{}
		col.SetNewName(c[0]) // Run Column.SetNewName()
		exp := c[1]
		act := col.newName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Column.NewName()
func TestColumnType(t *testing.T) {
	cases := []string{
		// Numeric types
		"smallint",
		"integer",
		"bigint",
		"decimal",
		"numeric",
		"real",
		"double precision",
		"smallserial",
		"serial",
		"bigserial",
		// Monetary types
		"money",
		// Character types
		"character varying",
		"varchar",
		"character varying(250)",
		"varchar(100)",
		"text",
		// Binary data types
		"bytea",
		// Date/time types
		"timestamp(6) without time zone",
		"timestamp(0) with time zone",
		"date",
		"time(6) without time zone",
		"time(0) with time zone",
		"interval YEAR",
		"interval MINUTE TO SECOND(6)",
		// Boolean type
		"boolean",
		// Geometric types
		"point",
		"line",
		"lseg",
		"box",
		"path",
		"polygon",
		"circle",
		// Network address types
		"cidr",
		"inet",
		"macaddr",
		"macaddr8",
		// UUID type
		"uuid",
	}
	for _, v := range cases {
		col := Column{}
		col.colType = v

		exp := v
		act := col.Type() // Run Column.Type()
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Column.SetType()
func TestColumnSetType(t *testing.T) {
	cases := []string{
		// Numeric types
		"smallint",
		"integer",
		"bigint",
		"decimal",
		"numeric",
		"real",
		"double precision",
		"smallserial",
		"serial",
		"bigserial",
		// Monetary types
		"money",
		// Character types
		"character varying",
		"varchar",
		"character varying(250)",
		"varchar(100)",
		"text",
		// Binary data types
		"bytea",
		// Date/time types
		"timestamp(6) without time zone",
		"timestamp(0) with time zone",
		"date",
		"time(6) without time zone",
		"time(0) with time zone",
		"interval YEAR",
		"interval MINUTE TO SECOND(6)",
		// Boolean type
		"boolean",
		// Geometric types
		"point",
		"line",
		"lseg",
		"box",
		"path",
		"polygon",
		"circle",
		// Network address types
		"cidr",
		"inet",
		"macaddr",
		"macaddr8",
		// UUID type
		"uuid",
	}
	for _, v := range cases {
		col := Column{}
		col.SetType(v) // Run Column.SetType()

		exp := v
		act := col.colType
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}
