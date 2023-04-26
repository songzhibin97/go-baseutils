package priorityqueue

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/queues"
	"sync"
)

var _ queues.Queue[any] = (*QueueSafe[any])(nil)

func NewSafeWith[E any](comparator bcomparator.Comparator[E]) *QueueSafe[E] {
	return &QueueSafe[E]{
		unsafe: NewWith[E](comparator),
	}
}

type QueueSafe[E any] struct {
	unsafe *Queue[E]
	lock   sync.Mutex
}

func (s *QueueSafe[E]) Enqueue(value E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Enqueue(value)

}

func (s *QueueSafe[E]) Dequeue() (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Dequeue()
}

func (s *QueueSafe[E]) Peek() (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Peek()
}

func (s *QueueSafe[E]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *QueueSafe[E]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *QueueSafe[E]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *QueueSafe[E]) Values() []E {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *QueueSafe[E]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *QueueSafe[E]) UnmarshalJSON(bytes []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(bytes)
}

func (s *QueueSafe[E]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
