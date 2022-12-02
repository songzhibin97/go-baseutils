package bcache

import "time"

type Cache[K comparable, V any] interface {
	Set(k K, v V, d time.Duration)
	SetDefault(k K, v V)
	SetNoExpire(k K, v V)
	SetIfAbsent(k K, v V, d time.Duration) bool
	Replace(k K, v V, d time.Duration) bool
	Delete(k K)
	Get(k K) (V, bool)
	GetWithExpire(k K) (V, time.Time, bool)
	Count() int
	Clear()
	Load(data []byte) error
	Export() ([]byte, error)
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}
