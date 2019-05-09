package maps_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apahl/collect/maps"
)

func TestMap_Keys(t *testing.T) {
	m := maps.Map{3: 0, 5: 0, 7: 0}
	assert.Equal(t, m.Keys(), []maps.T{3, 5, 7})
}

func TestIsEq(t *testing.T) {
	m1 := maps.Map{3: 0, 5: 0, 7: 0}
	m2 := maps.Map{3: 0, 5: 0, 7: 0}
	assert.True(t, maps.IsEq(m1, m2))
	m3 := maps.Map{3: 0, 5: 1, 7: 0}
	assert.False(t, maps.IsEq(m1, m3))
	m4 := maps.Map{3: 0, 5: 0}
	assert.False(t, maps.IsEq(m1, m4))
	m5 := maps.Map{3: 0, 5: 0, 7: 2}
	assert.False(t, maps.IsEq(m1, m5))
}

func ExampleMap_Keys() {
	m := maps.Map{3: 0, 5: 0, 7: 0}
	fmt.Println(m.Keys()) // => [3 5 7]
}
