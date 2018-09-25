package slices_test

import (
	"fmt"
	"testing"

	"github.com/apahl/collect/slices"
)

func TestIntIsEq(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 5}
	if !slices.IntIsEq(a, b) {
		t.Error("a and b should be equal.")
	}
	if slices.IntIsEq(a, c) {
		t.Error("a and c should NOT be equal.")
	}
	if !slices.IntIsEq([]int{1, 0, 4}, []int{1, 0, 4}) {
		t.Error("test should  be equal.")
	}
}

func TestSum(t *testing.T) {
	if i, o := []int{1, 2, 3, 6}, 12; slices.Sum(i) != o {
		t.Errorf("Sum(%v) should be %v", i, o)
	}
}

func ExampleIntIsEq() {
	s1 := []int{1, 2, 6}
	s2 := []int{1, 2, 6}
	fmt.Println(slices.IntIsEq(s1, s2)) // => true
}

func ExampleSum() {
	s := []int{1, 2, 3, 6}
	fmt.Println(slices.Sum(s)) // => 12
}
