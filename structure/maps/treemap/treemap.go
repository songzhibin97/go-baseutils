// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/maps"
	"github.com/songzhibin97/go-baseutils/structure/trees/redblacktree"
)

// Assert Map implementation
var _ maps.Map[int, any] = (*Map[int, any])(nil)

// Map holds the elements in a red-black tree
type Map[K, V any] struct {
	tree  *redblacktree.Tree[K, V]
	zeroK K
	zeroV V
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith[K, V any](comparator bcomparator.Comparator[K]) *Map[K, V] {
	return &Map[K, V]{tree: redblacktree.NewWith[K, V](comparator)}
}

// NewWithIntComparator instantiates a tree map with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator[V any]() *Map[int, V] {
	return &Map[int, V]{tree: redblacktree.NewWithIntComparator[V]()}
}

// NewWithStringComparator instantiates a tree map with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[V any]() *Map[string, V] {
	return &Map[string, V]{tree: redblacktree.NewWithStringComparator[V]()}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	m.tree.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.tree.Empty()
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.tree.Size()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[K, V]) Min() (key K, value V) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value
	}
	return m.zeroK, m.zeroV
}

// Max returns the maximum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[K, V]) Max() (key K, value V) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value
	}
	return m.zeroK, m.zeroV
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the map is empty, or because
// all keys in the map are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Floor(key K) (foundkey K, foundvalue V) {
	node, found := m.tree.Floor(key)
	if found {
		return node.Key, node.Value
	}
	return m.zeroK, m.zeroV
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the map is empty, or because
// all keys in the map are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Ceiling(key K) (foundkey K, foundvalue V) {
	node, found := m.tree.Ceiling(key)
	if found {
		return node.Key, node.Value
	}
	return m.zeroK, m.zeroV
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	bf := strings.Builder{}
	bf.WriteString("TreeMap\nmap[")
	it := m.Iterator()
	for it.Next() {
		bf.WriteString(fmt.Sprintf("(%v:%v) ", it.Key(), it.Value()))
	}
	bf.WriteString("]")
	return bf.String()
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error {
	return m.tree.UnmarshalJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	return m.tree.MarshalJSON()
}
