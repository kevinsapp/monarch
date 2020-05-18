package sql

import "github.com/iancoleman/strcase"

// Database ...
type Database struct {
	name    string
	newName string
	owner   string
}

// Name ...
func (t *Database) Name() string {
	return t.name
}

// SetName ...
func (t *Database) SetName(name string) {
	t.name = strcase.ToSnake(name)
}

// NewName ...
func (t *Database) NewName() string {
	return t.newName
}

// SetNewName ...
func (t *Database) SetNewName(name string) {
	t.newName = strcase.ToSnake(name)
}

// Owner ...
func (t *Database) Owner() string {
	return t.owner
}

// SetOwner ...
func (t *Database) SetOwner(owner string) {
	t.owner = strcase.ToSnake(owner)
}
