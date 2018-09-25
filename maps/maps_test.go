package maps_test

import (
	"fmt"
	"testing"

	"github.com/apahl/collect/maps"
	"github.com/apahl/collect/slices"
)

func TestKeys(t *testing.T) {
	if v, r := map[int]bool{3: true, 5: true, 7: true}, []int{3, 5, 7}; !slices.IntIsEq(maps.IntBoolKeys(v), r) {
		t.Error("Keys() should be [3 5 7].")
	}
}

func TestStringIntIsEq(t *testing.T) {
	m1 := map[string]int{"eins": 1, "zwei": 2, "drei": 3}
	m2 := map[string]int{"eins": 1, "drei": 3, "zwei": 2}
	if !maps.StringIntIsEq(m1, m2) {
		t.Error("Maps m1 and m2 should be equal.")
	}
	m3 := map[string]int{"eins": 1, "zwei": 2, "drei": 0}
	if maps.StringIntIsEq(m1, m3) {
		t.Error("Maps m1 and m3 should NOT be equal.")
	}
	m4 := map[string]int{"eins": 1, "zwei": 2}
	if maps.StringIntIsEq(m1, m4) {
		t.Error("Maps m1 and m4 should NOT be equal.")
	}
	m5 := map[string]int{"eins": 1, "zwei": 2, "vier": 4}
	if maps.StringIntIsEq(m1, m5) {
		t.Error("Maps m1 and m5 should NOT be equal.")
	}

}

func ExampleIntBoolKeys() {
	m := map[int]bool{3: true, 5: true, 7: true}
	fmt.Println(maps.IntBoolKeys(m)) // => [3 5 7]
}
