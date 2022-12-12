package entity

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var reqtpl = `// Package {{.Req.Package}} Code generated, DO NOT EDIT.
package {{.Req.Package}}
{{$basep := .Base.Package}}
import "{{.Module}}{{.Base.Path}}"

{{- $eMap:=.FullEntities}}
{{range $k,$v := .EntityMap}}
// get create update delete req for {{$k}}
type (
    // Get{{batch $k}}NoForeign get {{batch $k}} no foreign key
    Get{{batch $k}}NoForeign interface {
{{- range $k1,$v1 := $v.Attrs}}
{{- if $v1.Fuzzy}}
        {{$basep}}.{{$v1.AttributeName}}
{{- else if (or $v1.Exact $v1.Batch)}}
{{- if not $v1.Ref.Type}}
        {{$basep}}.{{$k}}{{$v1.AttributeName}}
{{- end}}
{{- end}}
{{- end}}
    }
    // Get{{batch $k}}Pure get {{batch $k}} pure, 仅仅作用在单一实体表上
    Get{{batch $k}}Pure interface {
{{- range $k1,$v1 := $v.Attrs}}
{{- if $v1.Fuzzy}}
        {{$basep}}.{{$v1.AttributeName}}
{{- else if (or $v1.Exact $v1.Batch)}}
{{- if $v1.Ref.Type}}
        {{$basep}}.{{firstUpper $v1.Ref.Role}}{{- if $v1.Batch}}{{batch $v1.Ref.AttributeID}}{{- else}}{{$v1.Ref.AttributeID}}{{- end}}
{{- else}}
        {{$basep}}.{{$k}}{{$v1.AttributeName}}
{{- end}}
{{- end}}
{{- end}}
    }
    // Get{{batch $k}}PO get {{batch $k}} po, 用于 persistence 上的 get 操作
    Get{{batch $k}}PO interface {
        Get{{batch $k}}Pure
        {{$basep}}.PageOrderOperator
    }

    // Get{{batch $k}}Req get {{batch $k}} req, 用于对外 service 中使用
    Get{{batch $k}}Req interface {
        Get{{batch $k}}Pure
{{- range $k1,$v1 := $v.Attrs}}
{{- if and $v1.Ref.EntityName (not $v1.Batch)}}
{{$e:= index $eMap $v1.Ref.EntityName}}
        // start. use for {{firstUpper $v1.Ref.Role}} filter
{{- range $refKey, $refAttr:=$e.Attrs}}
{{- if $refAttr.Fuzzy}}
        {{$basep}}.{{$refAttr.AttributeName}}
{{- end}}
{{- if or $refAttr.Exact $refAttr.Batch}}

{{- if $refAttr.Ref.Type}}
        {{$basep}}.{{firstUpper $refAttr.Ref.Role}}{{- if $refAttr.Batch}}{{batch $refAttr.Ref.AttributeID}}{{- else}}{{$refAttr.Ref.AttributeID}}{{- end}}
{{- else}}
        {{$basep}}.{{$v1.Ref.EntityName}}{{$refAttr.AttributeName}}
{{- end}}

{{- end}}
{{- end}}
        // end.
{{end}}
{{- end}}
        {{$basep}}.Fields
        {{$basep}}.PageOrderOperator
        {{$basep}}.Validator
    }

    Delete{{$k}}Req interface {
        {{$basep}}.{{$k}}ID
        {{$basep}}.Operator
        {{$basep}}.Validator
    }

    Update{{$k}}Req interface {
{{- range $k1,$v1 := $v.Attrs}}
{{- if or $v1.Upsert $v1.Update}}
{{- if $v1.Ref.Type}}
        {{$basep}}.{{firstUpper $v1.Ref.Role}}{{$v1.Ref.AttributeID}}
{{- else}}
        {{$basep}}.{{$k}}{{$v1.AttributeName}}
{{- end}}
{{- end}}
{{- end}}
        {{$basep}}.Operator
        {{$basep}}.Validator
    }
    Create{{$k}}Req interface {
{{- range $k1,$v1 := $v.Attrs}}
{{- if or $v1.Upsert $v1.Insert}}
{{- if $v1.Ref.Type}}
        {{$basep}}.{{firstUpper $v1.Ref.Role}}{{$v1.Ref.AttributeID}}
{{- else}}
        {{$basep}}.{{$k}}{{$v1.AttributeName}}
{{- end}}
{{- end}}
{{- end}}
        {{$basep}}.Operator
        {{$basep}}.Validator
    }
)
{{- end}}
`

var reqTempl = template.Must(template.New("req-interface").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch,
		"part":       entity.Part}).Parse(reqtpl))

func WriteReq(writer io.Writer, domain *entity.Group) {
	_ = reqTempl.Execute(writer, domain)
}
