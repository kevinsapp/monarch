package sql

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// SQL templates for TABLE operaions
const (
	// CreateTableTmpl is a SQL template for creating tables.
	CreateTableTmpl string = `-- Table: {{.Name}}

CREATE TABLE {{.Name}} (
	id uuid DEFAULT gen_random_uuid() NOT NULL,

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT {{.Name}}_pkey PRIMARY KEY (id)
);
`
	// DropTableTmpl is a SQL template for dropping tables.
	DropTableTmpl string = `-- Table: {{.Name}}

DROP TABLE {{.Name}};
`

	// RenameTableTmpl is a SQL template for renaming tables.
	RenameTableTmpl string = `-- Table: {{.Name}}

ALTER TABLE {{.Name}} RENAME TO {{.NewName}};
`

	// AddColumnTmpl is a SQL template for adding columns to a table.
	AddColumnTmpl string = `-- Table: {{.Name}}

ALTER TABLE {{.Name}}{{range .Columns}}
ADD COLUMN {{.Name}} {{.Type}}{{end}};
`
	// DropColumnTmpl is a SQL template for dropping columns from a table.
	DropColumnTmpl string = `-- Table: {{.Name}}

ALTER TABLE {{.Name}}{{range .Columns}}
DROP COLUMN {{.Name}}{{end}};
`

	// RenameColumnTmpl is a SQL template for renaming columns in a table.
	RenameColumnTmpl string = `-- Table: {{.Name}}

ALTER TABLE {{.Name}}{{range .Columns}}
RENAME COLUMN {{.Name}} TO {{.NewName}}{{end}};
`
)

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

// Column ...
type Column struct {
	name    string
	newName string
	colType string
}

// Name ...
func (c *Column) Name() string {
	return c.name
}

// SetName ...
func (c *Column) SetName(name string) {
	c.name = strcase.ToSnake(name)
}

// NewName ...
func (c *Column) NewName() string {
	return c.newName
}

// SetNewName ...
func (c *Column) SetNewName(name string) {
	c.newName = strcase.ToSnake(name)
}

// Type ...
func (c *Column) Type() string {
	return c.colType
}

// SetType ...
func (c *Column) SetType(name string) {
	c.colType = strcase.ToSnake(name)
	c.colType = strings.ToUpper(c.colType)
}
