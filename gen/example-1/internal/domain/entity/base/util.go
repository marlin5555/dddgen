// Package base Code generated, DO NOT EDIT.
package base

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
