// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/structure/sets"
)

// Assert Set implementation
var _ sets.Set[int] = (*Set[int])(nil)

// Set holds elements in go's native map
type Set[E comparable] struct {
	items map[E]struct{}
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[E comparable](values ...E) *Set[E] {
	set := &Set[E]{items: make(map[E]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[E]) Add(items ...E) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[E]) Remove(items ...E) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[E]) Contains(items ...E) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[E]) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[E]) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set[E]) Clear() {
	set.items = make(map[E]struct{})
}

// Values returns all items in the set.
func (set *Set[E]) Values() []E {
	values := make([]E, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Set[E]) String() string {
	b := strings.Builder{}
	b.WriteString("HashSet\n")
	for k := range set.items {
		b.WriteString(fmt.Sprintf("(key:%v) ", k))
	}
	return b.String()
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "another".
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set[E]) Intersection(another *Set[E]) *Set[E] {
	result := New[E]()

	// Iterate over smaller set (optimization)
	if set.Size() <= another.Size() {
		for item := range set.items {
			if _, contains := another.items[item]; contains {
				result.Add(item)
			}
		}
	} else {
		for item := range another.items {
			if _, contains := set.items[item]; contains {
				result.Add(item)
			}
		}
	}

	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[E]) Union(another *Set[E]) *Set[E] {
	result := New[E]()

	for item := range set.items {
		result.Add(item)
	}
	for item := range another.items {
		result.Add(item)
	}

	return result
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[E]) Difference(another *Set[E]) *Set[E] {
	result := New[E]()

	for item := range set.items {
		if _, contains := another.items[item]; !contains {
			result.Add(item)
		}
	}

	return result
}

// UnmarshalJSON @implements json.Unmarshaler
func (set *Set[E]) UnmarshalJSON(bytes []byte) error {
	elements := []E{}
	err := json.Unmarshal(bytes, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// MarshalJSON @implements json.Marshaler
func (set *Set[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}
