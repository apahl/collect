//go:generate genny -in=$GOFILE -out=maps.go gen "KeyType=int,string ValueType=bool,int,string"

// Package maps implements helper functions for map
package maps

import (
	"github.com/cheekybits/genny/generic"
)

// KeyType is the type for map keys.
type KeyType generic.Type

// ValueType is the type for map values.
type ValueType generic.Type

// KeyTypeValueTypeKeys returns a slice of KeyType map keys
func KeyTypeValueTypeKeys(m map[KeyType]ValueType) []KeyType {
	result := make([]KeyType, len(m))
	i := 0
	for k := range m {
		result[i] = k
		i++
	}
	return result
}

// KeyTypeValueTypeIsEq returns true when two maps have the same keys and values
// (using value semantics).
func KeyTypeValueTypeIsEq(m1, m2 map[KeyType]ValueType) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok == true {
			if v2 != v1 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
