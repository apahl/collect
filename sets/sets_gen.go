//go:generate genny -in=$GOFILE -out=sets.go gen "KeyType=int8,int16,int32,int"

package sets

import (
	"fmt"
	"strings"

	"github.com/cheekybits/genny/generic"
)

// KeyType is the set type.
type KeyType generic.Type

// KeyTypeSet is an implementation of a set holding integers.
// It uses a map in the background.
type KeyTypeSet map[KeyType]bool

func (s KeyTypeSet) String() string {
	var b strings.Builder
	for k := range s {
		fmt.Fprintf(&b, "%d ", k)
	}
	result := fmt.Sprintf("KeyTypeSet{ %s}", b.String())
	return result
}

// Len returns the length (the number of members)
// of a set.
func (s KeyTypeSet) Len() int {
	return len(s)
}

// Add adds an element to the set.
// If the element is already present, nothing happens.
func (s KeyTypeSet) Add(el KeyType) {
	s[el] = true
}

// Contains tests whether an element is present in a set.
func (s KeyTypeSet) Contains(el KeyType) bool {
	_, result := s[el]
	return result
}

// IsEq tests for equality of two sets
func (s KeyTypeSet) IsEq(other KeyTypeSet) bool {
	slen := s.Len()
	olen := other.Len()
	if slen != olen {
		return false
	}
	if slen == 0 {
		return true
	}
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// NewKeyTypeSet creates a new empty set
func NewKeyTypeSet() KeyTypeSet {
	result := make(KeyTypeSet)
	return result
}

// KeyTypeSetFromSlice creates a set from a slice
func KeyTypeSetFromSlice(s []KeyType) KeyTypeSet {
	result := make(KeyTypeSet)
	for _, k := range s {
		result.Add(k)
	}
	return result
}
