package sqlt

import (
	"bytes"
	"text/template"
)

// SQL templates for DATABASE operaions
const (
	// CreateDBTmpl is a SQL template for creating databases.
	CreateDBTmpl string = `CREATE DATABASE {{.Name}} OWNER {{.Owner}};`

	// CopyTableTmpl is a SQL template for copying databases.
	CopyDBTmpl string = `CREATE DATABASE {{.CopyTargetName}} TEMPLATE {{.Name}};`

	// DropTableTmpl is a SQL template for dropping databases.
	DropDBTmpl string = `DROP DATABASE IF EXISTS {{.Name}};`

	// RenameTableTmpl is a SQL template for renaming databases.
	RenameDBTmpl string = `ALTER DATABASE {{.Name}} RENAME TO {{.NewName}};`
)

// SQL templates for TABLE operaions
const (
	// CreateTableTmpl is a SQL template for creating tables.
	CreateTableTmpl string = `CREATE TABLE {{.Name}} (
	PRIMARY KEY (id),
	id bigserial NOT NULL,
	{{- range .Columns}}
	{{.Name}} {{.Type}},
	{{- end}}

	-- Specify additional fields here.

	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL
);`

	// DropTableTmpl is a SQL template for dropping tables.
	DropTableTmpl string = `DROP TABLE {{.Name}};`

	// RenameTableTmpl is a SQL template for renaming tables.
	RenameTableTmpl string = `ALTER TABLE {{.Name}} RENAME TO {{.NewName}};`

	// AddColumnTmpl is a SQL template for adding columns to a table.
	AddColumnTmpl string = `ALTER TABLE {{.Name}}{{range $i, $col := .Columns}}{{if $i}},{{end}}
ADD COLUMN {{$col.Name}} {{$col.Type}}{{end}};`

	// DropColumnTmpl is a SQL template for dropping columns from a table.
	DropColumnTmpl string = `ALTER TABLE {{.Name}}{{$l := len .Columns}}{{range $i, $col := .Columns}}{{if $i}},{{end}}
DROP COLUMN IF EXISTS {{$col.Name}}{{end}};`

	// RenameColumnTmpl is a SQL template for renaming columns in a table.
	RenameColumnTmpl string = `{{ $table := .Name }}{{range .Columns}}ALTER TABLE {{ $table }}
RENAME COLUMN {{.Name}} TO {{.NewName}};

{{end}}`

	// CreateDefaultIndexTmpl is a SQL template for creating an index of the default type on one column.
	CreateDefaultIndexTmpl string = `CREATE INDEX {{.Name}} ON {{.TableName}} ({{.ColumnName}});`

	// DropIndexTmpl is a SQL template for dropping and index.
	DropIndexTmpl string = `DROP INDEX IF EXISTS {{.Name}};`
)

// ProcessTmpl applies a data structure to a SQL template and returns a string.
func ProcessTmpl(data interface{}, sqlt string) (string, error) {
	// Initialize a template.
	var s string
	t := template.New("t")

	// Parse the template.
	t, err := t.Parse(sqlt)
	if err != nil {
		return s, err
	}

	// Apply the data structure to the template and write the result to a buffer.
	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		return s, err
	}

	// Get contents of the byte buffer as a string.
	s = b.String()

	return s, err
}
