package slice_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apahl/collect/slice"
)

func TestIsEq(t *testing.T) {
	a := []slice.T{1, 2, 3}
	b := []slice.T{1, 2, 3}
	c := []slice.T{1, 2, 5}
	assert.True(t, slice.IsEq(a, b))
	assert.False(t, slice.IsEq(a, c))
}
func TestSum(t *testing.T) {
	a := []slice.T{1, 2, 0, 3}
	c := []slice.T{1, 0, 2, 5}
	assert.Equal(t, slice.Sum(a), slice.T(6))
	assert.Equal(t, slice.Sum(c), slice.T(8))
}

func ExampleIsEq() {
	s1 := []slice.T{1, 2, 6}
	s2 := []slice.T{1, 2, 6}
	fmt.Println(slice.IsEq(s1, s2)) // => true
}

func ExampleSum() {
	s1 := []slice.T{1, 2, 6}
	fmt.Println(slice.Sum(s1)) // => 9
}
