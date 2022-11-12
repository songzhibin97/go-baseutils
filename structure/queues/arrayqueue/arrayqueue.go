// Package arrayqueue implements a queue backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
package arrayqueue

import (
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/structure/lists/arraylist"
	"github.com/songzhibin97/go-baseutils/structure/queues"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

// Queue holds elements in an array-list
type Queue[E any] struct {
	list *arraylist.List[E]
}

// New instantiates a new empty queue
func New[E any]() *Queue[E] {
	return &Queue[E]{list: arraylist.New[E]()}
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[E]) Enqueue(value E) {
	queue.list.Add(value)
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[E]) Dequeue() (value E, ok bool) {
	value, ok = queue.list.Get(0)
	if ok {
		queue.list.Remove(0)
	}
	return
}

// Peek returns first element of the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[E]) Peek() (value E, ok bool) {
	return queue.list.Get(0)
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[E]) Empty() bool {
	return queue.list.Empty()
}

// Size returns number of elements within the queue.
func (queue *Queue[E]) Size() int {
	return queue.list.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[E]) Clear() {
	queue.list.Clear()
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue[E]) Values() []E {
	return queue.list.Values()
}

// String returns a string representation of container
func (queue *Queue[E]) String() string {
	b := strings.Builder{}
	b.WriteString("ArrayQueue\n")
	for index, value := range queue.list.Values() {
		b.WriteString(fmt.Sprintf("(index:%d value:%v) ", index, value))
	}
	return b.String()
}

// Check that the index is within bounds of the list
func (queue *Queue[E]) withinRange(index int) bool {
	return index >= 0 && index < queue.list.Size()
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue[E]) UnmarshalJSON(bytes []byte) error {
	return queue.list.UnmarshalJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue[E]) MarshalJSON() ([]byte, error) {
	return queue.list.MarshalJSON()
}
