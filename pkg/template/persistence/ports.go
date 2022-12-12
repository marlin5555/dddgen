package persistence

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var persisttpl = `// Package {{.PortsPersistence.Package}} Code generated, DO NOT EDIT.
package {{.PortsPersistence.Package}}

import (
    "context"

    "{{.Module}}{{.Req.Path}}"
    "{{.Module}}{{.PO.Path}}"
)
{{$req:=.Req.Package}}
{{$po:=.PO.Package}}
{{- range $k,$v := .EntityMap}}
// {{$k}}DAO {{$k}} dao
type {{$k}}DAO interface {
    Transaction
    FieldChecker
    Get(context.Context, {{$req}}.Get{{batch $k}}PO) ({{$po}}.{{batch $k}}, uint32, error)
    Create(context.Context, {{$req}}.Create{{$k}}Req) ({{$v.IDType}},error)
    Update(context.Context, {{$req}}.Update{{$k}}Req) error
    Delete(context.Context, {{$req}}.Delete{{$k}}Req) error
}
{{end}}
`
var persistenceTempl = template.Must(template.New("persistence-ports").Funcs(template.FuncMap{"firstLower": util.FirstLower,
	"firstUpper": util.FirstUpper,
	"batch":      util.Batch}).Parse(persisttpl))

func WritePersistencePorts(writer io.Writer, domain *entity.Group) {
	_ = persistenceTempl.Execute(writer, domain)
}

// table column gen interface template
var persistctpl = `// Package {{.PortsPersistence.Package}} Code generated, DO NOT EDIT.
package {{.PortsPersistence.Package}}
import(
    "{{.Module}}{{.Base.Path}}"

    "gorm.io/gorm"
)

// Transaction 执行事务
type Transaction interface {
	Transaction(fun func(tx *gorm.DB) error) error
}

// FieldChecker own IsRelated BaseFields
type FieldChecker interface {
	IsRelated(fields base.Fields) bool
	BaseFields() base.Fields
}
`
var persistencePortsCTempl = template.Must(template.New("persistence-ports").Parse(persistctpl))

func WritePersistencePortsCommon(writer io.Writer, common *entity.Common) {
	_ = persistencePortsCTempl.Execute(writer, common)
}
