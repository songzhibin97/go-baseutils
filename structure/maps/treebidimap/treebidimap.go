// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/base/anytostring"
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/maps"
	"github.com/songzhibin97/go-baseutils/structure/trees/redblacktree"
)

// Assert Map implementation
var _ maps.BidiMap[int, int] = (*Map[int, int])(nil)

// Map holds the elements in two red-black trees.
type Map[K, V any] struct {
	forwardMap      redblacktree.Tree[K, V]
	inverseMap      redblacktree.Tree[V, K]
	keyComparator   bcomparator.Comparator[K]
	valueComparator bcomparator.Comparator[V]
}

// NewWith instantiates a bidirectional map.
func NewWith[K, V any](keyComparator bcomparator.Comparator[K], valueComparator bcomparator.Comparator[V]) *Map[K, V] {
	return &Map[K, V]{
		forwardMap:      *redblacktree.NewWith[K, V](keyComparator),
		inverseMap:      *redblacktree.NewWith[V, K](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewWithIntComparators instantiates a bidirectional map with the IntComparator for key and value, i.e. keys and values are of type int.
func NewWithIntComparators() *Map[int, int] {
	return NewWith(bcomparator.IntComparator(), bcomparator.IntComparator())
}

// NewWithStringComparators instantiates a bidirectional map with the StringComparator for key and value, i.e. keys and values are of type string.
func NewWithStringComparators() *Map[string, string] {
	return NewWith(bcomparator.StringComparator(), bcomparator.StringComparator())
}

// Put inserts element into the map.
func (m *Map[K, V]) Put(key K, value V) {
	if d, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(d)
	}
	if d, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(d)
	}
	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d, true
	}
	var zero V
	return zero, false
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, V]) GetKey(value V) (key K, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d, true
	}
	var zero K
	return zero, false
}

// Remove removes the element from the map by key.
func (m *Map[K, V]) Remove(key K) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (ordered).
func (m *Map[K, V]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
func (m *Map[K, V]) Values() []V {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	bf := strings.Builder{}
	bf.WriteString("TreeBidiMap\nmap[")
	it := m.Iterator()
	for it.Next() {
		bf.WriteString(fmt.Sprintf("(%v:%v) ", it.Key(), it.Value()))
	}
	bf.WriteString("]")
	return bf.String()
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(bytes, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			var nk K
			err = m.keyComparator.Unmarshal([]byte(key), &nk)
			if err != nil {
				return err
			}
			var nv V
			err = m.valueComparator.Unmarshal([]byte(anytostring.AnyToString(value)), &nv)
			if err != nil {
				return err
			}
			m.Put(nk, nv)
		}
	}
	return err
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	elements := make(map[string]string)
	it := m.Iterator()
	for it.Next() {
		k, err := m.keyComparator.Marshal(it.Key())
		if err != nil {
			return nil, err
		}
		v, err := m.valueComparator.Marshal(it.Value())
		if err != nil {
			return nil, err
		}
		elements[string(k)] = string(v)
	}
	return json.Marshal(&elements)
}
