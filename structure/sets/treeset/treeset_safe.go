package treeset

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/sets"
	"sync"
)

var _ sets.Set[any] = (*SetSafe[any])(nil)

func NewSafeWith[E any](comparator bcomparator.Comparator[E], values ...E) *SetSafe[E] {
	return &SetSafe[E]{
		unsafe: NewWith[E](comparator, values...),
	}
}

func NewSafeWithIntComparator(values ...int) *SetSafe[int] {
	return &SetSafe[int]{
		unsafe: NewWithIntComparator(values...),
	}
}

func NewSafeWithStringComparator(values ...string) *SetSafe[string] {
	return &SetSafe[string]{
		unsafe: NewWithStringComparator(values...),
	}
}

type SetSafe[E any] struct {
	unsafe *Set[E]
	lock   sync.Mutex
}

func (s *SetSafe[E]) Add(items ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Add(items...)

}

func (s *SetSafe[E]) Remove(items ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Remove(items...)

}

func (s *SetSafe[E]) Contains(items ...E) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Contains(items...)
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

func (s *SetSafe[E]) Intersection(another *Set[E]) *Set[E] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Intersection(another)
}

func (s *SetSafe[E]) Union(another *Set[E]) *Set[E] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Union(another)
}

func (s *SetSafe[E]) Difference(another *Set[E]) *Set[E] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Difference(another)
}

func (s *SetSafe[E]) UnmarshalJSON(bytes []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(bytes)
}

func (s *SetSafe[E]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
