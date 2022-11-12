package treeset

import (
	"github.com/songzhibin97/go-baseutils/structure/containers"
	"github.com/songzhibin97/go-baseutils/structure/trees/redblacktree"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithIndex[any] = (*Iterator[any])(nil)

// Iterator returns a stateful iterator whose values can be fetched by an index.
type Iterator[E any] struct {
	index    int
	iterator redblacktree.Iterator[E, struct{}]
	tree     *redblacktree.Tree[E, struct{}]
}

// Iterator holding the iterator's state
func (set *Set[E]) Iterator() Iterator[E] {
	return Iterator[E]{index: -1, iterator: set.tree.Iterator(), tree: set.tree}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Next() bool {
	if iterator.index < iterator.tree.Size() {
		iterator.index++
	}
	return iterator.iterator.Next()
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.iterator.Prev()
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[E]) Value() E {
	return iterator.iterator.Key()
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
	iterator.iterator.Begin()
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[E]) End() {
	iterator.index = iterator.tree.Size()
	iterator.iterator.End()
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
