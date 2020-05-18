package sqlt

import "github.com/iancoleman/strcase"

// Table ...
type Table struct {
	name    string
	newName string
	columns []Column
}

// Name ...
func (t *Table) Name() string {
	return t.name
}

// SetName ...
func (t *Table) SetName(name string) {
	t.name = strcase.ToSnake(name)
}

// NewName ...
func (t *Table) NewName() string {
	return t.newName
}

// SetNewName ...
func (t *Table) SetNewName(name string) {
	t.newName = strcase.ToSnake(name)
}

// Columns ...
func (t *Table) Columns() []Column {
	return t.columns
}

// SetColumns ...
func (t *Table) SetColumns(cols []Column) {
	t.columns = cols
}

// AddColumn ...
func (t *Table) AddColumn(col Column) {
	t.columns = append(t.columns, col)
}
