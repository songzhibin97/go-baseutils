package skipset

import (
	"github.com/songzhibin97/go-baseutils/structure/sets"
	"sync"
)

var _ sets.Set[any] = (*SetSafe[any])(nil)

type SetSafe[E any] struct {
	unsafe *Set[E]
	lock   sync.Mutex
}

func (s *SetSafe[E]) AddB(value E) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.AddB(value)
}

func (s *SetSafe[E]) Add(values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Add(values...)

}

func (s *SetSafe[E]) ContainsB(value E) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.ContainsB(value)
}

func (s *SetSafe[E]) Contains(values ...E) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Contains(values...)
}

func (s *SetSafe[E]) RemoveB(value E) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RemoveB(value)
}

func (s *SetSafe[E]) Remove(values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Remove(values...)

}

func (s *SetSafe[E]) Range(f func(value E) bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Range(f)

}

func (s *SetSafe[E]) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Len()
}

func (s *SetSafe[E]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *SetSafe[E]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *SetSafe[E]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *SetSafe[E]) Values() []E {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *SetSafe[E]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}
