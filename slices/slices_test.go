package slices_test

import (
	"testing"

	"github.com/apahl/collect/slices"
)

func TestSlices(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	if !slices.AreEqual(a, b) {
		t.Error("Expected slices to be equal")
	}
	b[0] = 2
	if slices.AreEqual(a, b) {
		t.Error("Expected slices to be NOT equal (have different values))")
	}

	c := []int{3, 1, 2, 4}
	if slices.AreEqual(a, c) {
		t.Error("Expected slices to be NOT equal (have different orders)")
	}

	d := []int{1, 2, 3, 4, 5}
	if slices.AreEqual(a, d) {
		t.Error("Expected slices to be NOT equal (have different lengths)")
	}
}
