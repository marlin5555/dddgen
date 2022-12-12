package entity

import (
	"fmt"
	"go/ast"
	"sort"

	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	utilf "github.com/marlin5555/dddgen/pkg/template/util"
)

type DDDTagSetting map[string]string

type DB struct {
	DataBaseName string
}
type Groups struct {
	*ShareColumn
	GroupMap map[string]*Group
}

type ShareColumn struct {
	*Common
	Imports             []string
	ShareCombinationMap map[string]*AttrCombination
}

func (sc ShareColumn) GetGroup() *Group {
	m := map[string]*Entity{}
	for k, combination := range sc.ShareCombinationMap {
		cs := map[string]*Attribute{}
		for _, column := range combination.Columns {
			cs[column.AttributeName] = column
		}
		m[k] = &Entity{
			Attrs:           cs,
			AttrCombination: combination,
		}
	}
	return &Group{
		Name:          util.CommonFile,
		Common:        sc.Common,
		EntityImports: sc.Imports,
		EntityMap:     m,
	}
}

type PackageAndPath struct {
	Package string
	Path    string
}

type Common struct {
	Module           string
	Base             PackageAndPath
	Entity           PackageAndPath
	Req              PackageAndPath
	Rsp              PackageAndPath
	Persistence      PackageAndPath
	PO               PackageAndPath
	POToEntity       PackageAndPath
	Sql              PackageAndPath
	Zero             PackageAndPath
	PortsPersistence PackageAndPath
	PortsRepo        PackageAndPath
	ImplRepo         PackageAndPath
	Factory          PackageAndPath
	Log              PackageAndPath
	Config           PackageAndPath
}

type Group struct {
	*Common
	Name          string
	EntityImports []string
	PoImports     []string
	EntityMap     map[string]*Entity
	FullEntities  map[string]*Entity
}

type AttrCombination struct {
	IsTable      bool
	TypeName     string // e.g. Account
	Columns      []*Attribute
	Combinations []*AttrCombination
}

func (a AttrCombination) GetIDType() string {
	for _, column := range a.Columns {
		if column.AttributeName == "ID" {
			return column.AttributeType
		}
	}
	return ""
}

type Entity struct {
	*DB
	*Group
	TableName string // e.g. account (auto gen)/ t_account
	IDType    string // e.g. string uint32 ...
	Attrs     map[string]*Attribute
	*AttrCombination
}

type Attribute struct {
	Shared        bool // 是否是共享的 column
	Fuzzy         bool // 是否是扩展的 Fuzzy attribute
	Batch         bool // 是否是扩展的 Batch attribute
	Exact         bool // 是否支持 Exact query
	Upsert        bool // 是否支持 Upsert 标签
	Insert        bool // 是否支持 Insert 标签
	Update        bool // 是否支持 Update 标签
	Simple        bool
	Detail        bool
	IDFlag        IDFlag
	Ref           Ref           // <本实体>.<外键> 引用   <其他实体>.<主键>
	Fer           []Ref         // <本实体>.<主键> 被引用 <其他实体>.<外键>
	ColumnName    string        // e.g. id from gorm column tag
	ColumnType    string        // e.g. varchar(36)
	AttributeName string        // e.g. ID
	OrigAttrName  string        // e.g. ID
	AttributeType string        // e.g. string / time.Time
	Import        string        // e.g. time
	Settings      DDDTagSetting // e.g. exact/fuzzy
	Gorm          string        // gorm info
	Obj           *ast.Object   // obj
}

func (a *Attribute) GetFuzzy(belongEntity *Entity) *Attribute {
	return &Attribute{
		Shared:        a.Shared,
		Fuzzy:         true,
		ColumnName:    a.ColumnName,
		ColumnType:    a.ColumnType,
		AttributeName: fuzzy(a.AttributeName, belongEntity.TypeName),
		AttributeType: a.AttributeType,
		Import:        a.Import,
		Settings:      a.Settings,
		Gorm:          a.Gorm,
		Obj:           a.Obj,
	}
}

func fuzzy(aName, eName string) string {
	return fmt.Sprintf("%s%s%s", util.Fuzzy, eName, aName)
}

func (a *Attribute) GetBatch() *Attribute {
	return &Attribute{
		Shared:        a.Shared,
		Fuzzy:         a.Fuzzy,
		Batch:         true,
		ColumnName:    a.ColumnName,
		ColumnType:    a.ColumnType,
		AttributeName: utilf.Batch(a.AttributeName),
		OrigAttrName:  a.AttributeName,
		AttributeType: a.AttributeType,
		Import:        a.Import,
		Ref:           a.Ref,
		Fer:           a.Fer,
		Settings:      a.Settings,
		Gorm:          a.Gorm,
		Obj:           a.Obj,
	}
}

func (a *Attribute) SetExact() *Attribute {
	a.Exact = true
	return a
}

func (a *Attribute) SetUpsert() *Attribute {
	a.Upsert = true
	return a
}

func (a *Attribute) SetUpdate() *Attribute {
	a.Update = true
	return a
}

func (a *Attribute) SetInsert() *Attribute {
	a.Insert = true
	return a
}

func (a *Attribute) SetSimple() *Attribute {
	a.Simple = true
	return a
}

func (a *Attribute) SetDetail() *Attribute {
	a.Detail = true
	return a
}

func (a *Attribute) SetRef(t RefType, ref string, role string) *Attribute {
	a.Ref = BuildRef(t, ref, role)
	return a
}

func (a *Attribute) SetIDFlag(t IDFlagType) *Attribute {
	a.IDFlag = BuildIDFlag(t, a.AttributeName, a.AttributeType)
	return a
}

func (a *Attribute) AppendFer(t RefType, fer []string) *Attribute {
	var r []Ref
	for _, s := range fer {
		r = append(r, BuildRef(t, s, ""))
	}
	a.Fer = append(a.Fer, r...)
	return a
}

func RefEntities(group *Group) []string {
	var arrays []string
	for k, e := range group.EntityMap {
		if k != "" {
			arrays = append(arrays, k)
		}
		for _, a := range e.Attrs {
			if a.Ref.EntityName != "" {
				arrays = append(arrays, a.Ref.EntityName)
			}
		}
	}
	return distinct(arrays)
}

func distinct(input []string) []string {
	m := map[string]interface{}{}
	var r []string
	for _, e := range input {
		if _, ok := m[e]; !ok {
			m[e] = nil
			r = append(r, e)
		}
	}
	sort.StringSlice(r).Sort()

	return r
}
