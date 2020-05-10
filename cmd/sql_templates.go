package cmd

// SQL templates for TABLE operaions
const (
	// sqltCreateTable is a SQL template for creating tables.
	sqltCreateTable string = `-- Table: {{.Name}}

CREATE TABLE {{.Name}} (
	id uuid DEFAULT gen_random_uuid() NOT NULL,

	-- Specify additional fields here.


	-- Timestamps
	created_at timestamp(6) without time zone NOT NULL,
	updated_at timestamp(6) without time zone NOT NULL,
	CONSTRAINT {{.Name}}_pkey PRIMARY KEY (id)
);
`
	// sqltDropTable is a SQL template for dropping tables.
	sqltDropTable string = `-- Table: {{.Name}}

DROP TABLE {{.Name}};
`
)
