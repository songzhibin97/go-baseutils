package maps

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Map interface that all maps implement
type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []E
	// String() string
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap[K comparable, V comparable] interface {
	GetKey(value V) (key K, found bool)

	Map[K, V]
}
