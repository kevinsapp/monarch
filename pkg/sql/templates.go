package sql

// Table ...
type Table struct {
	Name    string
	NewName string
	Columns []Column
}

// Column ...
type Column struct {
	Name    string
	NewName string
	Type    string
}

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
