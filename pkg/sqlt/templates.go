package sqlt

import (
	"bytes"
	"text/template"
)

// SQL templates for DATABASE operaions
const (
	// CreateDBTmpl is a SQL template for creating databases.
	// TODO: consider adding options:
	// ENCODING = 'UTF8'
	// LC_COLLATE = 'en_US.utf8'
	// LC_CTYPE = 'en_US.utf8'
	// TABLESPACE = pg_default
	// CONNECTION LIMIT = -1
	CreateDBTmpl string = `-- Database: {{.Name}}

CREATE DATABASE {{.Name}}
	WITH 
	OWNER = {{.Owner}};
`
	// DropTableTmpl is a SQL template for dropping databases.
	DropDBTmpl string = `-- Database: {{.Name}}

DROP DATABASE IF EXISTS {{.Name}};
`
)

// SQL templates for TABLE operaions
const (
	// CreateTableTmpl is a SQL template for creating tables.
	CreateTableTmpl string = `CREATE TABLE {{.Name}} (
	id bigint NOT NULL,
	{{range .Columns}}{{.Name}} {{.Type}},
	{{end}}

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT {{.Name}}_pkey PRIMARY KEY (id)
);
`
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
