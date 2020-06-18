package sqlt

import "github.com/iancoleman/strcase"

// ForeignKeyConstraint ...
type ForeignKeyConstraint struct {
	name                 string
	referencedTableName  string
	referencingTableName string
}

// Name ...
func (f *ForeignKeyConstraint) Name() string {
	return f.name
}

// SetName ...
func (f *ForeignKeyConstraint) SetName(name string) {
	f.name = strcase.ToSnake(name)
}

// ReferencedTableName ...
func (f *ForeignKeyConstraint) ReferencedTableName() string {
	return f.referencedTableName
}

// SetReferencedTableName ...
func (f *ForeignKeyConstraint) SetReferencedTableName(name string) {
	f.referencedTableName = strcase.ToSnake(name)
}

// ReferencingTableName ...
func (f *ForeignKeyConstraint) ReferencingTableName() string {
	return f.referencingTableName
}

// SetReferencingTableName ...
func (f *ForeignKeyConstraint) SetReferencingTableName(name string) {
	f.referencingTableName = strcase.ToSnake(name)
}
