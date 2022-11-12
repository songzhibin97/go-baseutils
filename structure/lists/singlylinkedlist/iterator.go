package singlylinkedlist

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Assert Iterator implementation
var _ containers.IteratorWithIndex[any] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[E any] struct {
	list    *List[E]
	index   int
	element *element[E]
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[E]) Iterator() Iterator[E] {
	return Iterator[E]{list: list, index: -1, element: nil}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == 0 {
		iterator.element = iterator.list.first
	} else {
		iterator.element = iterator.element.next
	}
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[E]) Value() E {
	return iterator.element.value
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator[E]) Index() int {
	return iterator.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[E]) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) NextTo(f func(index int, value E) bool) bool {
	for iterator.Next() {
		index, value := iterator.Index(), iterator.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}
