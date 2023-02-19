package sets_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/apahl/collect/sets"
)

func TestSimpleSet(t *testing.T) {
	s := sets.NewSimpleSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(2) // duplicate, should not added again
	if s.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", s.Len())
	}
	if !s.Contains(1) {
		t.Error("Expected 1 to be in the set")
	}
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected 3 items, got %d", len(slice))
	}
	sort.Ints(slice)
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Errorf("Expected [1, 2, 3], got %v", slice)
	}

	s2 := sets.NewSimpleSetFromSlice([]int{2, 3, 4, 5})
	slice = s.Union(s2).ToSlice()
	if len(slice) != 5 {
		t.Errorf("Expected 5 items, got %d", len(slice))
	}
	sort.Ints(slice)
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 || slice[3] != 4 || slice[4] != 5 {
		t.Errorf("Expected [1, 2, 3, 4, 5], got %v", slice)
	}

	slice = s.Intersect(s2).ToSlice()
	if len(slice) != 2 {
		t.Errorf("Expected 2 items, got %d", len(slice))
	}
	sort.Ints(slice)
	if slice[0] != 2 || slice[1] != 3 {
		t.Errorf("Expected [2, 3], got %v", slice)
	}

	slice = s.Difference(s2).ToSlice()
	if len(slice) != 1 {
		t.Errorf("Expected 1 item, got %d", len(slice))
	}
	if slice[0] != 1 {
		t.Errorf("Expected [1], got %v", slice)
	}
}

// ---------------------------------------------------------------------------

type Employee struct {
	id   int
	name string
	age  int
}

func (m Employee) Hash() int {
	return m.id
}

func TestIntHashSet(t *testing.T) {
	s := sets.NewIntHashSet[Employee]()
	s.Add(Employee{id: 1, name: "Alice", age: 20})
	s.Add(Employee{id: 2, name: "Bob", age: 21})
	s.Add(Employee{id: 3, name: "Charlie", age: 22})
	s.Add(Employee{id: 2, name: "Bob", age: 21}) // duplicate, should not added again
	if s.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", s.Len())
	}
	if !s.Contains(Employee{id: 1, name: "Alice", age: 20}) {
		t.Error("Expected Alice to be in the set")
	}
	if !s.Contains(Employee{id: 2, name: "Bob", age: 21}) {
		t.Error("Expected Bob to be in the set")
	}
	s.Remove(Employee{id: 2, name: "Bob", age: 21})
	if s.Len() != 2 {
		t.Errorf("Expected 2 items, got %d. Bob was not removed.", s.Len())
	}

	s = sets.NewIntHashSetFromSlice(
		[]Employee{
			{id: 1, name: "Alice", age: 20},
			{id: 2, name: "Bob", age: 21},
			{id: 3, name: "Charlie", age: 22},
			{id: 2, name: "Bob", age: 21}, // should not be added again
		},
	)
	if s.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", s.Len())
	}
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected 3 items, got %d", len(slice))
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].id < slice[j].id
	})
	if slice[0].id != 1 || slice[1].id != 2 || slice[2].id != 3 {
		t.Errorf("Expected Employees in order [{1 Alice 20} {2 Bob 21} {3 Charlie 22}], got %v", slice)
	}

	s2 := sets.NewIntHashSetFromSlice(
		[]Employee{
			{id: 2, name: "Bob", age: 21},
			{id: 3, name: "Charlie", age: 22},
			{id: 4, name: "David", age: 23},
			{id: 5, name: "Eve", age: 24},
		},
	)

	slice = s.Union(s2).ToSlice()
	if len(slice) != 5 {
		t.Errorf("Expected 5 items, got %d", len(slice))
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].id < slice[j].id
	})
	if slice[0].id != 1 || slice[1].id != 2 || slice[2].id != 3 || slice[3].id != 4 || slice[4].id != 5 {
		t.Errorf("Expected Employees in order [{1 Alice 20} {2 Bob 21} {3 Charlie 22} {4 David 23} {5 Eve 24}], got %v", slice)
	}

	slice = s.Intersect(s2).ToSlice()
	if len(slice) != 2 {
		t.Errorf("Expected 2 items, got %d", len(slice))
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].id < slice[j].id
	})
	if slice[0].id != 2 || slice[1].id != 3 {
		t.Errorf("Expected Employees in order [{2 Bob 21} {3 Charlie 22}], got %v", slice)
	}

	slice = s.Difference(s2).ToSlice()
	if len(slice) != 1 {
		t.Errorf("Expected 1 item, got %d", len(slice))
	}
	if slice[0].id != 1 {
		t.Errorf("Expected [{1 Alice 20}], got %v", slice)
	}
}

// ---------------------------------------------------------------------------

type Person struct {
	name string
	age  int
}

func (m Person) Hash() string {
	return fmt.Sprintf("%s (%d)", m.name, m.age)
}

func TestStringHashSet(t *testing.T) {
	s := sets.NewStringHashSet[Person]()
	s.Add(Person{name: "Alice", age: 20})
	s.Add(Person{name: "Bob", age: 21})
	s.Add(Person{name: "Charlie", age: 22})
	s.Add(Person{name: "Bob", age: 21}) // duplicate, should not added again
	if s.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", s.Len())
	}
	if !s.Contains(Person{name: "Alice", age: 20}) {
		t.Error("Expected Alice to be in the set")
	}
	if !s.Contains(Person{name: "Bob", age: 21}) {
		t.Error("Expected Bob to be in the set")
	}
	s.Remove(Person{name: "Bob", age: 21})
	if s.Len() != 2 {
		t.Errorf("Expected 2 items, got %d. Bob was not removed.", s.Len())
	}
	// Try to remove a value that is not in the set
	// This should not cause a panic
	s.Remove(Person{name: "Bob", age: 21})

	s = sets.NewStringHashSetFromSlice(
		[]Person{
			{name: "Alice", age: 20},
			{name: "Bob", age: 21},
			{name: "Charlie", age: 22},
			{name: "Bob", age: 21}, // should not be added again
		},
	)
	if s.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", s.Len())
	}
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected 3 items, got %d", len(slice))
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].name < slice[j].name
	})
	if slice[0].name != "Alice" || slice[1].name != "Bob" || slice[2].name != "Charlie" {
		t.Errorf("Expected Persons in order [{Alice 20} {Bob 21} {Charlie 22}], got %v", slice)
	}

	s2 := sets.NewStringHashSetFromSlice(
		[]Person{
			{name: "Bob", age: 21},
			{name: "Charlie", age: 22},
			{name: "David", age: 23},
			{name: "Eve", age: 24},
		},
	)
	slice = s.Union(s2).ToSlice()
	if len(slice) != 5 {
		t.Errorf("Expected 5 items, got %d", len(slice))
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].name < slice[j].name
	})
	if slice[0].name != "Alice" || slice[1].name != "Bob" || slice[2].name != "Charlie" || slice[3].name != "David" || slice[4].name != "Eve" {
		t.Errorf("Expected Persons in order [{Alice 20} {Bob 21} {Charlie 22} {David 23} {Eve 24}], got %v", slice)
	}

	slice = s.Intersect(s2).ToSlice()
	if len(slice) != 2 {
		t.Errorf("Expected 2 items, got %d", len(slice))
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].name < slice[j].name
	})
	if slice[0].name != "Bob" || slice[1].name != "Charlie" {
		t.Errorf("Expected Persons in order [{Bob 21} {Charlie 22}], got %v", slice)
	}

	slice = s.Difference(s2).ToSlice()
	if len(slice) != 1 {
		t.Errorf("Expected 1 item, got %d", len(slice))
	}
	if slice[0].name != "Alice" {
		t.Errorf("Expected [{Alice 20}], got %v", slice)
	}
}
