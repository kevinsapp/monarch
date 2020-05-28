package sqlt

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

// Index ...
type Index struct {
	tableName  string
	columnName string
}

// Name will generate a name from TableName and ColumnName followed by `_mnrk_idx`.
func (idx *Index) Name() string {
	name := fmt.Sprintf("%s_%s_mnrk_idx", idx.TableName(), idx.ColumnName())
	return name
}

// TableName ...
func (idx *Index) TableName() string {
	return idx.tableName
}

// SetTableName ...
func (idx *Index) SetTableName(name string) {
	idx.tableName = strcase.ToSnake(name)
}

// ColumnName ...
func (idx *Index) ColumnName() string {
	return idx.columnName
}

// SetColumnName ...
func (idx *Index) SetColumnName(name string) {
	idx.columnName = strcase.ToSnake(name)
}
