package sets_test

import (
	"fmt"
	"testing"

	"github.com/apahl/collect/sets"
)

func TestIsEq(t *testing.T) {
	s1 := sets.IntSet{}
	s1.Add(1)
	s1.Add(1)
	if s1.Len() != 1 {
		t.Error("Length should be 1.")
	}
	s2 := sets.IntSetFromSlice([]int{1})
	if !s1.IsEq(s2) {
		t.Error("Sets should be equal.")
	}
	s3 := sets.NewIntSet()
	if s1.IsEq(s3) {
		t.Error("Sets should not be equal.")
	}
	s4 := sets.NewIntSet()
	s4.Add(3)
	if s1.IsEq(s4) {
		t.Error("Sets should not be equal.")
	}
}

func ExampleIntSet_IsEq() {
	s1 := sets.IntSet{}
	s1.Add(1)
	s1.Add(2)
	s2 := sets.NewIntSet()
	s2.Add(2)
	s2.Add(1)
	fmt.Println(s1.IsEq(s2)) // => true
}
