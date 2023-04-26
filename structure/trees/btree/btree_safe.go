package btree

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/trees"
	"sync"
)

var _ trees.Tree[any] = (*TreeSafe[any, any])(nil)


func NewSafeWith[K, V any](order int, comparator bcomparator.Comparator[K]) *TreeSafe[K, V] {
	return &TreeSafe[K, V]{
		unsafe: NewWith[K, V](order, comparator),
	}
}

func NewSafeWithIntComparator[V any](order int) *TreeSafe[int, V] {
	return &TreeSafe[int, V]{
		unsafe: NewWithIntComparator[V](order),
	}
}

func NewSafeWithStringComparator[V any](order int) *TreeSafe[string, V] {
	return &TreeSafe[string, V]{
		unsafe: NewWithStringComparator[V](order),
	}
}

type TreeSafe[K, V any] struct {
	unsafe *Tree[K, V]
	lock   sync.Mutex
}

func (s *TreeSafe[K, V]) Put(key K, value V) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Put(key, value)

}

func (s *TreeSafe[K, V]) Get(key K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Get(key)
}

func (s *TreeSafe[K, V]) GetNode(key K) *Node[K, V] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.GetNode(key)
}

func (s *TreeSafe[K, V]) Remove(key K) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Remove(key)

}

func (s *TreeSafe[K, V]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *TreeSafe[K, V]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *TreeSafe[K, V]) Keys() []K {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Keys()
}

func (s *TreeSafe[K, V]) Values() []V {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *TreeSafe[K, V]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *TreeSafe[K, V]) Height() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Height()
}

func (s *TreeSafe[K, V]) Left() *Node[K, V] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Left()
}

func (s *TreeSafe[K, V]) LeftKey() K {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.LeftKey()
}

func (s *TreeSafe[K, V]) LeftValue() V {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.LeftValue()
}

func (s *TreeSafe[K, V]) Right() *Node[K, V] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Right()
}

func (s *TreeSafe[K, V]) RightKey() K {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RightKey()
}

func (s *TreeSafe[K, V]) RightValue() V {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RightValue()
}

func (s *TreeSafe[K, V]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *TreeSafe[K, V]) UnmarshalJSON(bytes []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(bytes)
}

func (s *TreeSafe[K, V]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
