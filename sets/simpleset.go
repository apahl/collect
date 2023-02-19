// Package sets provides different set implementations.
// `SimpleSet` is a simple set implementation,
// using the built-in `comparable` type.
// `IntHashSet` is available for types that implement the IntHashable interface,
// having a Hash() method that returns an int.
// `StringHashSet` is available for types that implement the StringHashable interface,
// having a Hash() method that returns a string.
package sets

// go test ./...

// SimpleSet is a simple set implementation,
// It can be used with generic `comparable` types.
// Internally it uses a map.
type SimpleSet[T comparable] map[T]bool

// NewSimpleSet creates a new empty SimpleSet.
func NewSimpleSet[T comparable]() SimpleSet[T] {
	result := make(map[T]bool)
	return result
}

// NewSimpleSetFromSlice creates a new SimpleSet from a slice.
func NewSimpleSetFromSlice[T comparable](slice []T) SimpleSet[T] {
	result := make(map[T]bool)
	for _, v := range slice {
		result[v] = true
	}
	return result
}

// Add adds a value to the set.
func (s SimpleSet[T]) Add(v T) {
	s[v] = true
}

// Remove removes a value from the set.
func (s SimpleSet[T]) Remove(v T) {
	delete(s, v)
}

// Contains returns true if the value is in the set.
func (s SimpleSet[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Len returns the number of elements in the set.
func (s SimpleSet[T]) Len() int {
	return len(s)
}

// ToSlice returns a slice containing all the elements in the set.
func (s SimpleSet[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}

// Union returns a new set containing the union of the two sets.
func (s SimpleSet[T]) Union(other SimpleSet[T]) SimpleSet[T] {
	result := NewSimpleSet[T]()
	for k := range s {
		result.Add(k)
	}
	for k := range other {
		result.Add(k)
	}
	return result
}

// Intersect returns a new set containing the intersection of the two sets.
func (s SimpleSet[T]) Intersect(other SimpleSet[T]) SimpleSet[T] {
	result := NewSimpleSet[T]()
	for k := range s {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// Difference returns a new set containing the difference of the two sets.
func (s SimpleSet[T]) Difference(other SimpleSet[T]) SimpleSet[T] {
	result := NewSimpleSet[T]()
	for k := range s {
		if !other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}
