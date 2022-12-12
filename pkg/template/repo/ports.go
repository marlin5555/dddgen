package repo

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var repoporttpl = `// Package {{.PortsRepo.Package}} Code generated, DO NOT EDIT.
package {{.PortsRepo.Package}}
import (
	"context"

    "{{.Module}}{{.Entity.Path}}"
    "{{.Module}}{{.Req.Path}}"
)
{{$entity:=.Entity.Package}}
{{$req:=.Req.Package}}
// {{firstUpper .Name}}Repository {{firstUpper .Name}} repo
type {{firstUpper .Name}}Repository interface {
{{- range $k,$v := .EntityMap}}
	// Create{{$k}} create {{$k}}
	Create{{$k}}(context.Context, {{$req}}.Create{{$k}}Req) ({{$v.IDType}}, error)
	// Update{{$k}} update {{$k}}
	Update{{$k}}(context.Context, {{$req}}.Update{{$k}}Req) error
	// Delete{{$k}} delete {{$k}}
	Delete{{$k}}(context.Context, {{$req}}.Delete{{$k}}Req) error
	// Get{{batch $k}} general get {{$k}} method
	Get{{batch $k}}(context.Context, {{$req}}.Get{{batch $k}}Req) ({{$entity}}.{{batch $k}}, uint32, error)
{{end}}
}
`
var repoportTempl = template.Must(template.New("repo-ports").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch}).
	Parse(repoporttpl))

func WriteRepoPorts(writer io.Writer, domain *entity.Group) {
	_ = repoportTempl.Execute(writer, domain)
}
