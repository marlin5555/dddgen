// Package potoentity Code generated, DO NOT EDIT.
package potoentity

import (
	"fmt"
	"strings"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence"
)

// fieldChecker field checker
type fieldChecker struct {
	cmap map[string]bool // column name map
	tmap map[string]bool // table name map
}

// hasTable check has tableName
func (c *fieldChecker) hasTable(tableName string) bool {
	if c == nil {
		return false
	}
	if c.tmap == nil {
		return false
	}
	_, ok := c.tmap[tableName]
	return ok
}

// hasColumn check has tableName.columnName
func (c *fieldChecker) hasColumn(tableName, columnName string) bool {
	if c == nil {
		return false
	}
	if c.cmap == nil {
		return false
	}
	_, ok := c.cmap[fmt.Sprintf(persistence.COLUMN_PATTERN, tableName, columnName)]
	return ok
}

// buildChecker build checker
func buildChecker(fields base.Fields) fieldChecker {
	r := fieldChecker{
		cmap: map[string]bool{},
		tmap: map[string]bool{},
	}
	for _, f := range fields.GetFieldNames() {
		r.cmap[f] = true
		r.tmap[strings.Split(f, ".")[0]] = true
	}
	return r
}

// addFields append fields
func (c *fieldChecker) addFields(fields base.Fields) {
	for _, f := range fields.GetFieldNames() {
		c.cmap[f] = true
		c.tmap[strings.Split(f, ".")[0]] = true
	}
}

// newTable table construct
func newTable(name string) table {
	return table{tTame: name}
}

// table 方便构建 column 对象
type table struct {
	tTame string
}

// col 使用 table 构建 column 对象
func (t *table) col(c string) *column {
	return &column{
		table: t,
		cName: c,
	}
}

// column 对 tableName/columnName 进行封装，方便进行 check
type column struct {
	*table
	cName string
}

// check 对 column 进行检查
func (c *column) check(ck check) predicate {
	return func() bool {
		return ck(c.tTame, c.cName)
	}
}

// check 用于检查 column 的抽象方法，具体实现参见 fieldChecker.hasColumn
type check func(table, column string) bool

// predicate 执行（ do ）时需要提前进行的检查
type predicate func() bool

// do 具体的执行动作，检查结果由 predResult 作为入参传入
type do func(predResult bool, value interface{})

// tuple 方便执行 do (pred(), attr)
type tuple struct {
	pred predicate
	attr interface{}
	do   do
}

// builderAppend 执行 tuple 指定的行为
func builderAppend(t tuple) {
	t.do(t.pred(), t.attr)
}
