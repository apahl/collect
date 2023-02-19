package sets

// IntHashable defines the interface for a type that can be hashed to an int.
type IntHashable[T any] interface {
	Hash() int
}

// IntHashSet is a set of IntHashable values.
type IntHashSet[T IntHashable[T]] map[int]T

// NewIntHashSet creates a new empty IntHashSet.
func NewIntHashSet[T IntHashable[T]]() IntHashSet[T] {
	result := make(map[int]T)
	return result
}

// NewIntHashSetFromSlice creates a new IntHashSet from a slice.
func NewIntHashSetFromSlice[T IntHashable[T]](slice []T) IntHashSet[T] {
	result := make(map[int]T)
	for _, v := range slice {
		result[v.Hash()] = v
	}
	return result
}

// Add adds a value to the set.
// If the value is already in the set, it is not added again.
func (i IntHashSet[T]) Add(v T) {
	i[v.Hash()] = v
}

// Remove removes a value from the set.
// If the value is not in the set, nothing happens.
func (i IntHashSet[T]) Remove(v T) {
	delete(i, v.Hash())
}

// Contains returns true if the value is in the set.
func (i IntHashSet[T]) Contains(v T) bool {
	_, ok := i[v.Hash()]
	return ok
}

// Len returns the number of elements in the set.
func (i IntHashSet[T]) Len() int {
	return len(i)
}

// ToSlice returns a slice containing all the elements in the set.
func (i IntHashSet[T]) ToSlice() []T {
	result := make([]T, 0, len(i))
	for _, v := range i {
		result = append(result, v)
	}
	return result
}

// Union returns a new set containing the union of the two sets.
func (i IntHashSet[T]) Union(other IntHashSet[T]) IntHashSet[T] {
	result := NewIntHashSet[T]()
	for _, v := range i {
		result.Add(v)
	}
	for _, v := range other {
		result.Add(v)
	}
	return result
}

// Intersect returns a new set containing the intersection of the two sets.
func (i IntHashSet[T]) Intersect(other IntHashSet[T]) IntHashSet[T] {
	result := NewIntHashSet[T]()
	for _, v := range i {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns a new set containing the difference of the two sets.
func (i IntHashSet[T]) Difference(other IntHashSet[T]) IntHashSet[T] {
	result := NewIntHashSet[T]()
	for _, v := range i {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// ---------------------------------------------------------------------------

// StringHashable defines the interface for a type that can be hashed to a string.
type StringHashable[T any] interface {
	Hash() string
}

// StringHashSet is a set of StringHashable values.
type StringHashSet[T StringHashable[T]] map[string]T

// NewStringHashSet creates a new empty StringHashSet.
func NewStringHashSet[T StringHashable[T]]() StringHashSet[T] {
	result := make(map[string]T)
	return result
}

// NewStringHashSetFromSlice creates a new StringHashSet from a slice.
func NewStringHashSetFromSlice[T StringHashable[T]](slice []T) StringHashSet[T] {
	result := make(map[string]T)
	for _, v := range slice {
		result[v.Hash()] = v
	}
	return result
}

// Add adds a value to the set.
func (s StringHashSet[T]) Add(v T) {
	s[v.Hash()] = v
}

// Remove removes a value from the set.
// If the value is not in the set, nothing happens.
func (s StringHashSet[T]) Remove(v T) {
	delete(s, v.Hash())
}

// Contains returns true if the value is in the set.
func (s StringHashSet[T]) Contains(v T) bool {
	_, ok := s[v.Hash()]
	return ok
}

// Len returns the number of elements in the set.
func (s StringHashSet[T]) Len() int {
	return len(s)
}

// ToSlice returns a slice containing all the elements in the set.
func (s StringHashSet[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		result = append(result, v)
	}
	return result
}

// Union returns a new set containing the union of the two sets.
func (s StringHashSet[T]) Union(other StringHashSet[T]) StringHashSet[T] {
	result := NewStringHashSet[T]()
	for _, v := range s {
		result.Add(v)
	}
	for _, v := range other {
		result.Add(v)
	}
	return result
}

// Intersect returns a new set containing the intersection of the two sets.
func (s StringHashSet[T]) Intersect(other StringHashSet[T]) StringHashSet[T] {
	result := NewStringHashSet[T]()
	for _, v := range s {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns a new set containing the difference of the two sets.
func (s StringHashSet[T]) Difference(other StringHashSet[T]) StringHashSet[T] {
	result := NewStringHashSet[T]()
	for _, v := range s {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}
