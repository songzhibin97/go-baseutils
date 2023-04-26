package lscq

import (
	"fmt"
	"unsafe"

	"github.com/songzhibin97/go-baseutils/structure/queues"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

type Queue[E any] struct {
	q    *PointerQueue
	zero E
}

func New[E any]() *Queue[E] {
	return &Queue[E]{
		q: NewPointer(),
	}
}

func (q *Queue[E]) Enqueue(value E) {
	_ = fmt.Sprintf("%p", &value) // TODO make generic variables escape
	q.q.Enqueue(unsafe.Pointer(&value))
}

func (q *Queue[E]) Dequeue() (value E, ok bool) {
	data, ok := q.q.Dequeue()
	if !ok {
		return q.zero, ok
	}
	return *(*E)(data), ok
}

func (q *Queue[E]) Peek() (value E, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (q *Queue[E]) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (q *Queue[E]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (q *Queue[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (q *Queue[E]) Values() []E {
	//TODO implement me
	panic("implement me")
}

func (q *Queue[E]) String() string {
	//TODO implement me
	panic("implement me")
}
