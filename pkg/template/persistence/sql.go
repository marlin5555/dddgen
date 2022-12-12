package persistence

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
	"io"
	"text/template"
)

// table column gen interface template
var sqltpl = `// Package {{.Sql.Package}} Code generated, DO NOT EDIT.
package {{.Sql.Package}}
import (
    "context"
    "fmt"
    "reflect"
    "time"

    "{{.Module}}{{.Base.Path}}"
    "{{.Module}}{{.Req.Path}}"
    "{{.Module}}{{.PortsPersistence.Path}}"
    "{{.Module}}{{.PO.Path}}"

    "gorm.io/gorm"
)

{{- $e:=.TypeName}}
{{- $po:=.PO.Package}}
// New{{$e}}DAO 构造方法
func New{{$e}}DAO(db *gorm.DB) {{.PortsPersistence.Package}}.{{$e}}DAO {
    return &{{$e}}DAO{
        db: db,
    }
}

// {{$e}}DAO {{$e}} 数据访问对象
type {{$e}}DAO struct {
    db *gorm.DB
}

// Transaction 事务执行
func (e {{$e}}DAO) Transaction(fun func(tx *gorm.DB) error) error {
    return e.db.Transaction(func(tx *gorm.DB) (err error) {
        return fun(tx)
    })
}

// IsRelated check fields is related EventBusDAO
func (e {{$e}}DAO) IsRelated(fields {{.Base.Package}}.Fields) bool {
	return checkFields(fields, {{.PO.Package}}.{{$e}}{}.TableName())
}

// BaseFields {{$e}}DAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e {{$e}}DAO) BaseFields() {{.Base.Package}}.Fields {
	return {{firstLower $e}}BaseFields{}
}

type {{firstLower $e}}BaseFields struct{}
// GetFieldNames return event set dao base fields
func ({{firstLower $e}}BaseFields) GetFieldNames() []string {
	return []string{
{{- range .Attrs}}
{{- if .Simple}}
		fmt.Sprintf(COLUMN_PATTERN, {{$po}}.{{$e}}{}.TableName(), "{{.ColumnName}}"),
{{- end}}
{{- end}}
	}
}

func (e {{$e}}DAO) listBaseDB(ctx context.Context, r {{.Req.Package}}.Get{{batch $e}}PO) *gorm.DB {
    eT := newTable({{.PO.Package}}.{{$e}}{}.TableName())
    db := e.db.WithContext(ctx).Model({{.PO.Package}}.{{$e}}{}).Where(eT.col("deleted_at").notDelete())
    mdb := mydb{db: db}

    for _, t := range []tuple{
{{- range $k,$v := .Attrs}}
{{- $attrName:= printf "%s%s" $e $v.AttributeName}}
{{- if $v.Ref.Type}}{{- $attrName = printf "%s%s" (firstUpper $v.Ref.Role) $v.Ref.AttributeID}}{{- end}}
{{- if $v.Batch}}
        {predicate: eT.col("{{$v.ColumnName}}").in(), param: r.Get{{$attrName}}()},
{{- end}}
{{- if $v.Fuzzy}}
        {predicate: eT.col("{{$v.ColumnName}}").like(), param: r.Get{{$v.AttributeName}}()},
{{- end}}
{{- if $v.Exact}}
        {predicate: eT.col("{{$v.ColumnName}}").eq(), param: r.Get{{$attrName}}()},
{{- end}}
{{- end}}
    } {
        mdb.appendWhere(t)
    }
    return mdb.db
}

// Get 通用 get 方法
func (e {{$e}}DAO) Get(ctx context.Context, r {{.Req.Package}}.Get{{batch $e}}PO) (po.{{batch $e}}, uint32, error) {
    var total int64
    if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    var result []{{.PO.Package}}.{{$e}}
    if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
        return nil, 0, err
    }
    return result, uint32(total), nil

}
func (e {{$e}}DAO) Update(ctx context.Context, r {{.Req.Package}}.Update{{$e}}Req) error {
    return e.db.WithContext(ctx).Updates(convertUpdate{{$e}}Req(r)).Error
}

func (e {{$e}}DAO) Delete(ctx context.Context, r {{.Req.Package}}.Delete{{$e}}Req) error {
    return e.db.WithContext(ctx).Updates(convertDelete{{$e}}Req(r)).Error
}

func (e {{$e}}DAO) Create(ctx context.Context, r {{.Req.Package}}.Create{{$e}}Req) ({{.IDType}},error) {
	pobj := convertCreate{{$e}}Req(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().({{.IDType}}), err
	}
	return pobj.ID, nil
}

func convertUpdate{{$e}}Req(r {{.Req.Package}}.Update{{$e}}Req) {{.PO.Package}}.{{$e}} {
    return {{.PO.Package}}.{{$e}}{
{{- range $k,$v := .Attrs}}
{{- if or $v.Update $v.Upsert}}
{{- $f:=printf "%s%s%s" "Get" $e $v.AttributeName}}
{{- if $v.Ref.Type}}
{{- $f = printf "%s%s%s" "Get" (firstUpper $v.Ref.Role) $v.Ref.AttributeID}}
{{- end}}
        {{$v.AttributeName}}: r.{{$f}}(),
{{- end}}
{{- end}}
        Updater: r.GetRtx(),
    }
}

func convertDelete{{$e}}Req(r {{.Req.Package}}.Delete{{$e}}Req) {{.PO.Package}}.{{$e}} {
    return {{.PO.Package}}.{{$e}}{
        ID : r.Get{{$e}}ID(),
        DeletedAt : time.Now().Unix(),
        Updater: r.GetRtx(),
    }
}

func convertCreate{{$e}}Req(r {{.Req.Package}}.Create{{$e}}Req) {{.PO.Package}}.{{$e}} {
    return {{.PO.Package}}.{{$e}}{
{{- range $k,$v := .Attrs}}
  {{- if or $v.Insert $v.Upsert}}
    {{- $f:=printf "%s%s%s" "Get" $e $v.AttributeName}}
    {{- if $v.Ref.Type}}
      {{- $f = printf "%s%s%s" "Get" (firstUpper $v.Ref.Role) $v.Ref.AttributeID}}
    {{- end}}
        {{$v.AttributeName}}: r.{{$f}}(),
  {{- end}}
  {{- if and $v.IDFlag.Type (auto $v.IDFlag.Type)}}
    {{- if eq $v.IDFlag.AttrType "string"}}
        {{$v.AttributeName}}:uuid(),
    {{- end}}
  {{- end}}
{{- end}}
        Creator: r.GetRtx(),
    }
}
`

var sqlTempl = template.Must(template.New("persistence-sql").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch, "auto": entity.Auto}).
	Parse(sqltpl))

func WriteSQL(writer io.Writer, entity *entity.Entity) {
	_ = sqlTempl.Execute(writer, entity)
}

var sqlcommontpl = `// Package {{.Sql.Package}} Code generated, DO NOT EDIT.
package {{.Sql.Package}}
import(
    "fmt"
    "reflect"
	"strings"
    "sync/atomic"

    "{{.Module}}{{.Base.Path}}"
    config "{{.Module}}{{.Config.Path}}"

    guuid "github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func uuid() string{
    return guuid.New().String()
}

// NewSQLiteDB new sqlLite client
func NewSQLiteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared&loc=Asia%2fShanghai"), &gorm.Config{})
	if err != nil {
		panic("Connect to sqlite error: " + err.Error())
	}
	return db
}

// NewMySQLDB new mysql db
func NewMySQLDB(config config.DBConfig) *gorm.DB {
	// mysql config
	username := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	timeout := config.Timeout
	dbname := config.DBName
	// dsn param
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?collation=utf8mb4_0900_as_cs&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, dbname, timeout)
	// connect to mysql and get db client
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Connect to mysql error: " + err.Error())
	}
	return db
}

const (
    UnDeleted = 0 // 软删除：未删除
    COLUMN_PATTERN = "%s.%s" // 用于构造 table.column
)

func needOrder(o {{.Base.Package}}.Order) bool {
    return o.OrderStr() != ""
}

func needPage(pq {{.Base.Package}}.PageQuery) bool {
    if pq.GetOffset() < 0 {
        return false
    }
    if pq.GetLimit() <= 0 {
        return false
    }
    return true
}

// pageQuery 进行分页查询，结果放到dest中
func pageQuery(d *gorm.DB, pgo {{.Base.Package}}.PageOrder, dest interface{}) error {
    db := d
    if needPage(pgo) {
        db = db.Offset(pgo.GetOffset()).Limit(pgo.GetLimit())
    }
    if needOrder(pgo) {
        db = db.Order(pgo.OrderStr())
    }
    return db.Find(dest).Error
}

type table struct {
    tableName   string
    joinCounter atomic.Value
}

func newTable(name string) *table {
    var v = atomic.Value{}
    v.Store(0)
    return &table{
        tableName:   name,
        joinCounter: v,
    }
}

func (t *table) col(col string) column {
    return column{
        table:      t,
        columnName: col,
    }
}

type column struct {
    *table
    columnName string
}

func (c column) notDelete() string {
    return fmt.Sprintf("%s.%s = %d", c.tableName, c.columnName, UnDeleted)
}

func (c column) eq() predicate {
    return func(param interface{}) (Category, string, interface{}) {
        if notZero(param) {
            return WHERE, fmt.Sprintf("%s.%s = ?", c.tableName, c.columnName), param
        }
        return WHERE, "", nil
    }
}

func (c column) in() predicate {
    return func(param interface{}) (Category, string, interface{}) {
        switch reflect.TypeOf(param).Kind() {
        case reflect.Array, reflect.Slice:
            if notZero(param) {
                return WHERE, fmt.Sprintf("%s.%s IN ?", c.tableName, c.columnName), param
            }
        }
        return WHERE, "", nil
    }
}

func (c column) like() predicate {
    return func(param interface{}) (Category, string, interface{}) {
        switch reflect.TypeOf(param).Kind() {
        case reflect.String:
            if notZero(param) {
                return WHERE, fmt.Sprintf("%s.%s like '%%%s%%'", c.tableName, c.columnName, param), nil
            }
        }
        return WHERE, "", nil
    }
}

type mydb struct {
    db *gorm.DB
}

type Category int

const (
    WHERE Category = 1
    JOIN  Category = 2
)

func (d mydb) appendWhere(t tuple) {
    var c Category
    var q string
    var args interface{}
    if c, q, args = t.predicate(t.param); q == "" {
        return
    }
    switch c {
    case WHERE:
        if args == nil {
            d.db = d.db.Where(q)
            return
        }
        d.db = d.db.Where(q, args)
        return
    case JOIN:
        d.db.Joins(q, args)
    }
}

type tuple struct {
    param     interface{}
    predicate predicate
}

type predicate = func(param interface{}) (Category, string, interface{})

func notZero(v interface{}) bool {
    return !reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

func checkFields(fields base.Fields, table string) bool {
	for _, field := range fields.GetFieldNames() {
		if table == strings.Split(field, ".")[0] {
			return true
		}
	}
	return false
}
`

var sqlCommonTempl = template.Must(template.New("persistence-sql-common").Parse(sqlcommontpl))

func WriteCommonSQL(writer io.Writer, common *entity.Common) {
	_ = sqlCommonTempl.Execute(writer, common)
}
