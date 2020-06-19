package sqlt

import "github.com/iancoleman/strcase"

// ForeignKey ...
type ForeignKey struct {
	name                 string
	referencedTableName  string
	referencingTableName string
}

// Name ...
func (f *ForeignKey) Name() string {
	return f.name
}

// SetName ...
func (f *ForeignKey) SetName(name string) {
	f.name = strcase.ToSnake(name)
}

// ReferencedTableName ...
func (f *ForeignKey) ReferencedTableName() string {
	return f.referencedTableName
}

// SetReferencedTableName ...
func (f *ForeignKey) SetReferencedTableName(name string) {
	f.referencedTableName = strcase.ToSnake(name)
}

// ReferencingTableName ...
func (f *ForeignKey) ReferencingTableName() string {
	return f.referencingTableName
}

// SetReferencingTableName ...
func (f *ForeignKey) SetReferencingTableName(name string) {
	f.referencingTableName = strcase.ToSnake(name)
}
