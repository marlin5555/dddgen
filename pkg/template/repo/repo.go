package repo

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var repotpl = `// Package {{.PortsRepo.Package}} Code generated, DO NOT EDIT.
package {{.PortsRepo.Package}}
import (
    "context"

    "{{.Module}}{{.Entity.Path}}"
    "{{.Module}}{{.Base.Path}}"
    "{{.Module}}{{.Req.Path}}"
    "{{.Module}}{{.Zero.Path}}"
    "{{.Module}}{{.PortsPersistence.Path}}"
    "{{.Module}}{{.PortsRepo.Path}}"
    "{{.Module}}{{.PO.Path}}"
    "{{.Module}}{{.POToEntity.Path}}"
)

{{- $basePkg:=.Base.Package}}
{{- $poPkg:=.PO.Package}}
{{- $zeroPkg:=.Zero.Package}}
{{- $entityPkg:=.Entity.Package}}
{{- $reqPkg:=.Req.Package}}
{{- $ppPkg:=.PortsPersistence.Package}}
{{- $prPkg:=.PortsRepo.Package}}
{{- $potoePkg:=.POToEntity.Package}}
{{- $es:= entities .}}
// {{firstUpper .Name}}Repo ...
type {{firstUpper .Name}}Repo struct {
{{- range $es}}
    {{firstLower .}}DAO {{$ppPkg}}.{{.}}DAO
{{- end}}
}

// New{{firstUpper .Name}}Repo ...
func New{{firstUpper .Name}}Repo(
{{- range $es}}{{firstLower .}}DAO {{$ppPkg}}.{{.}}DAO,{{end}}) {{$prPkg}}.{{firstUpper .Name}}Repository {
    return &{{firstUpper .Name}}Repo{
{{- range $es}}
        {{firstLower .}}DAO: {{firstLower .}}DAO,
{{- end}}
	}
}

{{- range $k,$v := .EntityMap}}
// Create{{$k}} create {{$k}}
func (r *{{firstUpper .Name}}Repo) Create{{$k}}(ctx context.Context, req {{$reqPkg}}.Create{{$k}}Req) ({{$v.IDType}}, error) {
    return r.{{firstLower $k}}DAO.Create(ctx, req)
}

// Update{{$k}} update {{$k}}
func (r *{{firstUpper .Name}}Repo) Update{{$k}}(ctx context.Context, req {{$reqPkg}}.Update{{$k}}Req) error {
    return r.{{firstLower $k}}DAO.Update(ctx, req)
}
// Delete{{$k}} delete {{$k}}
func (r *{{firstUpper .Name}}Repo) Delete{{$k}}(ctx context.Context, req {{$reqPkg}}.Delete{{$k}}Req) error {
    return r.{{firstLower $k}}DAO.Delete(ctx, req)
}
// Get{{batch $k}} general get {{$k}} method
func (r *{{firstUpper .Name}}Repo) Get{{batch $k}}(ctx context.Context, rq {{$reqPkg}}.Get{{batch $k}}Req) ({{$entityPkg}}.{{batch $k}}, uint32, error) {
{{- $needRefTable:=false}}
{{- range $v1 := $v.Columns}}{{- if $v1.Ref.EntityName}}{{- $needRefTable = true}}{{- end}}{{- end}}
{{- if $needRefTable}}
    var err error
    var isStop bool
  {{- range $v1 := $v.Columns}}
    {{- if $v1.Ref.EntityName}}
    var {{firstLower $v1.AttributeName}}s {{$basePkg}}.{{$v1.Ref.EntityName}}IDs
    if {{firstLower $v1.AttributeName}}s, isStop, err = get{{$v1.Ref.EntityName}}IDsByRef(ctx, r.{{firstLower $v1.Ref.EntityName}}DAO, rq); isStop{
        return nil, 0, err
    }
    {{- end}}
  {{- end}}
{{- end}}
    // 做主表查询，获取到主表记录
	result, total, err := r.{{firstLower $k}}DAO.Get(ctx, struct {
		{{$reqPkg}}.Get{{batch $k}}NoForeign
		{{$basePkg}}.PageOrderOperator // 分页查询信息
{{- range $v1 := $v.Columns}}
{{- if $v1.Ref.EntityName}}
        {{$zeroPkg}}.{{firstUpper $v1.Ref.Role}}{{$v1.Ref.AttributeID}}
        {{$basePkg}}.{{firstUpper $v1.Ref.Role}}{{$v1.Ref.AttributeID}}s
{{- end}}
{{- end}}
	}{
		Get{{batch $k}}NoForeign: rq,
		PageOrderOperator: rq,
{{- range $v1 := $v.Columns}}
{{- if $v1.Ref.EntityName}}
        {{firstUpper $v1.Ref.Role}}{{$v1.Ref.AttributeID}}s: {{$zeroPkg}}.{{firstUpper $v1.Ref.Role}}{{$v1.Ref.AttributeID}}s{Value: {{firstLower $v1.AttributeName}}s.Get{{$v1.Ref.EntityName}}IDs()},
{{- end}}
{{- end}}
    })
    // 关联表查询
{{- range $v1:=$v.Columns}}
  {{- if $v1.Ref.Type}}
    {{- $ref:=$v1.Ref}}{{- $var:=batch (firstLower $ref.Role)}}
    var {{$var}} {{$poPkg}}.{{batch $ref.EntityName}}
    {{- $r:="result"}}
    {{- if not (eq (firstUpper $v1.Ref.Role) $v1.Ref.EntityName)}}
      {{- $r = printf "%s.%s%sTo%s%s(result)" $basePkg (firstUpper $v1.Ref.Role) (batch $v1.Ref.AttributeID) $v1.Ref.EntityName (batch $v1.Ref.AttributeID)}}
    {{- end}}
	if {{$var}}, err = get{{batch $ref.EntityName}}ByIDsIfRelated(ctx, r.{{firstLower $ref.EntityName}}DAO, {{$r}}, rq); err != nil {
		return nil, 0, err
	}
  {{- end}}
{{- end}}
    builder := {{$potoePkg}}.NewBuilder(rq,
        {{$potoePkg}}.With{{firstUpper (batch $k)}}({{$basePkg}}.To{{firstUpper $v.IDType}}KeyMap(result, 
            func(i interface{}) string { return i.({{$basePkg}}.{{$k}}ID).Get{{$k}}ID() })),
{{- range $v1:=$v.Columns}}
  {{- if $v1.Ref.Type}}
    {{- $var:=batch (firstLower $v1.Ref.Role)}}
    {{- $attrName:= printf "%s%s" $v1.Ref.EntityName $v1.Ref.AttributeID}}
        {{$potoePkg}}.With{{batch $v1.Ref.EntityName}}({{$basePkg}}.To{{firstUpper $v1.AttributeType}}KeyMap({{$var}}, 
            func(i interface{}) string { return i.({{$basePkg}}.{{$attrName}}).Get{{$attrName}}() })),
  {{- end}}
{{- end}}
        {{$potoePkg}}.WithFields(r.{{firstLower $k}}DAO.BaseFields()))
	builder.Build()
    return builder.Get{{firstUpper (batch $k)}}(), total, err
}

func need{{$k}}IDs(req {{$reqPkg}}.Get{{batch $k}}Pure) bool {
{{- range $k1,$v1 := $v.Attrs}}
{{- if $v1.Fuzzy}}
    if !{{$basePkg}}.IsZero(req.Get{{$v1.AttributeName}}()) {
        return true
    } 
{{- end}}
{{- if or $v1.Exact $v1.Batch}}
    {{- $f:= printf "%s%s%s" "Get" $k $v1.AttributeName}}
    {{- if $v1.Ref.Type}}
        {{- $f = printf "%s%s%s" "Get" (firstUpper $v1.Ref.Role) $v1.Ref.AttributeID}}
        {{- if $v1.Batch}}{{- $f = batch $f}}{{- end}}
    {{- end}} 
    if !{{$basePkg}}.IsZero(req.{{$f}}()) {
        return true
    }
{{- end}}
{{- end}}
	return false
}

func get{{$k}}IDsByRef(ctx context.Context, dao {{$ppPkg}}.{{$k}}DAO, pure {{$reqPkg}}.Get{{batch $k}}Pure) ({{$basePkg}}.{{$k}}IDs, bool, error) {
    if !need{{$k}}IDs(pure) {
        return {{$zeroPkg}}.{{$k}}IDs{}, false, nil
    }
    {{- $var:=firstLower (batch $k)}}
    var {{$var}} {{$poPkg}}.{{batch $k}}
    var err error
    {{$var}}, _, err = dao.Get(ctx, struct{
        zero.PageOrderOperator
        {{$reqPkg}}.Get{{batch $k}}Pure
    }{Get{{batch $k}}Pure:pure})
    if len({{$var}}.Get{{$k}}IDs()) == 0 || err != nil{
        return {{$zeroPkg}}.{{$k}}IDs{}, true, nil
    }
    return {{$zeroPkg}}.{{$k}}IDs{Value: {{$var}}.Get{{$k}}IDs()}, false, nil
}

func get{{batch $k}}ByIDsIfRelated(ctx context.Context, dao {{$ppPkg}}.{{$k}}DAO, ids {{$basePkg}}.{{$k}}IDs, fields {{$basePkg}}.Fields)({{$poPkg}}.{{batch $k}}, error){
    if !dao.IsRelated(fields) {
		return nil, nil
	}
	return get{{batch $k}}ByIDs(ctx, dao, ids)
}

func get{{batch $k}}ByIDs(ctx context.Context, dao {{$ppPkg}}.{{$k}}DAO, ids {{$basePkg}}.{{$k}}IDs) ({{$poPkg}}.{{batch $k}}, error) {
	{{- $result:=firstLower (batch $k)}}
    var {{$result}} po.{{batch $k}}
	var err error
	if {{$result}}, _, err = dao.Get(ctx, struct {
{{- range $k1,$v1 := $v.Attrs}}
{{- if $v1.Fuzzy}}
    {{$zeroPkg}}.{{$v1.AttributeName}}
{{- end}}
{{- if and (or $v1.Exact $v1.Batch)}}
{{- if (eq $v1.AttributeName "IDs")}}
    {{$basePkg}}.{{$k}}{{$v1.AttributeName}}
{{- else}}
{{- $attrName:=printf "%s%s" $k $v1.AttributeName}}
{{- if $v1.Ref.Type}}
{{- $attrName = printf "%s%s" (firstUpper $v1.Ref.Role) $v1.Ref.AttributeID}}
{{- if $v1.Batch}}{{- $attrName = batch $attrName}}{{- end}}
{{- end}}
    {{$zeroPkg}}.{{$attrName}}
{{- end}}
{{- end}}
{{- end}}
	{{$zeroPkg}}.PageOrderOperator
}{ {{$k}}IDs: ids }); err != nil {
		return nil, err
	}
	return {{$result}}, err
}

{{- range $k1,$v1 := $v.Attrs}}
  {{- if not $v1.Batch}}
    {{- if eq (print $v1.Ref.Type) "YAM_A"}}

func get{{batch $k}}By{{$v1.AttributeName}}sIfRelated(ctx context.Context, dao {{$ppPkg}}.{{$k}}DAO, ids {{$basePkg}}.{{batch $v1.AttributeName}}, fields {{$basePkg}}.Fields)({{$poPkg}}.{{$k}}s, error){
    if !dao.IsRelated(fields) {
		return nil, nil
	}
	return get{{$k}}sBy{{$v1.AttributeName}}s(ctx, dao, ids)
}

func get{{batch $k}}By{{$v1.Ref.EntityName}}{{batch $v1.Ref.AttributeID}}(ctx context.Context, dao {{$ppPkg}}.{{$k}}DAO, ids {{$basePkg}}.{{batch $v1.AttributeName}})({{$poPkg}}.{{$k}}s, error) {
	var {{firstLower $k}}s po.{{$k}}s
	var err error
	if {{firstLower $k}}s, _, err = dao.Get(ctx, struct {
      {{- range $k2,$v2 := $v.Attrs}}
        {{- if $v2.Fuzzy}}
    {{$zeroPkg}}.{{$v2.AttributeName}}
        {{- end}}
        {{- if and (or $v2.Exact $v2.Batch)}}
          {{- $attrName:= printf "%s%s" $k $v2.AttributeName}}
          {{- if $v2.Ref.Type}}{{- $attrName = $v2.AttributeName}}{{- end}}
          {{- if (eq $v2.AttributeName (batch $v1.AttributeName))}}
    {{$basePkg}}.{{$attrName}}
          {{- else}}
    {{$zeroPkg}}.{{$attrName}}
          {{- end}}
        {{- end}}
      {{- end}}
	{{$zeroPkg}}.PageOrderOperator
}{ {{batch $v1.AttributeName}}: ids }); err != nil {
		return nil, err
	}
	return {{firstLower $k}}s, err
}
    {{- end}}
  {{- end}}
{{- end}}
{{- end}}

`

var repoTempl = template.Must(template.New("repo").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch,
		"entities":   entity.RefEntities}).
	Parse(repotpl))

func WriteRepo(writer io.Writer, domain *entity.Group) {
	_ = repoTempl.Execute(writer, domain)
}
