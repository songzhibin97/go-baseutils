package linkedhashmap

import (
	"github.com/songzhibin97/go-baseutils/structure/maps"
	"sync"
)

var _ maps.Map[int, any] = (*MapSafe[int, any])(nil)

func NewSafe[K comparable, V any]() *MapSafe[K, V] {
	return &MapSafe[K, V]{
		unsafe: New[K, V](),
	}
}

type MapSafe[K comparable, V any] struct {
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

func (s *MapSafe[K, V]) Keys() []K {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Keys()
}

func (s *MapSafe[K, V]) Values() []V {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *MapSafe[K, V]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *MapSafe[K, V]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *MapSafe[K, V]) UnmarshalJSON(data []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.UnmarshalJSON(data)
}

func (s *MapSafe[K, V]) MarshalJSON() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.MarshalJSON()
}
