package sqlt

import (
	"github.com/iancoleman/strcase"
)

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
	// TODO: It appears that the interval data type includes some options which are upcased.
	// c.colType = strings.ToLower(name)
	c.colType = name
}
