package maps

// IntHashable defines the interface for a type that can be hashed to an int.
type IntHashable[T any] interface {
	Hash() int
}

// IntHashMap is a dictionary type that maps from any type that implements IntHashable to any type.
// Internally, it uses a map[int]T and a map[int]V.
type IntHashMap[T IntHashable[T], V any] struct {
	hashToKey map[int]T
	hashToVal map[int]V
}

// NewIntHashMap creates a new empty IntHashMap.
func NewIntHashMap[T IntHashable[T], V any]() IntHashMap[T, V] {
	return IntHashMap[T, V]{
		hashToKey: make(map[int]T),
		hashToVal: make(map[int]V),
	}
}

// Add adds a key-value pair to the map.
// The key needs to implement IntHashable.
// If the key is already in the map, the value is overwritten.
func (i IntHashMap[T, V]) Add(key T, val V) {
	hash := key.Hash()
	i.hashToKey[hash] = key
	i.hashToVal[hash] = val
}

// Get returns the value associated with the key.
// If the key is not in the map, the second return value is false.
func (i IntHashMap[T, V]) Get(key T) (V, bool) {
	hash := key.Hash()
	val, ok := i.hashToVal[hash]
	return val, ok
}

// Remove removes a key-value pair from the map.
// If the key is not in the map, nothing happens.
func (i IntHashMap[T, V]) Remove(key T) {
	hash := key.Hash()
	delete(i.hashToKey, hash)
	delete(i.hashToVal, hash)
}

// Contains returns true if the key is in the map.
func (i IntHashMap[T, V]) Contains(key T) bool {
	hash := key.Hash()
	_, ok := i.hashToKey[hash]
	return ok
}

// Len returns the number of key-value pairs in the map.
func (i IntHashMap[T, V]) Len() int {
	// DEBUG:
	if len(i.hashToKey) != len(i.hashToVal) {
		panic("hashToKey and hashToVal have different lengths")
	}
	return len(i.hashToKey)
}

// Keys returns a slice of all the keys in the map.
func (i IntHashMap[T, V]) Keys() []T {
	result := make([]T, 0, len(i.hashToKey))
	for _, key := range i.hashToKey {
		result = append(result, key)
	}
	return result
}

// Values returns a slice of all the values in the map.
func (i IntHashMap[T, V]) Values() []V {
	result := make([]V, 0, len(i.hashToKey))
	for _, val := range i.hashToVal {
		result = append(result, val)
	}
	return result
}

// Items returns a slice of all the key-value pairs in the map.
func (i IntHashMap[T, V]) Items() []struct {
	Key T
	Val V
} {
	result := make([]struct {
		Key T
		Val V
	}, 0, len(i.hashToKey))
	for hash, key := range i.hashToKey {
		val := i.hashToVal[hash]
		result = append(result, struct {
			Key T
			Val V
		}{key, val})
	}
	return result
}

// ---------------------------------------------------------------------------

// StringHashable defines the interface for a type that can be hashed to an int.
type StringHashable[T any] interface {
	Hash() string
}

// StringHashMap is a dictionary type that maps from any type that implements StringHashable to any type.
// Internally, it uses a map[string]T and a map[string]V.
type StringHashMap[T StringHashable[T], V any] struct {
	hashToKey map[string]T
	hashToVal map[string]V
}

// NewStringHashMap creates a new empty StringHashMap.
func NewStringHashMap[T StringHashable[T], V any]() StringHashMap[T, V] {
	return StringHashMap[T, V]{
		hashToKey: make(map[string]T),
		hashToVal: make(map[string]V),
	}
}

// Add adds a key-value pair to the map.
// The key needs to implement StringHashable.
// If the key is already in the map, the value is overwritten.
func (i StringHashMap[T, V]) Add(key T, val V) {
	hash := key.Hash()
	i.hashToKey[hash] = key
	i.hashToVal[hash] = val
}

// Get returns the value associated with the key.
// If the key is not in the map, the second return value is false.
func (i StringHashMap[T, V]) Get(key T) (V, bool) {
	hash := key.Hash()
	val, ok := i.hashToVal[hash]
	return val, ok
}

// Remove removes a key-value pair from the map.
// If the key is not in the map, nothing happens.
func (i StringHashMap[T, V]) Remove(key T) {
	hash := key.Hash()
	delete(i.hashToKey, hash)
	delete(i.hashToVal, hash)
}

// Contains returns true if the key is in the map.
func (i StringHashMap[T, V]) Contains(key T) bool {
	hash := key.Hash()
	_, ok := i.hashToKey[hash]
	return ok
}

// Len returns the number of key-value pairs in the map.
func (i StringHashMap[T, V]) Len() int {
	// DEBUG:
	if len(i.hashToKey) != len(i.hashToVal) {
		panic("hashToKey and hashToVal have different lengths")
	}
	return len(i.hashToKey)
}

// Keys returns a slice of all the keys in the map.
func (i StringHashMap[T, V]) Keys() []T {
	result := make([]T, 0, len(i.hashToKey))
	for _, key := range i.hashToKey {
		result = append(result, key)
	}
	return result
}

// Values returns a slice of all the values in the map.
func (i StringHashMap[T, V]) Values() []V {
	result := make([]V, 0, len(i.hashToKey))
	for _, val := range i.hashToVal {
		result = append(result, val)
	}
	return result
}

// Items returns a slice of all the key-value pairs in the map.
func (i StringHashMap[T, V]) Items() []struct {
	Key T
	Val V
} {
	result := make([]struct {
		Key T
		Val V
	}, 0, len(i.hashToKey))
	for hash, key := range i.hashToKey {
		val := i.hashToVal[hash]
		result = append(result, struct {
			Key T
			Val V
		}{key, val})
	}
	return result
}
