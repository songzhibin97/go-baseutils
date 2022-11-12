// Package linkedhashmap is a map that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package linkedhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/lists/doublylinkedlist"
	"github.com/songzhibin97/go-baseutils/structure/maps"
)

// Assert Map implementation
var _ maps.Map[int, any] = (*Map[int, any])(nil)

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map[K comparable, V any] struct {
	table    map[K]V
	ordering *doublylinkedlist.List[K]
}

// New instantiates a linked-hash-map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		table:    make(map[K]V),
		ordering: doublylinkedlist.New[K](),
	}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	if _, contains := m.table[key]; !contains {
		m.ordering.Append(key)
	}
	m.table[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	value, found = m.table[key]
	return
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.ordering.Size()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.ordering.Values()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.table = make(map[K]V)
	m.ordering.Clear()
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	bf := strings.Builder{}
	bf.WriteString("LinkedHashMap\nmap[")
	it := m.Iterator()
	for it.Next() {
		bf.WriteString(fmt.Sprintf("(%v:%v) ", it.Key(), it.Value()))
	}
	bf.WriteString("]")
	return bf.String()
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, V]) UnmarshalJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	index := make(map[K]int)
	var keys []K
	for key := range elements {
		keys = append(keys, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(data, esc)
	}

	byIndex := func(a, b K) int {
		key1 := a
		key2 := b
		index1 := index[key1]
		index2 := index[key2]
		return index1 - index2
	}

	bcomparator.Sort(keys, byIndex)

	m.Clear()

	for _, key := range keys {
		m.Put(key, elements[key])
	}

	return nil
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	it := m.Iterator()
	lastIndex := m.Size() - 1
	index := 0

	for it.Next() {
		km, err := json.Marshal(it.Key())
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(it.Value())
		if err != nil {
			return nil, err
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}

	buf.WriteRune('}')

	return buf.Bytes(), nil
}
