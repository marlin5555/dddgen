package entity

import "github.com/marlin5555/dddgen/pkg/code-gen/util"

type IDFlag struct {
	AttrName string
	AttrType string
	Type     IDFlagType
}
type IDFlagType int

const (
	_ IDFlagType = iota
	AUTO
	GIVEN
	NAMED
)

var strToIDFlagType = map[string]IDFlagType{
	util.AUTO:  AUTO,
	util.GIVEN: GIVEN,
	util.NAMED: NAMED,
}

var idFlagTypeToStr = map[IDFlagType]string{
	AUTO:  util.AUTO,
	GIVEN: util.GIVEN,
	NAMED: util.NAMED,
}

func (t IDFlagType) String() string {
	if s, ok := idFlagTypeToStr[t]; ok {
		return s
	}
	return ""
}

func ToIDFlagType(s string) IDFlagType {
	if i, ok := strToIDFlagType[s]; ok {
		return i
	}
	return 0
}

func BuildIDFlag(t IDFlagType, attrName string, attrType string) IDFlag {
	return IDFlag{
		AttrName: attrName,
		AttrType: attrType,
		Type:     t,
	}
}

func Auto(t IDFlagType) bool {
	return t == AUTO
}
