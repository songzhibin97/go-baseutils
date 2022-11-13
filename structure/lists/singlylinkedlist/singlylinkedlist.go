// Package singlylinkedlist implements the singly-linked list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package singlylinkedlist

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

// List holds the elements, where each element points to the next element
type List[E any] struct {
	first *element[E]
	last  *element[E]
	size  int
	zero  E
}

type element[E any] struct {
	value E
	next  *element[E]
}

// New instantiates a new list and adds the passed values, if any, to the list
func New[E any](values ...E) *List[E] {
	list := &List[E]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List[E]) Add(values ...E) {
	for _, value := range values {
		newElement := &element[E]{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Append appends a value (one or more) at the end of the list (same as Add())
func (list *List[E]) Append(values ...E) {
	list.Add(values...)
}

// Prepend prepends a values (or more)
func (list *List[E]) Prepend(values ...E) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element[E]{value: values[v], next: list.first}
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[E]) Get(index int) (E, bool) {

	if !list.withinRange(index) {
		return list.zero, false
	}

	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}

	return element.value, true
}

// Remove removes the element at the given index from the list.
func (list *List[E]) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeElement *element[E]
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
		beforeElement = element
	}

	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = beforeElement
	}
	if beforeElement != nil {
		beforeElement.next = element.next
	}

	element = nil

	list.size--
}

// Contains checks if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[E]) Contains(values ...E) bool {

	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if reflect.DeepEqual(element.value, value) {
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
func (list *List[E]) Values() []E {
	values := make([]E, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

// IndexOf returns index of provided element
func (list *List[E]) IndexOf(value E) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.Values() {
		if reflect.DeepEqual(element, value) {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (list *List[E]) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List[E]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List[E]) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Sort sort values (in-place) using.
func (list *List[E]) Sort(comparator bcomparator.Comparator[E]) {

	if list.size < 2 {
		return
	}

	values := list.Values()
	bcomparator.Sort(values, comparator)

	list.Clear()

	list.Add(values...)

}

// Swap swaps values of two elements at the given indices.
func (list *List[E]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var element1, element2 *element[E]
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[E]) Insert(index int, values ...E) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	list.size += len(values)

	var beforeElement *element[E]
	foundElement := list.first
	for e := 0; e != index; e, foundElement = e+1, foundElement.next {
		beforeElement = foundElement
	}

	if foundElement == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newElement := &element[E]{value: value}
			if i == 0 {
				list.first = newElement
			} else {
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, value := range values {
			newElement := &element[E]{value: value}
			beforeElement.next = newElement
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	}
}

// Set value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[E]) Set(index int, value E) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}

	foundElement := list.first
	for e := 0; e != index; {
		e, foundElement = e+1, foundElement.next
	}
	foundElement.value = value
}

// String returns a string representation of container
func (list *List[E]) String() string {
	b := strings.Builder{}
	b.WriteString("SinglyLinkedList\n")
	ln := 0
	for element := list.first; element != nil; element = element.next {
		b.WriteString(fmt.Sprintf("(index:%d value:%v) ", ln, element.value))
		ln++
	}
	return b.String()
}

// Check that the index is within bounds of the list
func (list *List[E]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

// UnmarshalJSON @implements json.Unmarshaler
func (list *List[E]) UnmarshalJSON(bytes []byte) error {
	var elements []E
	err := json.Unmarshal(bytes, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}

// MarshalJSON @implements json.Marshaler
func (list *List[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}
