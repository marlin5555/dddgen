package persistence

import (
	"io"
	"text/template"

	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
)

var builderTpl = `// Package {{.POToEntity.Package}} Code generated, DO NOT EDIT.
package {{.POToEntity.Package}}
import(
    "{{.Module}}{{.Entity.Path}}"
    "{{.Module}}{{.Base.Path}}"
    "{{.Module}}{{.PO.Path}}"
)
{{- $eMap:=.FullEntities}}
{{- $poPkg:=.PO.Package}}
{{- $entityPkg:=.Entity.Package}}
// Option EntityBuilder option
type Option func(builder *EntityBuilder)

// WithFields set fields
func WithFields(fields base.Fields) Option {
	return func(builder *EntityBuilder) {
		builder.checker.addFields(fields)
	}
}

{{- range $k,$v := .FullEntities}}
// With{{firstUpper (batch $k)}} set {{firstLower (batch $k)}}
func With{{firstUpper (batch $k)}}(pValue map[{{$v.IDType}}]{{$poPkg}}.{{firstUpper $k}}) Option {
	return func(builder *EntityBuilder) {
		builder.po{{firstUpper (batch $k)}} = pValue
	}
}
{{- end}}

// EntityBuilder builder support {{range $k,$v := .FullEntities}}{{$entityPkg}}.{{firstUpper $k}}/{{- end}}
type EntityBuilder struct {
{{- range $k,$v := .FullEntities}}
{{- $eNames:=firstUpper (batch $k)}}
{{- $eName:=firstUpper $k}}
{{- $eNameL:=firstLower $k}}
	/** {{firstLower $k}} **/
	po{{$eNames}} map[{{$v.IDType}}]{{$poPkg}}.{{$eName}}
	en{{$eNames}} map[{{$v.IDType}}]*{{$eNameL}}Entity
{{- end}}
	checker fieldChecker
}

// NewBuilder builder constructor
func NewBuilder(fields base.Fields, opts ...Option) *EntityBuilder {
	b := &EntityBuilder{
{{- range $k,$v := .FullEntities}}
{{- $eNames:=firstUpper (batch $k)}}
{{- $eName:=firstUpper $k}}
{{- $eNameL:=firstLower $k}}
	    po{{$eNames}}: map[{{$v.IDType}}]{{$poPkg}}.{{$eName}}{},
	    en{{$eNames}}: map[{{$v.IDType}}]*{{$eNameL}}Entity{},
{{- end}}
	}
	b.checker = buildChecker(fields)
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// Build build entity and relation
func (b *EntityBuilder) Build(){
    b.buildEntity()
    b.buildRelation()
}

{{- range $k,$v := .FullEntities}}
{{- $e:=firstUpper (batch $k)}}
// Get{{$e}} get {{firstUpper $k}} entities
func (b *EntityBuilder)Get{{$e}}() {{$entityPkg}}.{{$e}} {
	var r {{$entityPkg}}.{{$e}}
	for _, e := range b.en{{$e}} {
		r = append(r, e)
	}
	return r
}
{{- end}}

// buildEntity build entity
func (b *EntityBuilder) buildEntity() {
{{- range $k,$v := .FullEntities}}
	b.build{{firstUpper $k}}Entity()
{{- end}}
}

{{- range $k,$v := .FullEntities}}
func (b *EntityBuilder) build{{firstUpper $k}}Entity() {
	for k, p := range b.po{{firstUpper (batch $k)}} {
		b.en{{firstUpper (batch $k)}}[k] = build{{firstUpper $k}}(b.checker, p)
	}
}
{{- end}}

// buildRelation build relation
func (b *EntityBuilder) buildRelation() {
{{- range $k,$v := .FullEntities}}
	b.build{{firstUpper $k}}Relation()
{{- end}}
}

{{- range $k,$v := .FullEntities}}
func (b *EntityBuilder) build{{$k}}Relation() {
{{$needGetter:=false}}
{{- range $k1,$v1:=$v.Attrs}}
    {{- if not $v1.Batch}}
        {{- if $v1.Ref.Type}}{{- $needGetter = true}}{{- end}}
        {{- range $idx, $ref :=$v1.Fer}}{{- if part $ref.Type}}{{- $needGetter = true}}{{- end}}{{- end}}
    {{- end}}
{{- end}}

{{- if $needGetter}}
	// 进行 {{$k}} - reference 关联
	for k, e := range b.en{{batch $k}} {
{{- range $k1,$v1:=$v.Attrs}}
{{- if not $v1.Batch}}
{{- if $v1.Ref.Type}}
		if a, ok := b.en{{batch $v1.Ref.EntityName}}[b.po{{batch $k}}[k].{{$k1}}]; ok {
			e.{{firstLower $v1.Ref.Role}}Getter = {{firstLower $v1.Ref.Role}}Getter{value: a}
		}
{{- end}}
{{- range $idx, $ref :=$v1.Fer}}
{{- if part $ref.Type}}
        if a, ok:=b.en{{batch $ref.EntityName}}[b.po{{batch $k}}[k].{{$k1}}]; ok {
            e.{{firstLower $ref.Role}}Getter = {{firstLower $ref.Role}}Getter{value: a}
        }
{{- end}}
{{- end}}
{{- end}}
{{- end}}
	}
{{- end}}
}
{{- end}}
`

var builderTempl = template.Must(template.New("builder").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch,
		"part":       entity.Part}).Parse(builderTpl))

func WriteBuilder(writer io.Writer, entities map[string]*entity.Entity, common *entity.Common) {
	_ = builderTempl.Execute(writer, struct {
		*entity.Common
		FullEntities map[string]*entity.Entity
	}{Common: common, FullEntities: entities})
}

// table column gen interface template
var poToEntityTpl = `// Package {{.POToEntity.Package}} Code generated, DO NOT EDIT.
package {{.POToEntity.Package}}
{{- $entityPkg:=.Entity.Package}}
{{- $basePkg:=.Base.Package}}
{{- $poPkg:=.PO.Package}}
{{- $zeroPkg:=.Zero.Package}}
{{- $persistPkg:=.Persistence.Package}}
{{- $logPkg:=.Log.Package}}
import(
    "{{.Module}}{{.Entity.Path}}"
    "{{.Module}}{{.Base.Path}}"
    "{{.Module}}{{.Zero.Path}}"
    "{{.Module}}{{.Persistence.Path}}"
    "{{.Module}}{{.PO.Path}}"
    "{{.Module}}{{.Log.Path}}"
)
{{- range $k,$v := .EntityMap}}
// {{firstLower $k}}Getter impl {{$entityPkg}}.{{firstUpper $k}}Getter
type {{firstLower $k}}Getter struct {
	value {{$entityPkg}}.{{firstUpper $k}}
}
// Get{{$k}} return {{$entityPkg}}.{{$k}}
func (g {{firstLower $k}}Getter) Get{{$k}}() {{$entityPkg}}.{{$k}} {
	return g.value
}
  {{- range $k1, $v1:=$v.Attrs}}
    {{- if and (not $v1.Batch) (and $v1.Ref.Type (not (eq $v1.Ref.Role $v1.Ref.EntityName)))}}
      {{- $e:=$v1.Ref.Role}}
      {{- $e1:=$v1.Ref.EntityName}}
// {{firstLower $e}}Getter impl {{$entityPkg}}.{{firstUpper $e}}Getter
type {{firstLower $e}}Getter struct {
	value {{$entityPkg}}.{{firstUpper $e1}}
}
// Get{{firstUpper $e}} return {{$entityPkg}}.{{firstUpper $e1}}
func (g {{firstLower $e}}Getter) Get{{firstUpper $e}}() {{$entityPkg}}.{{firstUpper $e1}} {
	return g.value
}
    {{- end}}
  {{- end}}
{{- end}}


{{- range $k,$v := .EntityMap}}
// {{firstLower $k}}Entity {{firstLower $k}} entity impl {{$entityPkg}}.{{firstUpper $k}}
type {{firstLower $k}}Entity struct {
  {{- range $k1,$v1:=$v.Attrs}}
    {{- if not (or $v1.Batch $v1.Fuzzy)}}
      {{- if $v1.Shared}}
    {{$basePkg}}.{{$k1}}
      {{- else if $v1.Ref.Type}}
    {{firstLower $v1.Ref.Role}}Getter
      {{- else}}
    {{$basePkg}}.{{$k}}{{$k1}}
      {{- end}}

      {{- range $idx,$ref:=$v1.Fer}}
        {{- if part $ref.Type}}
    {{firstLower $ref.Role}}Getter
        {{- end}}
      {{- end}}
    {{- end}}
  {{- end}}
}

  {{- range $k1,$v1:=$v.Attrs}}
    {{- if not (or $v1.Ref.Type (or $v1.Fuzzy $v1.Batch))}}
      {{- $attrName:=$k1}}
      {{- if not $v1.Shared}}{{- $attrName = printf "%s%s" $k $k1}}{{- end}}
func (e *{{firstLower $k}}Entity) set{{$attrName}}(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.{{$attrName}}); ok {
			e.{{$attrName}} = w
			return
		}
		{{$logPkg}}.InfofWithFuncName("value[%+v] not a base.{{$attrName}}", value)
	}
	// 零值处理
	e.{{$attrName}} = {{$zeroPkg}}.{{$attrName}}{}
}
    {{- end}}
  {{- end}}

// build{{$k}} build entity
func build{{$k}}(checker fieldChecker, po {{$poPkg}}.{{$k}}) *{{firstLower $k}}Entity {
	e := &{{firstLower $k}}Entity{}
	t := newTable({{$persistPkg}}.T_{{$k}})
	for _, tup := range []tuple{
  {{- range $k1,$v1:=$v.Attrs}}
    {{- if not (or $v1.Ref.Type (or $v1.Fuzzy $v1.Batch))}}
      {{- $attrName:=$k1}}
      {{- if not $v1.Shared}}{{- $attrName = printf "%s%s" $k $k1}}{{- end}}
        {pred: t.col("{{$v1.ColumnName}}").check(checker.hasColumn), attr: po, do: e.set{{$attrName}}},
    {{- end}}
  {{- end}}
	} {
		builderAppend(tup)
	}
	return e
}
{{- end}}
`

var poToEntityTempl = template.Must(template.New("persistence-po-to-entity").Funcs(template.FuncMap{"firstLower": util.FirstLower,
	"firstUpper": util.FirstUpper,
	"batch":      util.Batch, "part": entity.Part}).Parse(poToEntityTpl))

func WritePOToEntity(writer io.Writer, domain *entity.Group) {
	_ = poToEntityTempl.Execute(writer, domain)
}

var builderCommonTpl = `// Package {{.POToEntity.Package}} Code generated, DO NOT EDIT.
package {{.POToEntity.Package}}
import (
	"fmt"
	"strings"

    "{{.Module}}{{.Base.Path}}"
    "{{.Module}}{{.Persistence.Path}}"
)
{{- $persistPkg:=.Persistence.Package}}
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
	_, ok := c.cmap[fmt.Sprintf({{$persistPkg}}.COLUMN_PATTERN, tableName, columnName)]
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

`

var builderCommonTempl = template.Must(template.New("persistence-builder-common").Funcs(template.FuncMap{"firstLower": util.FirstLower,
	"firstUpper": util.FirstUpper,
	"batch":      util.Batch, "part": entity.Part}).Parse(builderCommonTpl))

func WriteBuilderCommon(writer io.Writer, common *entity.Common) {
	_ = builderCommonTempl.Execute(writer, common)
}
