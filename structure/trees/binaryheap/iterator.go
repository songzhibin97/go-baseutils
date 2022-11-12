package binaryheap

import (
	"github.com/songzhibin97/go-baseutils/structure/containers"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithIndex[any] = (*Iterator[any])(nil)

// Iterator returns a stateful iterator whose values can be fetched by an index.
type Iterator[E any] struct {
	heap  *Heap[E]
	index int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (heap *Heap[E]) Iterator() Iterator[E] {
	return Iterator[E]{heap: heap, index: -1}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Next() bool {
	if iterator.index < iterator.heap.Size() {
		iterator.index++
	}
	return iterator.heap.withinRange(iterator.index)
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[E]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.heap.withinRange(iterator.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[E]) Value() E {
	start, end := evaluateRange(iterator.index)
	if end > iterator.heap.Size() {
		end = iterator.heap.Size()
	}
	tmpHeap := NewWith(iterator.heap.Comparator)
	for n := start; n < end; n++ {
		value, _ := iterator.heap.list.Get(n)
		tmpHeap.Push(value)
	}
	for n := 0; n < iterator.index-start; n++ {
		tmpHeap.Pop()
	}
	value, _ := tmpHeap.Pop()
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

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[E]) End() {
	iterator.index = iterator.heap.Size()
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

// numOfBits counts the number of bits of an int
func numOfBits(n int) uint {
	var count uint
	for n != 0 {
		count++
		n >>= 1
	}
	return count
}

// evaluateRange evaluates the index range [start,end) of same level nodes in the heap as the index
func evaluateRange(index int) (start int, end int) {
	bits := numOfBits(index+1) - 1
	start = 1<<bits - 1
	end = start + 1<<bits
	return
}
