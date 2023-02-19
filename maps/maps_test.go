package maps_test

import (
	"sort"
	"testing"

	"github.com/apahl/collect/maps"
)

type Employee struct {
	id   int
	name string
	age  int
}

func (m Employee) Hash() int {
	return m.id
}

func TestIntHashMap(t *testing.T) {
	m := maps.NewIntHashMap[Employee, int]()
	m.Add(Employee{id: 1, name: "Alice", age: 20}, 1000)
	m.Add(Employee{id: 2, name: "Bob", age: 21}, 2000)
	m.Add(Employee{id: 3, name: "Charlie", age: 22}, 3000)
	m.Add(Employee{id: 2, name: "Bob", age: 21}, 4000) // duplicate, value will be overwritten
	if m.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", m.Len())
	}
	if val, ok := m.Get(Employee{id: 2, name: "Bob", age: 21}); !ok || val != 4000 {
		t.Errorf("Expected Bob to have a salary of 4000, got %d.", val)
	}
	items := m.Items()
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}
	m.Remove(Employee{id: 2, name: "Bob", age: 21}) // Bob is fired.
	if m.Len() != 2 {
		t.Errorf("Expected 2 items, got %d", m.Len())
	}
	m.Add(Employee{id: 2, name: "Bob", age: 21}, 5000) // Bob gets re-hired with a better salry.
	if val, ok := m.Get(Employee{id: 2, name: "Bob", age: 21}); !ok || val != 5000 {
		t.Errorf("Expected Bob to have a salary of 5000, got %d.", val)
	}

	keys := m.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].name < keys[j].name
	})
	if keys[0].name != "Alice" || keys[1].name != "Bob" || keys[2].name != "Charlie" {
		t.Errorf("Expected Employees in order [{1 Alice 20} {2 Bob 21} {3 Charlie 22}], got %v", keys)
	}

	values := m.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	if values[0] != 1000 || values[1] != 3000 || values[2] != 5000 {
		t.Errorf("Expected Salaries in order [1000 3000 5000], got %v", values)
	}

}

// ---------------------------------------------------------------------------

type Person struct {
	name string
	age  int
}

func (m Person) Hash() string {
	return m.name
}

func TestStringHashMap(t *testing.T) {
	m := maps.NewStringHashMap[Person, int]()
	m.Add(Person{name: "Alice", age: 20}, 1000)
	m.Add(Person{name: "Bob", age: 21}, 2000)
	m.Add(Person{name: "Charlie", age: 22}, 3000)
	m.Add(Person{name: "Bob", age: 21}, 4000) // duplicate, value will be overwritten
	if m.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", m.Len())
	}
	if val, ok := m.Get(Person{name: "Bob", age: 21}); !ok || val != 4000 {
		t.Errorf("Expected Bob to have a salary of 4000, got %d.", val)
	}
	items := m.Items()
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}
	m.Remove(Person{name: "Bob", age: 21}) // Bob is fired.
	if m.Len() != 2 {
		t.Errorf("Expected 2 items, got %d", m.Len())
	}
	m.Add(Person{name: "Bob", age: 21}, 5000) // Bob gets re-hired with a better salry.
	if val, ok := m.Get(Person{name: "Bob", age: 21}); !ok || val != 5000 {
		t.Errorf("Expected Bob to have a salary of 5000, got %d.", val)
	}

	keys := m.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].name < keys[j].name
	})
	if keys[0].name != "Alice" || keys[1].name != "Bob" || keys[2].name != "Charlie" {
		t.Errorf("Expected Persons in order [{Alice 20} {Bob 21} {Charlie 22}], got %v", keys)
	}

	values := m.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	if values[0] != 1000 || values[1] != 3000 || values[2] != 5000 {
		t.Errorf("Expected Salaries in order [1000 3000 5000], got %v", values)
	}

}
