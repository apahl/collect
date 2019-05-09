// Package slice implements generic slice utilities.
// To be used with https://github.com/ncw/gotemplate
package slice

// T is the generic type
// template type Slice(T)
type T int

// IsEq checks for equality of two int slices a, b
func IsEq(a, b []T) bool {
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

// Sum returns the sum of all elements of the slice.
// Is defined for numbers and strings.
func Sum(a []T) T {
	var result T
	for _, i := range a {
		result += i
	}
	return result
}
