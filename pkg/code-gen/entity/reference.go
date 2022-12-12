package entity

import (
	"encoding/json"
	"strings"

	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	utilf "github.com/marlin5555/dddgen/pkg/template/util"
)

type RefType int

const (
	_ RefType = iota
	REF
	FER
	// 1 - 1
	MUSTA
	TSUMA
	MAYA
	YAMA
	// 1 - n
	MUSTS
	TSUMS
	MAYS
	YAMS
	// m - n
	SMAYMUSTS
	SMUSTMAYS
	SMAYS
	SYAMS
	SMUSTS
	STSUMS
)

var Relations = []string{util.REF, util.FER, util.MUSTA, util.TSUMA, util.MAYA, util.YAMA,
	util.MUSTS, util.TSUMS, util.MAYS, util.YAMS,
	util.SMAYMUSTS, util.SMUSTMAYS, util.SMAYS, util.SYAMS, util.SMUSTS, util.STSUMS}
var RefRelations = []string{util.TSUMA, util.YAMA, util.YAMS, util.TSUMS}
var PartFerRelations = []string{util.MUSTA, util.MAYA}
var strToType = map[string]RefType{
	util.REF:       REF,
	util.FER:       FER,
	util.MUSTA:     MUSTA,
	util.TSUMA:     TSUMA,
	util.MAYA:      MAYA,
	util.YAMA:      YAMA,
	util.MUSTS:     MUSTS,
	util.TSUMS:     TSUMS,
	util.MAYS:      MAYS,
	util.YAMS:      YAMS,
	util.SMAYMUSTS: SMAYMUSTS,
	util.SMUSTMAYS: SMUSTMAYS,
	util.SMAYS:     SMAYS,
	util.SYAMS:     SYAMS,
	util.SMUSTS:    SMUSTS,
	util.STSUMS:    STSUMS,
}
var typeTostr = map[RefType]string{
	REF:       util.REF,
	FER:       util.FER,
	MUSTA:     util.MUSTA,
	TSUMA:     util.TSUMA,
	MAYA:      util.MAYA,
	YAMA:      util.YAMA,
	MUSTS:     util.MUSTS,
	TSUMS:     util.TSUMS,
	MAYS:      util.MAYS,
	YAMS:      util.YAMS,
	SMAYMUSTS: util.SMAYMUSTS,
	SMUSTMAYS: util.SMUSTMAYS,
	SMAYS:     util.SMAYS,
	SYAMS:     util.SYAMS,
	SMUSTS:    util.SMUSTS,
	STSUMS:    util.STSUMS}

func (t RefType) String() string {
	if s, ok := typeTostr[t]; ok {
		return s
	}
	return ""
}

func ToRefType(s string) RefType {
	if i, ok := strToType[s]; ok {
		return i
	}
	return 0
}

type Ref struct {
	Type        RefType // 引用类型
	Role        string  // 被引用实体别名，当为空时，用被引用实体名填充
	EntityName  string  // (被)引用实体
	AttributeID string  // (被)引用主(外)键
}

func (r *Ref) String() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

func BuildRef(t RefType, ref string, role string) Ref {
	ss := strings.Split(ref, ".")
	return Ref{Type: t, Role: utilf.IF(role == "", ss[0], role), EntityName: ss[0], AttributeID: ss[1]}
}

func Part(t RefType) bool {
	for _, relation := range PartFerRelations {
		if t.String() == relation {
			return true
		}
	}
	return false
}
