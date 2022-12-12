// Package db Code generated, DO NOT EDIT.
package db

import (
	"fmt"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	config "github.com/marlin5555/dddgen/gen/example-1/pkg/conf"

	guuid "github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func uuid() string {
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
	UnDeleted      = 0       // 软删除：未删除
	COLUMN_PATTERN = "%s.%s" // 用于构造 table.column
)

func needOrder(o base.Order) bool {
	return o.OrderStr() != ""
}

func needPage(pq base.PageQuery) bool {
	if pq.GetOffset() < 0 {
		return false
	}
	if pq.GetLimit() <= 0 {
		return false
	}
	return true
}

// pageQuery 进行分页查询，结果放到dest中
func pageQuery(d *gorm.DB, pgo base.PageOrder, dest interface{}) error {
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
