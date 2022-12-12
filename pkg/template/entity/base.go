package entity

import (
	"io"
	"text/template"

	"github.com/marlin5555/dddgen/pkg/code-gen/entity"
	"github.com/marlin5555/dddgen/pkg/template/util"
)

// generate entity/base/{{group}}.go entity/base/common.go entity/base/util.go

// table column gen interface template
var tpl = `// Package {{.Base.Package}} Code generated, DO NOT EDIT.
{{- $eMap:=.FullEntities}}
{{- $needZero:=false}}
{{- range $k,$v := .EntityMap}}
{{- range $k1,$v1 := $v.Attrs}}
{{- if not $v1.Batch}}
{{- range $idx,$ref := $v1.Fer}}
{{- $refE := index $eMap $ref.EntityName}}
{{- $refC := index $refE.Attrs $ref.AttributeID}}
{{- $key := printf "%s%s" (firstUpper $refC.Ref.Role) $refC.Ref.AttributeID}}
{{- if not (eq $key (printf "%s%s" $k "ID"))}}
{{- if not (eq (batch $key) (printf "%s%s" $k (batch $k1)))}}
{{- $needZero = true}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
package {{.Base.Package}}
{{- if $needZero}}
import(
    "{{.Module}}{{.Zero.Path}}"
)
{{- end}}
{{- range $k,$v := .EntityMap}}
// base interface for {{$k}}
type (
  {{- range $k1,$v1 := $v.Attrs -}}
    {{- $key := (printf "%s%s" $k $k1)}}
    {{- if $v1.Fuzzy -}}{{- $key = $k1 -}}{{- end -}}
    {{- if not (or $v1.Shared $v1.Ref.Type)}}
    // {{$key}} interface has Get{{$key}}
    {{$key}} interface {
       Get{{$key}}() {{if $v1.Batch}}[]{{end}}{{$v1.AttributeType}}
    }
    {{- else if and (not (eq $v1.Ref.Role $v1.Ref.EntityName)) (not $v1.Batch)}}
      {{- $interface := printf "%s%s" (firstUpper $v1.Ref.Role) $v1.Ref.AttributeID}}
    // {{$interface}} interface has Get{{$interface}}
    {{$interface}} interface {
       Get{{$interface}}() {{$v1.AttributeType}}
    }
      {{- $interface = printf "%s%s" (firstUpper $v1.Ref.Role) (batch $v1.Ref.AttributeID)}}
    // {{$interface}} interface has Get{{$interface}}
    {{$interface}} interface {
       Get{{$interface}}() []{{$v1.AttributeType}}
    }
    {{- end}}
  {{- end}}
)
{{- end}}

{{- range $k,$v := .EntityMap}}
// To{{$k}}IDs convert {{$k}}ID s to IDs
func To{{$k}}IDs[T {{$k}}ID](arrays []T) []{{$v.IDType}}{
	return To{{firstUpper $v.IDType}}Set(arrays, func(i interface{}) {{$v.IDType}} { return i.(T).Get{{$k}}ID() })
}

  {{- range $k1,$v1 := $v.Attrs}}
    {{- if not $v1.Batch}}
      {{- range $idx,$ref := $v1.Fer}}
        {{- $refE := index $eMap $ref.EntityName}}
        {{- $refC := index $refE.Attrs $ref.AttributeID}}
        {{- $key := printf "%s%s" (firstUpper $refC.Ref.Role) $refC.Ref.AttributeID}}
        {{- if not (eq $key (printf "%s%s" $k "ID"))}}
// To{{batch $key}} convert {{$key}} s to []{{$refE.IDType}}
func To{{batch $key}}[T {{$key}}](arrays []T)[]{{$refE.IDType}}{
    return To{{firstUpper $refE.IDType}}Set(arrays, func(i interface{}) {{$refE.IDType}}{ return i.(T).Get{{$key}}()})
}
          {{- if not (eq (batch $key) (printf "%s%s" $k (batch $k1)))}}
// {{batch $key}}To{{$k}}{{batch $k1}} convert {{batch $key}} to {{$k}}{{batch $k1}}
func {{batch $key}}To{{$k}}{{batch $k1}}(i {{batch $key}}){{$k}}{{batch $k1}}{
    return zero.{{$k}}{{batch $k1}}{Value: i.Get{{batch $key}}()}
}
          {{- end}}
        {{- end}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end}}
`

var interfaceTempl = template.Must(template.New("interface").
	Funcs(template.FuncMap{"firstLower": util.FirstLower,
		"firstUpper": util.FirstUpper,
		"batch":      util.Batch}).Parse(tpl))

func WriteInterface(writer io.Writer, domain *entity.Group) {
	_ = interfaceTempl.Execute(writer, domain)
}

// table common column gen interface template
var tpl1 = `// Package {{.Base.Package}} Code generated, DO NOT EDIT.
package {{.Base.Package}}
{{range .EntityImports}}
import "{{.}}"
{{end}}
{{range $k,$v := .EntityMap}}
{{- range $k1,$v1 := $v.Attrs}}
{{- $key := $k1}}
{{- if $v1.Shared }}
// {{$key}} interface has Get{{$key}}
type {{$key}} interface {
   Get{{$key}}() {{$v1.AttributeType}}
}
{{end}}
{{- end}}
{{- end}}

type (
	// RspCode interface return int32
	RspCode interface {
		GetCode() int32
	}

	// RspMsg interface return string
	RspMsg interface {
		GetMsg() string
	}

	// RspErr interface return error
	RspErr interface {
		GetError() error
		IsOK() bool
		NotOK() bool
	}
)
type (
	// Rtx interface GetRtx return string
	Rtx interface {
		GetRtx() string
	}

	// Operator combination interface: Rtx
	Operator interface {
		Rtx
	}
)

type (
	// Offset interface GetOffset return int
	Offset interface {
		GetOffset() int
	}

	// Limit interface GetLimit return int
	Limit interface {
		GetLimit() int
	}

	// Order interface OrderStr return string
	Order interface {
		OrderStr() string
	}

	// PageQuery combination interface: Offset Limit
	PageQuery interface {
		Offset
		Limit
	}

	// PageOrder combination interface: PageQuery Order
	PageOrder interface {
		PageQuery
		Order
	}

	// PageOrderOperator combination interface: PageOrder Operator
	PageOrderOperator interface {
		PageOrder
		Operator
	}
)

// Fields interface return field names []string
type Fields interface {
	GetFieldNames() []string
}

// Validator interface Validate when ok return nil, others return error
type Validator interface {
	Validate() error
}
`

var commonInterfaceTempl = template.Must(template.New("common-interface").Parse(tpl1))

func WriteCommonInterface(writer io.Writer, domain *entity.Group) {
	_ = commonInterfaceTempl.Execute(writer, domain)
}

var utiltpl = `// Package {{.Base.Package}} Code generated, DO NOT EDIT.
package {{.Base.Package}}

import "reflect"

// IsZero 判断 v 是否是零值
func IsZero(v interface{}) bool {
	return reflect.DeepEqual(reflect.Zero(reflect.TypeOf(v)).Interface(), v)
}
// ToKeys 获取 map[K]V 的 key set
func ToKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

type getStringKey func(interface{}) string

// ToStringKeyMap 将 arrays 转换成 map[string]T
func ToStringKeyMap[T any](arrays []T, f getStringKey) map[string]T {
	m := make(map[string]T)
	for _, b := range arrays {
		m[f(b)] = b
	}
	return m
}

type getUint32Key func(interface{}) uint32

// ToUint32KeyMap 将 arrays 转换成 map[uint32]T
func ToUint32KeyMap[T any](arrays []T, f getUint32Key) map[uint32]T {
	m := make(map[uint32]T)
	for _, b := range arrays {
		m[f(b)] = b
	}
	return m
}

// MergeUint32IDs 合并 idsList
func MergeUint32IDs(iDsList ...[]uint32) []uint32 {
	m := make(map[uint32]bool)
	var r []uint32
	for _, iDs := range iDsList {
		for _, id := range iDs {
			if _, ok := m[id]; !ok {
				r = append(r, id)
				m[id] = true
			}
		}
	}
	return r
}

func distinctUint32IDs(iDs []uint32) []uint32 {
	return MergeUint32IDs(iDs)
}

// ToUint32Set 将 arrays 转换成去重后的 []uint32
func ToUint32Set[T any](arrays []T, f getUint32Key) []uint32 {
	var iDs []uint32
	for _, a := range arrays {
		iDs = append(iDs, f(a))
	}
	return distinctUint32IDs(iDs)
}

// MergeStringIDs 合并 idsList
func MergeStringIDs(iDsList ...[]string) []string {
	m := make(map[string]bool)
	var r []string
	for _, iDs := range iDsList {
		for _, id := range iDs {
			if _, ok := m[id]; !ok {
				r = append(r, id)
				m[id] = true
			}
		}
	}
	return r
}

func distinctStringIDs(iDs []string) []string {
	return MergeStringIDs(iDs)
}

// ToStringSet 将 arrays 转换成去重后的 []string
func ToStringSet[T any](arrays []T, f getStringKey) []string {
	var iDs []string
	for _, a := range arrays {
		iDs = append(iDs, f(a))
	}
	return distinctStringIDs(iDs)
}

// ToPointers 对 array T 取指针成 array *T
func ToPointers[T any](arrays []T) []*T {
	var r []*T
	for idx := range arrays {
		r = append(r, &arrays[idx])
	}
	return r
}
`

var utilInterfaceTempl = template.Must(template.New("util-interface").Parse(utiltpl))

func WriteUtilInterface(writer io.Writer, domain *entity.Group) {
	_ = utilInterfaceTempl.Execute(writer, domain)
}
