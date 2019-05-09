// Package set implements generic set functionality.
// To be used with https://github.com/ncw/gotemplate

// Package maps implements helper functions for map
package maps

// T is the generic key type
// template type Map(T, S)
type T int

// S is the generic value type
type S int

// Map is the generic map type
type Map map[T]S

// Keys returns a slice of int map keys
func (m Map) Keys() []T {
	result := make([]T, len(m))
	i := 0
	for k := range m {
		result[i] = k
		i++
	}
	return result
}

// IsEq returns true when two maps have the same keys and values
// (using value semantics).
func IsEq(m1, m2 Map) bool {
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

// SameKeys returns true when two maps have the same keys
// (using value semantics).
func SameKeys(m1, m2 Map) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1 := range m1 {
		if _, ok := m2[k1]; !ok {
			return false
		}
	}
	return true
}
