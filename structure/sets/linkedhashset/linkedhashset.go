// Package linkedhashset is a set that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Note that insertion-order is not affected if an element is re-inserted into the set.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package linkedhashset

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/structure/lists/doublylinkedlist"
	"github.com/songzhibin97/go-baseutils/structure/sets"
)

// Assert Set implementation
var _ sets.Set[int] = (*Set[int])(nil)

// Set holds elements in go's native map
type Set[E comparable] struct {
	table    map[E]struct{}
	ordering *doublylinkedlist.List[E]
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[E comparable](values ...E) *Set[E] {
	set := &Set[E]{
		table:    make(map[E]struct{}),
		ordering: doublylinkedlist.New[E](),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
// Note that insertion-order is not affected if an element is re-inserted into the set.
func (set *Set[E]) Add(items ...E) {
	for _, item := range items {
		if _, contains := set.table[item]; !contains {
			set.table[item] = itemExists
			set.ordering.Append(item)
		}
	}
}

// Remove removes the items (one or more) from the set.
// Slow operation, worst-case O(n^2).
func (set *Set[E]) Remove(items ...E) {
	for _, item := range items {
		if _, contains := set.table[item]; contains {
			delete(set.table, item)
			index := set.ordering.IndexOf(item)
			set.ordering.Remove(index)
		}
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[E]) Contains(items ...E) bool {
	for _, item := range items {
		if _, contains := set.table[item]; !contains {
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
	return set.ordering.Size()
}

// Clear clears all values in the set.
func (set *Set[E]) Clear() {
	set.table = make(map[E]struct{})
	set.ordering.Clear()
}

// Values returns all items in the set.
func (set *Set[E]) Values() []E {
	values := make([]E, set.Size())
	it := set.Iterator()
	for it.Next() {
		values[it.Index()] = it.Value()
	}
	return values
}

// String returns a string representation of container
func (set *Set[E]) String() string {
	b := strings.Builder{}
	b.WriteString("LinkedHashSet\n")
	it := set.Iterator()
	for it.Next() {
		b.WriteString(fmt.Sprintf("(key:%v) ", it.Value()))
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
		for item := range set.table {
			if _, contains := another.table[item]; contains {
				result.Add(item)
			}
		}
	} else {
		for item := range another.table {
			if _, contains := set.table[item]; contains {
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

	for item := range set.table {
		result.Add(item)
	}
	for item := range another.table {
		result.Add(item)
	}

	return result
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[E]) Difference(another *Set[E]) *Set[E] {
	result := New[E]()

	for item := range set.table {
		if _, contains := another.table[item]; !contains {
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
