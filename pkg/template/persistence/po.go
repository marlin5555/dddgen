package persistence

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var tpl = `// Package {{.PO.Package}} Code generated, DO NOT EDIT.
package {{.PO.Package}}
import (
{{- range .PoImports}}
    "{{.}}"
{{- end}}

    "{{.Module}}{{.Base.Path}}"
)
{{- $eMap:=.FullEntities}}

{{- range $k,$v := .EntityMap}}
// {{$k}} table name: {{$v.TableName}}
type {{$k}} struct {
{{- range $k1,$v1 := $v.Attrs -}}
{{if not (or $v1.Fuzzy $v1.Batch)}}
    {{$v1.AttributeName}} {{$v1.AttributeType}} ` + "`" + `gorm:"{{$v1.Gorm}}"` + "`" + `
{{- end}}
{{- end}}
}

type {{batch $k}} []{{$k}}

func (p {{batch $k}}) Get{{$k}}IDs()[]{{$v.IDType}}{
    return base.To{{$k}}IDs(p)
}

{{- range $k1,$v1 := $v.Attrs}}
{{- if and $v1.Ref.Type (not $v1.Batch)}}
{{- $refIDType:=(index $eMap $v1.Ref.EntityName).IDType}}
func (p {{batch $k}}) Get{{firstUpper $v1.Ref.Role}}{{batch $v1.Ref.AttributeID}}()[]{{$refIDType}}{
    return base.To{{firstUpper $v1.Ref.Role}}{{batch $v1.Ref.AttributeID}}(p)
}
{{- end}}
{{- end}}

// TableName {{$k}} impl schema.Tabler
func (p {{$k}}) TableName() string {
    return "{{$v.TableName}}"
}
{{- range $k1,$v1 := $v.Attrs}}
{{if not $v1.Fuzzy}}
{{- $key := (printf "%s%s" $k $k1)}}
{{- if $v1.Shared }}{{- $key = $k1}}{{end}}
{{- if $v1.Ref.Type}}
{{- $key = printf "%s%s" (firstUpper $v1.Ref.Role) $v1.Ref.AttributeID}}
{{- if $v1.Batch}}{{- $key = batch $key}}{{- end}}
{{- end}}
// Get{{$key}} {{$k}} impl base.{{$v1.AttributeName}}
func (p {{$k}}) Get{{$key}}() {{if $v1.Batch}}[]{{end}}{{$v1.AttributeType}} {
{{- if $v1.Batch}}
    return []{{$v1.AttributeType}}{p.{{$v1.OrigAttrName}}}
{{- else}}
    return p.{{$v1.AttributeName}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
`

var poTempl = template.Must(template.New("persistence-po").Funcs(template.FuncMap{"firstLower": util.FirstLower,
	"firstUpper": util.FirstUpper,
	"batch":      util.Batch}).Parse(tpl))

func WritePO(writer io.Writer, domain *entity.Group) {
	_ = poTempl.Execute(writer, domain)
}
