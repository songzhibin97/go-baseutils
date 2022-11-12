// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"encoding/json"
	"fmt"

	"github.com/songzhibin97/go-baseutils/base/anytostring"
	"github.com/songzhibin97/go-baseutils/structure/maps"
)

// Assert Map implementation
var _ maps.Map[int, any] = (*Map[int, any])(nil)

// Map holds the elements in go's native map
type Map[K comparable, V any] struct {
	m map[K]V
}

// New instantiates a hash map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]V)}
}

// Put inserts element into the map.
func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values (random order).
func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.m = make(map[K]V)
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(bytes, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.m[key] = value
		}
	}
	return err
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for key, value := range m.m {
		elements[anytostring.AnyToString(key)] = value
	}
	return json.Marshal(&elements)
}
