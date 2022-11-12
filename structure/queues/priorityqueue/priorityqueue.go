// Package priorityqueue implements a priority queue backed by binary queue.
//
// An unbounded priority queue based on a priority queue.
// The elements of the priority queue are ordered by a comparator provided at queue construction time.
//
// The heap of this queue is the least/smallest element with respect to the specified ordering.
// If multiple elements are tied for least value, the heap is one of those elements arbitrarily.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/Priority_queue
package priorityqueue

import (
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/queues"
	"github.com/songzhibin97/go-baseutils/structure/trees/binaryheap"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

// Queue holds elements in an array-list
type Queue[E any] struct {
	heap       *binaryheap.Heap[E]
	Comparator bcomparator.Comparator[E]
}

// NewWith instantiates a new empty queue with the custom comparator.
func NewWith[E any](comparator bcomparator.Comparator[E]) *Queue[E] {
	return &Queue[E]{heap: binaryheap.NewWith(comparator), Comparator: comparator}
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[E]) Enqueue(value E) {
	queue.heap.Push(value)
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[E]) Dequeue() (value E, ok bool) {
	return queue.heap.Pop()
}

// Peek returns top element on the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[E]) Peek() (value E, ok bool) {
	return queue.heap.Peek()
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[E]) Empty() bool {
	return queue.heap.Empty()
}

// Size returns number of elements within the queue.
func (queue *Queue[E]) Size() int {
	return queue.heap.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[E]) Clear() {
	queue.heap.Clear()
}

// Values returns all elements in the queue.
func (queue *Queue[E]) Values() []E {
	return queue.heap.Values()
}

// String returns a string representation of container
func (queue *Queue[E]) String() string {
	b := strings.Builder{}
	b.WriteString("PriorityQueue\n")
	for index, value := range queue.Values() {
		b.WriteString(fmt.Sprintf("(index:%d value:%v) ", index, value))
	}
	return b.String()
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue[E]) UnmarshalJSON(bytes []byte) error {
	return queue.heap.UnmarshalJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue[E]) MarshalJSON() ([]byte, error) {
	return queue.heap.MarshalJSON()
}
