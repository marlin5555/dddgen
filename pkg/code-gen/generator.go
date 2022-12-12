package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	"golang.org/x/tools/go/packages"
	"gorm.io/gorm/schema"
)

func main() {
	_ = os.MkdirAll(util.EntityPath, 0777)
	_ = os.MkdirAll(util.EntityBasePath, 0777)

	path := "./meta/example-1/entity"
	gendir := "gen/example-1"
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedImports | packages.NeedName |
		packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax}, path)
	if err != nil {
		panic(err)
	}
	g := Generator{}
	g.addPkg(pkgs[0], gendir)

	g.build()
	g.print()

	g.generateEntityInterface(gendir)
	g.generateFactory(gendir)
	g.generateEntityZero(gendir)
	g.generateEntity(gendir)
	g.generateReq(gendir)
	g.generateRsp(gendir)
	g.generatePO(gendir)
	g.generatePersistenceConstant(gendir)
	g.generatePersistencePorts(gendir)
	g.generatePOToEntity(gendir)
	g.generateSQL(gendir)
	g.generateRepoPorts(gendir)
	g.generateRepo(gendir)
}

func toName(p *packages.Package) []string {
	return p.Types.Scope().Names()
}

func getWriter(genDir, pathBase string, fileName string) (*bufio.Writer, *os.File) {
	p := path.Join(genDir, pathBase)
	os.MkdirAll(p, 0777)
	f, err := os.Create(path.Join(p, fileName) + ".go")
	if err != nil {
		log.Printf("%+v", err)
	}
	return bufio.NewWriter(f), f
}

func (g *Generator) print() {
	fmt.Print("\n\n====== attributes ======")
	for k, v := range g.pkg.attributes {
		for k1, v1 := range v {
			fmt.Printf("\n[%s][%s][%s]: ref=[%+v]; fer=%+v", k, k1, v1.AttributeName, v1.Ref, v1.Fer)
		}
	}
	fmt.Print("\n\n====== groups ======")
	for k, v := range g.pkg.groups.GroupMap {
		for k1, v1 := range v.EntityMap {
			for k2, attribute := range v1.Attrs {
				fmt.Printf("\n[%s][%s][%s]: ref=[%+v]; fer=%+v", k, k1, k2, attribute.Ref, attribute.Fer)
			}
		}
	}
}

// build produces the String method for the named type.
func (g *Generator) build() {
	for _, file := range g.pkg.files {
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
		}
	}
	fillTable(g.pkg.attributes, g.pkg.tables, g.pkg.attrCombins,
		g.pkg.groups, g.pkg.common, g.pkg.fullEntities)

}

func fillPackageAndPath(groups *entity.Groups, common *entity.Common) {
	groups.ShareColumn.Common = common
}

func fillUpsertAttribute(attribute *entity.Attribute) {
	if _, ok := attribute.Settings[util.UPSERT]; ok {
		attribute.SetUpsert()
	}
}
func fillInsertAttribute(attribute *entity.Attribute) {
	if _, ok := attribute.Settings[util.INSERT]; ok {
		attribute.SetInsert()
	}
}
func fillUpdateAttribute(attribute *entity.Attribute) {
	if _, ok := attribute.Settings[util.UPDATE]; ok {
		attribute.SetUpdate()
	}
}
func fillExactAttribute(attribute *entity.Attribute) {
	if _, ok := attribute.Settings[util.EXACT]; ok {
		attribute.SetExact()
	}
}
func fillSimpleAttribute(attribute *entity.Attribute) {
	if _, ok := attribute.Settings[util.SIMPLE]; ok {
		attribute.SetSimple()
	}
}
func fillDetailAttribute(attribute *entity.Attribute) {
	if _, ok := attribute.Settings[util.DETAIL]; ok {
		attribute.SetDetail()
	}
}

func fillRefAttribute(attribute *entity.Attribute) {
	var t entity.RefType
	var finalRef string
	var role string
	// 关系中的角色
	if r, ok := attribute.Settings[util.ROLE]; ok {
		role = r
	}
	for _, k := range entity.RefRelations {
		if ref, ok := attribute.Settings[k]; ok {
			t = entity.ToRefType(k)
			finalRef = ref
			break
		}
	}
	if t != 0 {
		attribute.SetRef(t, finalRef, role)
	}
}

func fillFefAttribute(attribute *entity.Attribute) {
	if fer, ok := attribute.Settings[util.FER]; ok {
		attribute.AppendFer(entity.FER, strings.Split(fer, ","))
	}
	if fer, ok := attribute.Settings[util.MUSTA]; ok {
		attribute.AppendFer(entity.MUSTA, strings.Split(fer, ","))
	}
	if fer, ok := attribute.Settings[util.MAYA]; ok {
		attribute.AppendFer(entity.MAYA, strings.Split(fer, ","))
	}
	if fer, ok := attribute.Settings[util.MUSTS]; ok {
		attribute.AppendFer(entity.MUSTS, strings.Split(fer, ","))
	}
	if fer, ok := attribute.Settings[util.MAYS]; ok {
		attribute.AppendFer(entity.MAYS, strings.Split(fer, ","))
	}
	if fer, ok := attribute.Settings[util.SMAYS]; ok {
		attribute.AppendFer(entity.SMAYS, strings.Split(fer, ","))
	}
	if fer, ok := attribute.Settings[util.SMUSTS]; ok {
		attribute.AppendFer(entity.SMUSTS, strings.Split(fer, ","))
	}
}

func fillIDFlag(attribute *entity.Attribute) {
	if idflag, ok := attribute.Settings[util.ID]; ok {
		attribute.SetIDFlag(entity.ToIDFlagType(strings.ToUpper(idflag)))
	}
}

func fillTable(attrs map[string]map[string]*entity.Attribute, entities map[string]*entity.Entity,
	combinations map[string]*entity.AttrCombination, groups *entity.Groups,
	common *entity.Common, fullEntity map[string]*entity.Entity) {
	fillPackageAndPath(groups, common)
	for _, table := range entities {
		if _, ok := attrs[table.TableName]; !ok {
			attrs[table.TableName] = map[string]*entity.Attribute{}
		}
		fillCombination(attrs[table.TableName], table.AttrCombination)
	}
	shareColumn := groups.ShareColumn
	for k, v := range combinations {
		if !v.IsTable {
			shareColumn.ShareCombinationMap[k] = v
		}
	}

	for k, v := range attrs {
		ent := entities[k]
		if ent.Group == nil {
			fmt.Printf("got error About[%s]", k)
			continue
		}
		group := groups.GroupMap[ent.Group.Name]
		entityMap := group.EntityMap
		if _, ok := entityMap[ent.TypeName]; !ok {
			entityMap[ent.TypeName] = ent
		}
		for _, attribute := range v {
			if _, ok := ent.Attrs[attribute.AttributeName]; !ok {
				ent.Attrs[attribute.AttributeName] = attribute
				fillExactAttribute(attribute)
				fillSimpleAttribute(attribute)
				fillDetailAttribute(attribute)
				fillUpsertAttribute(attribute)
				fillUpdateAttribute(attribute)
				fillInsertAttribute(attribute)
				fillRefAttribute(attribute)
				fillFefAttribute(attribute)
				fillIDFlag(attribute)
			}
			if _, ok := attribute.Settings[util.FUZZY]; ok {
				fuzzyAttr := attribute.GetFuzzy(ent)
				ent.Attrs[fuzzyAttr.AttributeName] = fuzzyAttr
			}
			if _, ok := attribute.Settings[util.BATCH]; ok {
				batchAttr := attribute.GetBatch()
				ent.Attrs[batchAttr.AttributeName] = batchAttr
			}
			group.PoImports = append(group.PoImports, attribute.Import)
			if attribute.Shared {
				shareColumn.Imports = append(shareColumn.Imports, attribute.Import)
			} else {
				group.EntityImports = append(group.EntityImports, attribute.Import)
			}
		}
	}
	shareColumn.Imports = trimArray(shareColumn.Imports)
	for _, group := range groups.GroupMap {
		group.EntityImports = trimArray(group.EntityImports)
		group.PoImports = trimArray(group.PoImports)
	}
	for _, group := range groups.GroupMap {
		for ename, entity := range group.EntityMap {
			fullEntity[ename] = entity
		}
	}
}

func trimArray(input []string) []string {
	resmap := map[string]interface{}{}
	res := []string{}
	for _, s := range input {
		if s == "" {
			continue
		}
		if _, ok := resmap[s]; ok {
			continue
		}
		resmap[s] = nil
		res = append(res, s)
	}
	return res
}

func fillCombination(m map[string]*entity.Attribute, cc *entity.AttrCombination) {
	for _, column := range cc.Columns {
		if _, ok := m[column.ColumnName]; !ok {
			m[column.ColumnName] = column
		}
	}
	for _, combination := range cc.Combinations {
		fillCombination(m, combination)
	}
}

func (f *_File) genDecl(node ast.Node) bool {
	ng, ok := node.(*ast.GenDecl)
	if !ok {
		return true
	}
	if ng.Tok != token.TYPE {
		return true
	}

	if len(ng.Specs) != 1 {
		log.Printf("****** not support type(xxx) ")
		return true
	}
	typ, ok := ng.Specs[0].(*ast.TypeSpec)
	if !ok {
		log.Printf("****** not support %+v ", typ)
		return true
	}
	attrs, ok := typ.Type.(*ast.StructType)
	if !ok {
		log.Printf("****** not support %+v ", typ)
		return true
	}
	structName := typ.Name.Name
	tableName, groupName := parseColumnCombinationComment(ng.Doc)(structName)
	if _, ok = f.pkg.attrCombins[structName]; !ok {
		f.pkg.attrCombins[structName] = &entity.AttrCombination{
			IsTable:      tableName != "",
			TypeName:     structName,
			Columns:      []*entity.Attribute{},
			Combinations: []*entity.AttrCombination{},
		}
	}
	cc := f.pkg.attrCombins[structName]
	for _, field := range attrs.Fields.List {
		parseColumn(field, f.pkg.attrCombins, cc)
	}

	if tableName != "" {
		if _, ok := f.pkg.tables[tableName]; ok {
			log.Printf("****** duplcate table %s", tableName)
			return false
		}
		var v *entity.Group
		if groupName != "" {
			if _, ok := f.pkg.groups.GroupMap[groupName]; !ok {
				f.pkg.groups.GroupMap[groupName] = &entity.Group{
					Name:         groupName,
					EntityMap:    map[string]*entity.Entity{},
					Common:       f.pkg.common,
					FullEntities: f.pkg.fullEntities,
				}
			}
			v = f.pkg.groups.GroupMap[groupName]
		}
		f.pkg.tables[tableName] = &entity.Entity{
			DB:              nil,
			Attrs:           map[string]*entity.Attribute{},
			Group:           v,
			TableName:       tableName,
			IDType:          cc.GetIDType(),
			AttrCombination: &entity.AttrCombination{TypeName: structName},
		}
		f.pkg.tables[tableName].AttrCombination = f.pkg.attrCombins[structName]
	}
	return false
}

func parseColumn(field *ast.Field, ccmap map[string]*entity.AttrCombination, cc *entity.AttrCombination) {
	if len(field.Names) > 1 {
		log.Printf("22222")
		return
	}
	if len(field.Names) == 0 {
		typ := field.Type.(*ast.Ident).Name
		if _, ok := ccmap[typ]; !ok {
			ccmap[typ] = &entity.AttrCombination{
				TypeName:     typ,
				Columns:      []*entity.Attribute{},
				Combinations: []*entity.AttrCombination{},
			}
		}
		cc.Combinations = append(cc.Combinations, ccmap[typ])
		return
	}
	colName, colType := parseColumnTag(field.Tag)

	attrType, importStr := parseAttributeType(field.Type)
	cc.Columns = append(cc.Columns, &entity.Attribute{
		Shared:        !cc.IsTable,
		ColumnName:    colName,
		ColumnType:    colType,
		AttributeName: field.Names[0].Name,
		AttributeType: attrType,
		Import:        importStr,
		Gorm:          parseGorm(field.Tag),
		Settings:      parseDDDTag(field.Tag),
		Obj:           field.Names[0].Obj,
	})

}

func parseAttributeType(typ ast.Expr) (string, string) {
	if i, ok := typ.(*ast.Ident); ok {
		return i.Name, ""
	}
	if i, ok := typ.(*ast.SelectorExpr); ok {
		return i.X.(*ast.Ident).Name + "." + i.Sel.Name, i.X.(*ast.Ident).Name
	}
	return "", ""
}

type combinationCommentProc func(string) (string, string)

func parseColumnCombinationComment(comment *ast.CommentGroup) combinationCommentProc {
	if len(comment.List) != 1 {
		println("type must has one comment")
		return nil
	}
	return parseTableName(comment.List[0].Text)
}

func parseTableName(comment string) combinationCommentProc {
	return func(combinationName string) (string, string) {
		s := strings.TrimPrefix(comment, "//")
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, combinationName)
		combinationSetting := schema.ParseTagSetting(s, ";")
		fmt.Printf("%+v", combinationSetting)
		var tableName, groupName string
		tableName, _ = combinationSetting["TABLE"]
		groupName, _ = combinationSetting["GROUP"]
		return strings.TrimSpace(tableName), strings.TrimSpace(groupName)
	}
}

func parseColumnTag(lit *ast.BasicLit) (string, string) {
	v, _ := strconv.Unquote(lit.Value)
	gormTag := reflect.StructTag(v).Get("gorm")
	tagSetting := schema.ParseTagSetting(gormTag, ";")
	return tagSetting["COLUMN"], tagSetting["TYPE"]
}

func parseGorm(lit *ast.BasicLit) string {
	v, _ := strconv.Unquote(lit.Value)
	return reflect.StructTag(v).Get("gorm")
}

func parseDDDTag(lit *ast.BasicLit) map[string]string {
	v, _ := strconv.Unquote(lit.Value)
	dddTag := reflect.StructTag(v).Get("ddd")
	return schema.ParseTagSetting(dddTag, ";")
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *_Package    // _Package we are scanning.

	trimPrefix  string
	lineComment bool
}

type _Package struct {
	name         string
	common       *entity.Common
	groups       *entity.Groups
	attrCombins  map[string]*entity.AttrCombination
	attributes   map[string]map[string]*entity.Attribute
	tables       map[string]*entity.Entity
	fullEntities map[string]*entity.Entity

	files []*_File
}

// _File holds a single parsed file and associated data.
type _File struct {
	fileSet *token.FileSet
	pkg     *_Package // _Package to which this file belongs.
	file    *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string   // Name of the constant type.
	values   []_Value // Accumulator for constant values of that type.

	trimPrefix  string
	lineComment bool
}

// _Value represents a declared constant.
type _Value struct {
	originalName string // The name of the constant.
	name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by _Value.String.
	value  uint64 // Will be converted to int64 when needed.
	signed bool   // Whether the constant is a signed type.
	str    string // The string representation given by the "go/constant" package.
}

func (g *Generator) addPkg(pkg *packages.Package, genDir string) {
	g.pkg = &_Package{
		name:  pkg.Name,
		files: make([]*_File, len(pkg.Syntax)),
		common: &entity.Common{
			Module: path.Join(util.Module, genDir) + "/",
			Base: entity.PackageAndPath{
				Package: util.EntityBasePackage,
				Path:    util.EntityBasePath,
			}, Entity: entity.PackageAndPath{
				Package: util.EntityPackage,
				Path:    util.EntityPath,
			}, Req: entity.PackageAndPath{
				Package: util.EntityReqPackage,
				Path:    util.EntityReqPath,
			}, Rsp: entity.PackageAndPath{
				Package: util.EntityRspPackage,
				Path:    util.EntityRspPath,
			}, Persistence: entity.PackageAndPath{
				Package: util.Persistence,
				Path:    util.PersistencePath,
			}, PO: entity.PackageAndPath{
				Package: util.PoPackage,
				Path:    util.PoPath,
			}, Sql: entity.PackageAndPath{
				Package: util.SqlPackage,
				Path:    util.SqlPath,
			}, Zero: entity.PackageAndPath{
				Package: util.EntityZeroPackage,
				Path:    util.EntityZeroPath,
			}, PortsPersistence: entity.PackageAndPath{
				Package: util.Persistence,
				Path:    util.PersistencePortsPath,
			}, PortsRepo: entity.PackageAndPath{
				Package: util.RepoPackage,
				Path:    util.RepoPortsPath,
			}, ImplRepo: entity.PackageAndPath{
				Package: util.RepoPackage,
				Path:    util.RepoPath,
			}, POToEntity: entity.PackageAndPath{
				Package: util.PoToEntityPackage,
				Path:    util.PoToEntityPath,
			}, Log: entity.PackageAndPath{
				Package: util.LogPackage,
				Path:    util.LogPath,
			}, Factory: entity.PackageAndPath{
				Package: util.FactoryPackage,
				Path:    util.FactoryPath,
			}, Config: entity.PackageAndPath{
				Package: util.ConfigPackage,
				Path:    util.ConfigPath,
			},
		},
		groups: &entity.Groups{
			ShareColumn: &entity.ShareColumn{
				Imports:             []string{},
				ShareCombinationMap: map[string]*entity.AttrCombination{},
			},
			GroupMap: map[string]*entity.Group{},
		},
		attrCombins:  map[string]*entity.AttrCombination{},
		attributes:   map[string]map[string]*entity.Attribute{},
		tables:       map[string]*entity.Entity{},
		fullEntities: map[string]*entity.Entity{},
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &_File{
			fileSet:     pkg.Fset,
			file:        file,
			pkg:         g.pkg,
			trimPrefix:  g.trimPrefix,
			lineComment: g.lineComment,
		}
	}
}
