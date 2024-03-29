package skipmap

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/maps"
	"sync"
)

var _ maps.Map[int, any] = (*MapSafe[int, any])(nil)

func NewSafe[K, V any](comparator bcomparator.Comparator[K]) *MapSafe[K, V] {
	return &MapSafe[K, V]{
		unsafe: New[K, V](comparator),
	}
}

type MapSafe[K any, V any] struct {
	unsafe *Map[K, V]
	lock   sync.Mutex
}

func (s *MapSafe[K, V]) Put(key K, value V) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Put(key, value)

}

func (s *MapSafe[K, V]) Get(key K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Get(key)
}

func (s *MapSafe[K, V]) Remove(key K) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Remove(key)

}

func (s *MapSafe[K, V]) Keys() []K {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Keys()
}

func (s *MapSafe[K, V]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *MapSafe[K, V]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *MapSafe[K, V]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *MapSafe[K, V]) Values() []V {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *MapSafe[K, V]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *MapSafe[K, V]) Store(key K, value V) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Store(key, value)

}

func (s *MapSafe[K, V]) Load(key K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Load(key)
}

func (s *MapSafe[K, V]) LoadAndDelete(key K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.LoadAndDelete(key)
}

func (s *MapSafe[K, V]) LoadOrStore(key K, value V) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.LoadOrStore(key, value)
}

func (s *MapSafe[K, V]) LoadOrStoreLazy(key K, f func() V) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.LoadOrStoreLazy(key, f)
}

func (s *MapSafe[K, V]) Delete(key K) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Delete(key)
}

func (s *MapSafe[K, V]) Range(f func(key K, value V) bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Range(f)

}

func (s *MapSafe[K, V]) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Len()
}
