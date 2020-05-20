package sqlt

import "github.com/iancoleman/strcase"

// Database ...
type Database struct {
	name    string
	newName string
	owner   string
}

// Name returns name.
func (t *Database) Name() string {
	return t.name
}

// SetName converts string arg to snake_case then sets name.
func (t *Database) SetName(name string) {
	t.name = strcase.ToSnake(name)
}

// NewName returns new name.
func (t *Database) NewName() string {
	return t.newName
}

// SetNewName converts string arg to snake_case then sets new name.
func (t *Database) SetNewName(name string) {
	t.newName = strcase.ToSnake(name)
}

// Owner returns owner.
func (t *Database) Owner() string {
	return t.owner
}

// SetOwner sets owner.
func (t *Database) SetOwner(owner string) {
	t.owner = owner
}
