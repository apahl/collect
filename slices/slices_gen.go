//go:generate genny -in=$GOFILE -out=slices.go gen "GenericType=int,string"

// Package slices implements additional functionality for slices
// and standard types
package slices

import (
	"github.com/cheekybits/genny/generic"
)

// GenericType is the genny generic type
type GenericType generic.Type

// GenericTypeIsEq checks for equality of two int slices a, b
func GenericTypeIsEq(a, b []GenericType) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
