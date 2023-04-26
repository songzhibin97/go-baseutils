package doublylinkedlist

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/lists"
	"sync"
)

var _ lists.List[any] = (*ListSafe[any])(nil)

func NewSafe[E any](values ...E) *ListSafe[E] {
	return &ListSafe[E]{
		unsafe: New(values...),
	}
}


type ListSafe[E any] struct {
	unsafe *List[E]
	lock   sync.Mutex
}

func (s *ListSafe[E]) Add(values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Add(values...)

}

func (s *ListSafe[E]) Append(values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Append(values...)

}

func (s *ListSafe[E]) Prepend(values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Prepend(values...)

}

func (s *ListSafe[E]) Get(index int) (E, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Get(index)
}

func (s *ListSafe[E]) Remove(index int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Remove(index)

}

func (s *ListSafe[E]) Contains(values ...E) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Contains(values...)
}

func (s *ListSafe[E]) Values() []E {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *ListSafe[E]) IndexOf(value E) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.IndexOf(value)
}

func (s *ListSafe[E]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *ListSafe[E]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *ListSafe[E]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *ListSafe[E]) Sort(comparator bcomparator.Comparator[E]) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Sort(comparator)

}

func (s *ListSafe[E]) Swap(i int, j int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Swap(i, j)

}

func (s *ListSafe[E]) Insert(index int, values ...E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Insert(index, values...)

}

func (s *ListSafe[E]) Set(index int, value E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Set(index, value)

}

func (s *ListSafe[E]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *ListSafe[E]) UnmarshalJSON(bytes []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(bytes)
}

func (s *ListSafe[E]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
