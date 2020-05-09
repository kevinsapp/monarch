package generate

// // Unit test createTableSQL()
// func TestCreateTableSQL(t *testing.T) {
// 	// Define the expected SQL string.
// 	exp := `--Up migration for users table

// CREATE TABLE users (
// 	id uuid DEFAULT gen_random_uuid() NOT NULL,

// 	-- Specify additional fields here.

// 	-- Timestamps
// 	created_at timestamp(6) without time zone NOT NULL,
// 	updated_at timestamp(6) without time zone NOT NULL,
// 	CONSTRAINT users_pkey PRIMARY KEY (id)
// );
// `
// 	// Run createTableSQL()
// 	act, err := createTableSQL("users")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if exp != act {
// 		t.Errorf("want %q; got %q", exp, act)
// 	}
// }

// // Unit test dropTableSQL()
// func TestDropTableSQL(t *testing.T) {
// 	exp := `-- Down migration for users table

// DROP TABLE users;
// `
// 	// Run dropTableSQL()
// 	act, err := dropTableSQL("users")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if exp != act {
// 		t.Errorf("want %q; got %q", exp, act)
// 	}
// }
