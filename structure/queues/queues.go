// Package queues provides an abstract Queue interface.
//
// In computer science, a queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence. By convention, the end of the sequence at which elements are added is called the back, tail, or rear of the queue, and the end at which elements are removed is called the head or front of the queue, analogously to the words used when people line up to wait for goods or services.
// The operation of adding an element to the rear of the queue is known as enqueue, and the operation of removing an element from the front is known as dequeue. Other operations may also be allowed, often including a peek or front operation that returns the value of the next element to be dequeued without remove it.
//
// Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
package queues

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Queue interface that all queues implement
type Queue[E any] interface {
	Enqueue(value E)
	Dequeue() (value E, ok bool)
	Peek() (value E, ok bool)

	containers.Container[E]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []E{}
	// String() string
}
