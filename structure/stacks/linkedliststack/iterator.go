package linkedliststack

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Assert Iterator implementation
var _ containers.IteratorWithIndex[any] = (*Iterator[any])(nil)

// Iterator returns a stateful iterator whose values can be fetched by an index.
type Iterator[E any] struct {
	stack *Stack[E]
	index int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (stack *Stack[E]) Iterator() Iterator[E] {
	return Iterator[E]{stack: stack, index: -1}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[E]) Value() E {
	value, _ := iterator.stack.list.Get(iterator.index) // in reverse (LIFO)
	return value
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
