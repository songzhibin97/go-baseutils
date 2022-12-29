package bobjectstorage

import (
	"github.com/songzhibin97/go-baseutils/base/bmap"
	"github.com/songzhibin97/go-baseutils/base/breflect"
)

var (
	anyMap = bmap.NewSafeAnyBMap[string, any]()
)

func Set[T any](key string, val T) {
	anyMap.Put(key, val)
}

func GetSafeAssertion[T any](key string) (T, bool) {
	var zero T
	v, ok := anyMap.Get(key)
	if !ok || breflect.IsNil(v) {
		return zero, false
	}
	nv, ok := v.(T)
	return nv, ok
}

func Get[T any](key string) T {
	var zero T
	v, ok := anyMap.Get(key)
	if !ok || breflect.IsNil(v) {
		return zero
	}
	return v.(T)
}
