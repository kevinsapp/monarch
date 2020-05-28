package sqlt

import (
	"testing"
)

// Unit test Index.Name()
func TestIndexName(t *testing.T) {
	cases := [][]string{
		// {tableName, columnName, expectedString}
		{"table", "column", "table_column_mnrk_idx"},
		{"table_name", "column_name", "table_name_column_name_mnrk_idx"},
		{"tableName", "column_name", "table_name_column_name_mnrk_idx"},
		{"TableName", "column_name", "table_name_column_name_mnrk_idx"},
		{"Table", "column", "table_column_mnrk_idx"},
		{"ManyWordsTable", "many_words_column", "many_words_table_many_words_column_mnrk_idx"},
	}
	for _, c := range cases {
		idx := Index{}
		idx.SetTableName(c[0]) // Run Index.SetTableName()
		idx.SetColumnName(c[1])
		exp := c[2]
		act := idx.Name()
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Index.TableName()
func TestIndexTableName(t *testing.T) {
	idx := Index{}
	idx.tableName = "test_table"

	exp := idx.tableName
	act := idx.TableName() // Run Index.TableName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Index.SetTableName()
func TestIndexSetTableName(t *testing.T) {
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
		idx := Index{}
		idx.SetTableName(c[0]) // Run Index.SetTableName()
		exp := c[1]
		act := idx.tableName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}

// Unit test Index.ColumnName()
func TestIndexColumnName(t *testing.T) {
	idx := Index{}
	idx.columnName = "phone_number"

	exp := idx.columnName
	act := idx.ColumnName() // Run Index.ColumnName()
	if exp != act {
		t.Errorf("want %q; got %q", exp, act)
	}
}

// Unit test Index.SetTableName()
func TestIndexSetColumnName(t *testing.T) {
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
		idx := Index{}
		idx.SetColumnName(c[0]) // Run Index.SetColumnName()
		exp := c[1]
		act := idx.columnName
		if exp != act {
			t.Errorf("want %q; got %q", exp, act)
		}
	}
}
