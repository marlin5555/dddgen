package entity

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen entity zero template
var entityZeroTempl = `// Package {{.Zero.Package}} Code generated, DO NOT EDIT.
package {{.Zero.Package}}
{{range .EntityImports}}
import "{{.}}"
{{end}}
{{$basep := .Base.Package}}
{{- range $k,$v := .EntityMap}}
// zero struct for {{$k}} entity start
{{$justOneEmptyLine:=$k}}
    {{- range $k1,$v1 := $v.Attrs}}
{{- $key := (printf "%s%s" $k $k1)}}
{{- if $v1.Fuzzy -}}{{- $key = $k1 -}}{{- end -}}
        {{- if not (or $v1.Shared $v1.Ref.Type)}}
// {{$key}} struct has Get{{$key}}
type {{$key}} struct {
   Value {{if $v1.Batch}}[]{{end}}{{$v1.AttributeType}}
}
// Get{{$key}} return {{$key}}.{{$v1.AttributeName}} impl {{$basep}}.{{$key}}
func (e {{$key}}) Get{{$key}}() {{if $v1.Batch}}[]{{end}}{{$v1.AttributeType}}{
    return e.Value
} 
        {{- else if and (not (eq $v1.Ref.Role $v1.Ref.EntityName)) (not $v1.Batch)}}
{{- $interface := printf "%s%s" (firstUpper $v1.Ref.Role) $v1.Ref.AttributeID}}
// {{$interface}} struct has Get{{$interface}}
type {{$interface}} struct {
   Value {{$v1.AttributeType}}
}
func (e {{$interface}}) Get{{$interface}}(){{$v1.AttributeType}}{
    return e.Value
}
{{- $interface = printf "%s%s" (firstUpper $v1.Ref.Role) (batch $v1.Ref.AttributeID)}}
// {{$interface}} struct has Get{{$interface}}
type {{$interface}} struct {
   Value []{{$v1.AttributeType}}
}
func (e {{$interface}}) Get{{$interface}}()[]{{$v1.AttributeType}}{
    return e.Value
}
        {{- end}}
    {{- end}}
{{- end}}
`
var zeroTempl = template.Must(template.New("zero").Funcs(template.FuncMap{"firstLower": util.FirstLower,
	"firstUpper": util.FirstUpper,
	"batch":      util.Batch}).Parse(entityZeroTempl))

func WriteZero(writer io.Writer, domain *entity.Group) {
	_ = zeroTempl.Execute(writer, domain)
}

// table column gen entity zero template
var entityZeroCommonTempl = `// Package {{.Zero.Package}} Code generated, DO NOT EDIT.
package {{.Zero.Package}}

import (
{{- range .EntityImports}}
    "{{.}}"
{{- end}}
    "strings"
)
{{$basep := .Base.Package}}
{{range $k,$v := .EntityMap}}
{{- range $k1,$v1 := $v.Attrs}}
{{- $key := $k1}}
{{- if $v1.Shared }}
// {{$key}} interface has Get{{$key}}
type {{$key}} struct {
   {{$v1.AttributeName}} {{$v1.AttributeType}}
}
// Get{{$key}} return {{$key}}.{{$v1.AttributeName}} impl {{$basep}}.{{$key}}
func (e {{$key}}) Get{{$key}}() {{$v1.AttributeType}}{
    return e.{{$v1.AttributeName}}
} 
{{end}}
{{- end}}
{{- end}}

// PageOrderOperator combination: PageQuery Order Operator
type PageOrderOperator struct {
	PageQuery
	Order
	Operator
}
// PageQuery ??????????????????
type PageQuery struct {
	PageIndex int  // ????????????
	PageSize  int  // ????????????
}

// GetOffset ???????????????????????? offset
func (q PageQuery) GetOffset() int {
	return (q.PageIndex - 1) * q.PageSize
}

// GetLimit ???????????????????????? limit
func (q PageQuery) GetLimit() int {
	return q.PageSize
}
// Order ????????????
type Order struct {
	SortedBy string 
	Order    string 
}

// OrderStr ??????????????????
func (o Order) OrderStr() string {
	if !o.IsEmptyStr() {
		return strings.Join([]string{o.SortedBy, o.Order}, " ")
	}
	return ""
}

// IsEmptyStr ????????????????????????????????????????????????
func (o Order) IsEmptyStr() bool {
	if o.Order == "" {
		return true
	}

	if o.SortedBy != "" {
		return false
	}

	o.SortedBy = "desc"
	return false
}

// Operator ?????????
type Operator struct {
	Rtx         string // ????????????
}
func (o Operator)GetRtx() string {
    return o.Rtx
}
`
var zeroCommonTempl = template.Must(template.New("common-zero").Parse(entityZeroCommonTempl))

func WriteCommonZero(writer io.Writer, domain *entity.Group) {
	_ = zeroCommonTempl.Execute(writer, domain)
}
