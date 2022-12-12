package persistence

import (
	"io"
	"text/template"

	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
)

// table column gen interface template
var constantTpl = `// Package {{.Persistence.Package}} Code generated, DO NOT EDIT.
package {{.Persistence.Package}}

// all tables
const (
{{- range $k,$v:=.Entities}}
    T_{{firstUpper $k}} = "{{$v.TableName}}"
{{- end}}
)

const (
    COLUMN_PATTERN = "%s.%s"
)
`

var constantTempl = template.Must(template.New("persistence-constant").Funcs(template.FuncMap{"firstLower": util.FirstLower,
	"firstUpper": util.FirstUpper,
	"batch":      util.Batch}).Parse(constantTpl))

func WriteConstant(writer io.Writer, entities map[string]*entity.Entity, common *entity.Common) {
	_ = constantTempl.Execute(writer, struct {
		*entity.Common
		Entities map[string]*entity.Entity
	}{Entities: entities, Common: common})
}
