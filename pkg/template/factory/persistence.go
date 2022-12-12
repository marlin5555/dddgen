package factory

import (
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"

	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
)

var persistenceTpl = `// Package {{.Factory.Package}} Code generated, DO NOT EDIT.
package {{.Factory.Package}}

import(
    "sync"

    "{{.Module}}{{.PortsPersistence.Path}}"
    "{{.Module}}{{.Sql.Path}}"
    "{{.Module}}{{.PO.Path}}"
    config "{{.Module}}{{.Config.Path}}"
    "{{.Module}}{{.Log.Path}}"
)

{{- $persistPkg:=.PortsPersistence.Package}}
{{- $sqlPkg:=.Sql.Package}}
{{- $poPkg:=.PO.Package}}
// PersistenceFactory 存储层数据访问对象工厂
type PersistenceFactory struct {
{{- range $k,$v:=.Entities}}
    {{firstLower $k}}Dao {{$persistPkg}}.{{$k}}DAO
{{- end}}
}

var (
    dbOnce         sync.Once
    persistFactory PersistenceFactory
)

// DBPersistenceFactory 构造 PersistenceFactory
func DBPersistenceFactory(conf config.Config) *PersistenceFactory {
    dbOnce.Do(func() {
        // 如果没有传入 db config 则使用 SQLite
        mydb := db.NewSQLiteDB()
        if conf.DBConfig().Host != "" {
            mydb = db.NewMySQLDB(conf.DBConfig())
        }
        persistFactory = PersistenceFactory{
{{- range $k,$v:=.Entities}}
            {{firstLower $k}}Dao: {{$sqlPkg}}.New{{$k}}DAO(mydb),
{{- end}}
        }
        var err error
{{- range $k,$v:=.Entities}}
        err = mydb.AutoMigrate(&{{$poPkg}}.{{$k}}{})
        log.InfofWithFuncName("Auto Migrate {{$k}} got err = %+v", err)
{{- end}}
    })
    return &persistFactory
}


{{- range $k,$v:=.Entities}}
// {{$k}}DAO 获取 {{$persistPkg}}.{{$k}}DAO
func (f *PersistenceFactory) {{$k}}DAO() {{$persistPkg}}.{{$k}}DAO {
    if f == nil {
        return nil
    }
    return f.{{firstLower $k}}Dao
}
{{- end}}

`

var persistenceTempl = template.Must(template.New("common-zero").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch,
		"part":       entity.Part}).Parse(persistenceTpl))

func WritePersistence(writer io.Writer, entities map[string]*entity.Entity, common *entity.Common) {
	_ = persistenceTempl.Execute(writer, struct {
		Entities map[string]*entity.Entity
		*entity.Common
	}{Common: common, Entities: entities})
}

var repoTpl = `// Package {{.Factory.Package}} Code generated, DO NOT EDIT.
package {{.Factory.Package}}

import(
    "sync"

    "{{.Module}}{{.PortsRepo.Path}}"
    implrepo "{{.Module}}{{.ImplRepo.Path}}"
)

{{- $persistPkg:=.PortsPersistence.Package}}
{{- $sqlPkg:=.Sql.Package}}
{{- $poPkg:=.PO.Package}}
{{- $repoPkg:=.PortsRepo.Package}}

// RepoFactory 存储层数据访问对象工厂
type RepoFactory struct {
{{- range $k,$v:=.Groups}}
    {{firstLower $k}}Repo {{$repoPkg}}.{{firstUpper $k}}Repository
{{- end}}
}

var (
    repoOnce    sync.Once
    repoFactory RepoFactory
)

// RepositoryFactory 构造 PersistenceFactory
func RepositoryFactory(pF *PersistenceFactory) *RepoFactory {
    repoOnce.Do(func() {
        repoFactory = RepoFactory{
{{- range $k,$v:=.Groups}}
{{- $es:= entities $v}}
            {{firstLower $k}}Repo: implrepo.New{{firstUpper $k}}Repo({{- range $es}}pF.{{firstUpper .}}DAO(),{{- end}}),
{{- end}}
        }
    })
    return &repoFactory
}

{{- range $k,$v:=.Groups}}
func (f *RepoFactory) {{firstUpper $k}}Repo() {{$repoPkg}}.{{firstUpper $k}}Repository{
    if f == nil {
        return nil
    }
    return f.{{firstLower $k}}Repo
}
{{- end}}

`
var repoTempl = template.Must(template.New("common-zero").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch,
		"part":       entity.Part,
		"entities":   entity.RefEntities}).Parse(repoTpl))

func WriteRepo(writer io.Writer, groups map[string]*entity.Group, common *entity.Common) {
	_ = repoTempl.Execute(writer, struct {
		Groups map[string]*entity.Group
		*entity.Common
	}{Common: common, Groups: groups})
}
