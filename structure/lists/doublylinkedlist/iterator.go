package doublylinkedlist

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Assert Iterator implementation
var _ containers.ReverseIteratorWithIndex[any] = (*Iterator[any])(nil)

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
	if iterator.index != 0 {
		iterator.element = iterator.element.next
	} else {
		iterator.element = iterator.list.first
	}
	return true
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == iterator.list.size-1 {
		iterator.element = iterator.list.last
	} else {
		iterator.element = iterator.element.prev
	}
	return iterator.list.withinRange(iterator.index)
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

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[E]) End() {
	iterator.index = iterator.list.size
	iterator.element = iterator.list.last
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Last() bool {
	iterator.End()
	return iterator.Prev()
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

// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) PrevTo(f func(index int, value E) bool) bool {
	for iterator.Prev() {
		index, value := iterator.Index(), iterator.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}
