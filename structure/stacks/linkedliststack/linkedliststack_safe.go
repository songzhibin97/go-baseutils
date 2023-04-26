package linkedliststack

import (
	"github.com/songzhibin97/go-baseutils/structure/stacks"
	"sync"
)

var _ stacks.Stack[any] = (*StackSafe[any])(nil)

type StackSafe[E any] struct {
	unsafe *Stack[E]
	lock   sync.Mutex
}

func (s *StackSafe[E]) Push(value E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Push(value)

}

func (s *StackSafe[E]) Pop() (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Pop()
}

func (s *StackSafe[E]) Peek() (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Peek()
}

func (s *StackSafe[E]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *StackSafe[E]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *StackSafe[E]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *StackSafe[E]) Values() []E {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *StackSafe[E]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *StackSafe[E]) UnmarshalJSON(bytes []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(bytes)
}

func (s *StackSafe[E]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
