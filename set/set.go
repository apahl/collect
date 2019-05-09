// Package set implements generic set functionality.
// To be used with https://github.com/ncw/gotemplate
package set

import (
	"fmt"
	"strings"
)

// T is the generic type
// template type Set(T)
type T int

// Set is an implementation of a set.
// It uses a map in the background.
type Set map[T]empty

// The empty value for the map
type empty struct{}

func (s Set) String() string {
	var b strings.Builder
	for k := range s {
		fmt.Fprintf(&b, "%v ", k)
	}
	result := fmt.Sprintf("Set{ %s}", b.String())
	return result
}

// Len returns the length (the number of members)
// of a set.
func (s Set) Len() int {
	return len(s)
}

// Add adds an element to the set.
// If the element is already present, nothing happens.
func (s Set) Add(el T) {
	s[el] = empty{}
}

// Contains tests whether an element is present in a set.
func (s Set) Contains(el T) bool {
	_, result := s[el]
	return result
}

// IsEq tests for equality of two sets
func (s Set) IsEq(other Set) bool {
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

// New creates a new empty set
func New() Set {
	result := make(Set)
	return result
}

// FromSlice creates a set from a slice
func FromSlice(s []T) Set {
	result := make(Set)
	for _, k := range s {
		result.Add(k)
	}
	return result
}

// ToSlice converts a set to a slice
func ToSlice(s Set) []T {
	result := []T{}
	for k := range s {
		result = append(result, k)
	}
	return result
}
