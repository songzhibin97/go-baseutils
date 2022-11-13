// Copyright (c) 2021, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package circularbuffer implements the circular buffer.
//
// In computer science, a circular buffer, circular queue, cyclic buffer or ring buffer is a data structure that uses a single, fixed-size buffer as if it were connected end-to-end. This structure lends itself easily to buffering data streams.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Circular_buffer
package circularbuffer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/songzhibin97/go-baseutils/structure/queues"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

// Queue holds values in a slice.
type Queue[E any] struct {
	values  []E
	start   int
	end     int
	full    bool
	maxSize int
	size    int
	zero    E
}

// New instantiates a new empty queue with the specified size of maximum number of elements that it can hold.
// This max size of the buffer cannot be changed.
func New[E any](maxSize int) *Queue[E] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}
	queue := &Queue[E]{maxSize: maxSize}
	queue.Clear()
	return queue
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[E]) Enqueue(value E) {
	if queue.Full() {
		queue.Dequeue()
	}
	queue.values[queue.end] = value
	queue.end = queue.end + 1
	if queue.end >= queue.maxSize {
		queue.end = 0
	}
	if queue.end == queue.start {
		queue.full = true
	}

	queue.size = queue.calculateSize()
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[E]) Dequeue() (value E, ok bool) {
	if queue.Empty() {
		return queue.zero, false
	}

	value, ok = queue.values[queue.start], true

	if !reflect.DeepEqual(value, queue.zero) {
		queue.values[queue.start] = queue.zero
		queue.start = queue.start + 1
		if queue.start >= queue.maxSize {
			queue.start = 0
		}
		queue.full = false
	}

	queue.size = queue.size - 1

	return
}

// Peek returns first element of the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[E]) Peek() (value E, ok bool) {
	if queue.Empty() {
		return queue.zero, false
	}
	return queue.values[queue.start], true
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[E]) Empty() bool {
	return queue.Size() == 0
}

// Full returns true if the queue is full, i.e. has reached the maximum number of elements that it can hold.
func (queue *Queue[E]) Full() bool {
	return queue.Size() == queue.maxSize
}

// Size returns number of elements within the queue.
func (queue *Queue[E]) Size() int {
	return queue.size
}

// Clear removes all elements from the queue.
func (queue *Queue[E]) Clear() {
	queue.values = make([]E, queue.maxSize, queue.maxSize)
	queue.start = 0
	queue.end = 0
	queue.full = false
	queue.size = 0
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue[E]) Values() []E {
	values := make([]E, queue.Size(), queue.Size())
	for i := 0; i < queue.Size(); i++ {
		values[i] = queue.values[(queue.start+i)%queue.maxSize]
	}
	return values
}

// String returns a string representation of container
func (queue *Queue[E]) String() string {
	b := strings.Builder{}
	b.WriteString("CircularBuffer\n")
	for index, value := range queue.Values() {
		b.WriteString(fmt.Sprintf("(index:%d value:%v) ", index, value))
	}
	return b.String()
}

// Check that the index is within bounds of the list
func (queue *Queue[E]) withinRange(index int) bool {
	return index >= 0 && index < queue.size
}

func (queue *Queue[E]) calculateSize() int {
	if queue.end < queue.start {
		return queue.maxSize - queue.start + queue.end
	} else if queue.end == queue.start {
		if queue.full {
			return queue.maxSize
		}
		return 0
	}
	return queue.end - queue.start
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue[E]) UnmarshalJSON(bytes []byte) error {
	var values []E
	err := json.Unmarshal(bytes, &values)
	if err == nil {
		for _, value := range values {
			queue.Enqueue(value)
		}
	}
	return err
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(queue.values[:queue.maxSize])
}
