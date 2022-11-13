// Package arraylist implements the array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package arraylist

import (
	"encoding/json"
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/lists"
	"reflect"
	"strings"
)

// Assert List implementation
var _ lists.List[any] = (*List[any])(nil)

// List holds the elements in a slice
type List[E any] struct {
	elements []E
	size     int
	zero     E
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New[E any](values ...E) *List[E] {
	list := &List[E]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value at the end of the list
func (l *List[E]) Add(values ...E) {
	l.growBy(len(values))
	for _, value := range values {
		l.elements[l.size] = value
		l.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (l *List[E]) Get(index int) (E, bool) {

	if !l.withinRange(index) {
		return l.zero, false
	}

	return l.elements[index], true
}

// Remove removes the element at the given index from the list.
func (l *List[E]) Remove(index int) {

	if !l.withinRange(index) {
		return
	}

	l.elements[index] = l.zero                           // cleanup reference
	copy(l.elements[index:], l.elements[index+1:l.size]) // shift to the left by one (slow operation, need ways to optimize this)
	l.size--

	l.shrink()
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (l *List[E]) Contains(values ...E) bool {

	for _, searchValue := range values {
		found := false
		for index := 0; index < l.size; index++ {
			if reflect.DeepEqual(l.elements[index], searchValue) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (l *List[E]) Values() []E {
	newElements := make([]E, l.size, l.size)
	copy(newElements, l.elements[:l.size])
	return newElements
}

// IndexOf returns index of provided element
func (l *List[E]) IndexOf(value E) int {
	if l.size == 0 {
		return -1
	}
	for index, element := range l.elements {
		if reflect.DeepEqual(element, value) {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (l *List[E]) Empty() bool {
	return l.size == 0
}

// Size returns number of elements within the list.
func (l *List[E]) Size() int {
	return l.size
}

// Clear removes all elements from the list.
func (l *List[E]) Clear() {
	l.size = 0
	l.elements = []E{}
}

// Sort sorts values (in-place) using.
func (l *List[E]) Sort(comparator bcomparator.Comparator[E]) {
	if len(l.elements) < 2 {
		return
	}
	bcomparator.Sort(l.elements[:l.size], comparator)
}

// Swap swaps the two values at the specified positions.
func (l *List[E]) Swap(i, j int) {
	if l.withinRange(i) && l.withinRange(j) {
		l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (l *List[E]) Insert(index int, values ...E) {

	if !l.withinRange(index) {
		// Append
		if index == l.size {
			l.Add(values...)
		}
		return
	}

	ln := len(values)
	l.growBy(ln)
	l.size += ln
	copy(l.elements[index+ln:], l.elements[index:l.size-ln])
	copy(l.elements[index:], values)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (l *List[E]) Set(index int, value E) {

	if !l.withinRange(index) {
		// Append
		if index == l.size {
			l.Add(value)
		}
		return
	}

	l.elements[index] = value
}

// String returns a string representation of container
func (l *List[E]) String() string {
	b := strings.Builder{}
	b.WriteString("ArrayList\n")
	for index, value := range l.elements[:l.size] {
		b.WriteString(fmt.Sprintf("(index:%d value:%v) ", index, value))
	}
	return b.String()
}

// Check that the index is within bounds of the list
func (l *List[E]) withinRange(index int) bool {
	return index >= 0 && index < l.size
}

func (l *List[E]) resize(cap int) {
	newElements := make([]E, cap, cap)
	copy(newElements, l.elements)
	l.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (l *List[E]) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(l.elements)
	if l.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		l.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (l *List[E]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(l.elements)
	if l.size <= int(float32(currentCapacity)*shrinkFactor) {
		l.resize(l.size)
	}
}

func (l *List[E]) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &l.elements)
	if err == nil {
		l.size = len(l.elements)
	}
	return err
}

// MarshalJSON @implements json.Marshaler
func (l *List[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.elements[:l.size])
}
