package binaryheap

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/trees"
	"sync"
)

var _ trees.Tree[any] = (*HeapSafe[any])(nil)

func NewSafeWith[E any](comparator bcomparator.Comparator[E]) *HeapSafe[E] {
	return &HeapSafe[E]{
		unsafe: NewWith[E](comparator),
	}
}

func NewSafeWithIntComparator() *HeapSafe[int] {
	return &HeapSafe[int]{
		unsafe: NewWithIntComparator(),
	}
}

func NewSafeWithStringComparator() *HeapSafe[string] {
	return &HeapSafe[string]{
		unsafe: NewWithStringComparator(),
	}
}

type HeapSafe[E any] struct {
	unsafe *Heap[E]
	lock   sync.Mutex
}

func (s *HeapSafe[E]) Push(values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Push(values...)

}

func (s *HeapSafe[E]) Pop() (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Pop()
}

func (s *HeapSafe[E]) Peek() (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Peek()
}

func (s *HeapSafe[E]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *HeapSafe[E]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *HeapSafe[E]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *HeapSafe[E]) Values() []E {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *HeapSafe[E]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *HeapSafe[E]) UnmarshalJSON(bytes []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(bytes)
}

func (s *HeapSafe[E]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
