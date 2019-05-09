package set_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apahl/collect/set"
)

func TestIsEq(t *testing.T) {
	s1 := set.Set{}
	s1.Add(1)
	s1.Add(1)
	assert.Equal(t, s1.Len(), 1)
	s2 := set.FromSlice([]set.T{1})
	assert.True(t, s1.IsEq(s2))
	s3 := set.New()
	assert.False(t, s1.IsEq(s3))
	s4 := set.New()
	s4.Add(3)
	assert.False(t, s1.IsEq(s4))
}

func ExampleSet_IsEq() {
	s1 := set.Set{}
	s1.Add(1)
	s1.Add(2)
	s2 := set.New()
	s2.Add(2)
	s2.Add(1)
	fmt.Println(s1.IsEq(s2)) // => true
}
