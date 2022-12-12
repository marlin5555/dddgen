package entity

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// generate entity/{{group}}.go entity/common.go

// table column gen interface template
var entitytpl = `// Package {{.Entity.Package}} Code generated, DO NOT EDIT.
package {{.Entity.Package}}
import "{{.Module}}{{.Base.Path}}"
{{$eMap:=.FullEntities}}
{{$base:=.Base.Package}}
{{- range $k,$v := .EntityMap}}
type(
{{$entityName:=.TypeName}}
    // {{$entityName}} entity
	{{$entityName}} interface {
{{- range $v1 := $v.Columns}}
{{- if not $v1.Ref.Type}}
		{{$base}}.{{$k}}{{$v1.AttributeName}}
{{- end}}
{{- if $v1.Ref.Role }}
        {{firstUpper $v1.Ref.Role}}Getter
{{- end}}
{{- range $idx, $ref:=$v1.Fer}}
{{- if (part $ref.Type) }}
        {{$ref.EntityName}}Getter
{{- end}}
{{- end}}
{{- end}}
{{- range $v1 := $v.Combinations}}
		{{$v1.TypeName}}
{{- end}}
	}
    // {{$entityName}}Getter {{$entityName}} getter
    {{$entityName}}Getter interface {
        Get{{$entityName}}() {{$entityName}}
    }
{{- range $v1 := $v.Columns}}
{{- range $ref := $v1.Fer}}
{{- $refE:=index $eMap $ref.Role}}
{{- $refC:=index $refE.Attrs $ref.AttributeID}}
{{- $refRole:=firstUpper $refC.Ref.Role}}
{{- if not (eq $refRole $entityName)}}
    {{$refRole}}Getter interface {
        Get{{$refRole}}() {{$entityName}}
    }
{{- end}}
{{- end}}
{{- end}}

	// {{batch .TypeName}} {{firstLower .TypeName}} s
	{{batch .TypeName}} []{{.TypeName}}
)
{{- end}}
`

var entityTempl = template.Must(template.New("entity-interface").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch,
		"part":       entity.Part}).
	Parse(entitytpl))

func WriteEntity(writer io.Writer, domain *entity.Group) {
	_ = entityTempl.Execute(writer, domain)
}

// table column gen interface template
var entityctpl = `// Package {{.Entity.Package}} Code generated, DO NOT EDIT.
package {{.Entity.Package}}
import "{{.Module}}{{.Base.Path}}"
{{$base:=.Base.Package}}
{{- range $k,$v := .EntityMap}}
type(
    // {{$k}} common combination
	{{$k}} interface {
{{- range $k1, $v1 := $v.Attrs}}
		{{$base}}.{{$k1}}
{{- end}}
	}
)
{{- end}}
`

var entityCommonTempl = template.Must(template.New("entity-common-interface").
	Funcs(template.FuncMap{"firstLower": util.FirstLower}).
	Parse(entityctpl))

func WriteCommonEntity(writer io.Writer, domain *entity.Group) {
	_ = entityCommonTempl.Execute(writer, domain)
}
