package entity

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var rsptpl = `// Package {{.Rsp.Package}} Code generated, DO NOT EDIT.
package {{.Rsp.Package}}
import "{{.Module}}{{.Entity.Path}}"
{{$entity:=.Entity.Package}}
{{- range $k,$v := .EntityMap}}
// {{batch .TypeName}}Rsp combination Rsp Total {{batch .TypeName}}
type {{batch .TypeName}}Rsp struct {
	Rsp
	Total uint32
	{{$entity}}.{{batch .TypeName}}
}
{{- end}}
`

var rspTempl = template.Must(template.New("rsp-interface").
	Funcs(template.FuncMap{"firstLower": util.FirstLower, "batch": util.Batch}).
	Parse(rsptpl))

func WriteRsp(writer io.Writer, domain *entity.Group) {
	_ = rspTempl.Execute(writer, domain)
}

// table column gen interface template
var rspctpl = `// Package {{.Rsp.Package}} Code generated, DO NOT EDIT.
package {{.Rsp.Package}}
import "{{.Module}}{{.Base.Path}}"
{{$base:=.Base.Package}}
// Rsp combination base.RspCode base.RspErr base.RspMsg
type Rsp interface {
	base.RspCode
	base.RspErr
	base.RspMsg
}
`
var rspCommonTempl = template.Must(template.New("rsp-common-interface").Parse(rspctpl))

func WriteCommonRsp(writer io.Writer, domain *entity.Group) {
	_ = rspCommonTempl.Execute(writer, domain)
}
