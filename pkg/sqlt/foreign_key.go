package sqlt

import (
	"github.com/iancoleman/strcase"
)

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

// ConstraintName ...
func (f *ForeignKey) ConstraintName() string {
	return f.ReferencingTableName() + "_" + f.ReferencedTableName() + "_mnrk_fkc"
}

// ReferencedTableName ...
func (f *ForeignKey) ReferencedTableName() string {
	return f.referencedTableName
}

// SetReferencedTableName ...
func (f *ForeignKey) SetReferencedTableName(name string) {
	f.referencedTableName = strcase.ToSnake(name)
}

// ReferencedColumnName ...
func (f *ForeignKey) ReferencedColumnName() string {
	return "id"
}

// ReferencingTableName ...
func (f *ForeignKey) ReferencingTableName() string {
	return f.referencingTableName
}

// SetReferencingTableName ...
func (f *ForeignKey) SetReferencingTableName(name string) {
	f.referencingTableName = strcase.ToSnake(name)
}

// ReferencingColumnName ...
func (f *ForeignKey) ReferencingColumnName() string {
	return f.ReferencedTableName() + "_id"
}

// ReferencingColumnType ...
func (f *ForeignKey) ReferencingColumnType() string {
	return "bigint"
}
