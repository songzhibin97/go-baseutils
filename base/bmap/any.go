package bmap

import (
	"encoding/json"
	"reflect"
	"sync"

	"github.com/songzhibin97/go-baseutils/base/bternaryexpr"
)

// =====================================================================================================================
// unsafe

func NewUnsafeAnyBMap[K comparable, V any]() *UnsafeAnyBMap[K, V] {
	return &UnsafeAnyBMap[K, V]{mp: map[K]V{}}
}

func NewUnsafeAnyBMapByMap[K comparable, V any](mp map[K]V) *UnsafeAnyBMap[K, V] {
	if mp == nil {
		return NewUnsafeAnyBMap[K, V]()
	}
	return &UnsafeAnyBMap[K, V]{mp: mp}
}

type UnsafeAnyBMap[K comparable, V any] struct {
	mp map[K]V
}

func (x *UnsafeAnyBMap[K, V]) ToMetaMap() map[K]V {
	return x.mp
}

func (x *UnsafeAnyBMap[K, V]) Keys() []K {
	return Keys(x.mp)
}

func (x *UnsafeAnyBMap[K, V]) Values() []V {
	return Values(x.mp)
}

func (x *UnsafeAnyBMap[K, V]) EqualFuncByMap(m map[K]V, eq func(V1 V, V2 V) bool) bool {
	return EqualFunc[map[K]V, map[K]V, K, V, V](x.mp, m, eq)
}

func (x *UnsafeAnyBMap[K, V]) EqualFuncByBMap(m AnyBMap[K, V], eq func(V1 V, V2 V) bool) bool {
	return EqualFunc[map[K]V, map[K]V, K, V, V](x.mp, m.ToMetaMap(), eq)
}

func (x *UnsafeAnyBMap[K, V]) Clear() {
	x.mp = make(map[K]V)
}

func (x *UnsafeAnyBMap[K, V]) CloneToMap() map[K]V {
	return Clone(x.mp)
}

func (x *UnsafeAnyBMap[K, V]) CloneToBMap() AnyBMap[K, V] {
	return NewUnsafeAnyBMapByMap(Clone(x.mp))
}

func (x *UnsafeAnyBMap[K, V]) CopyByMap(dst map[K]V) {
	Copy(dst, x.mp)
}

func (x *UnsafeAnyBMap[K, V]) CopyByBMap(dst AnyBMap[K, V]) {
	Copy(dst.ToMetaMap(), x.mp)
}

func (x *UnsafeAnyBMap[K, V]) DeleteFunc(del func(K, V) bool) {
	DeleteFunc(x.mp, del)
}

func (x *UnsafeAnyBMap[K, V]) Marshal() ([]byte, error) {
	if x.mp == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(x.mp)
}

func (x *UnsafeAnyBMap[K, V]) Unmarshal(data []byte) error {
	if x.mp == nil {
		x.mp = make(map[K]V)
	}
	return json.Unmarshal(data, &x.mp)
}

func (x *UnsafeAnyBMap[K, V]) Size() int {
	return len(x.mp)
}

func (x *UnsafeAnyBMap[K, V]) IsEmpty() bool {
	return len(x.mp) == 0
}

func (x *UnsafeAnyBMap[K, V]) IsExist(k K) bool {
	_, ok := x.mp[k]
	return ok
}

func (x *UnsafeAnyBMap[K, V]) ContainsKey(k K) bool {
	_, ok := x.mp[k]
	return ok
}

func (x *UnsafeAnyBMap[K, V]) ContainsValue(v V) bool {
	for _, v2 := range x.mp {
		if reflect.DeepEqual(v, v2) {
			return true
		}
	}
	return false
}

func (x *UnsafeAnyBMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range x.mp {
		f(k, v)
	}
}

func (x *UnsafeAnyBMap[K, V]) Get(k K) (V, bool) {
	v, ok := x.mp[k]
	return v, ok
}

func (x *UnsafeAnyBMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	v, ok := x.mp[k]
	return bternaryexpr.TernaryExpr(ok, v, defaultValue)
}

func (x *UnsafeAnyBMap[K, V]) Put(k K, v V) {
	x.mp[k] = v
}

func (x *UnsafeAnyBMap[K, V]) PuTIfAbsent(k K, v V) bool {
	_, ok := x.mp[k]
	if ok {
		return false
	}
	x.mp[k] = v
	return true
}

func (x *UnsafeAnyBMap[K, V]) Delete(k K) {
	delete(x.mp, k)
}

func (x *UnsafeAnyBMap[K, V]) DeleteIfPresent(k K) (V, bool) {
	v, ok := x.mp[k]
	if ok {
		delete(x.mp, k)
	}
	return v, ok
}

func (x *UnsafeAnyBMap[K, V]) MergeByMap(m map[K]V, f func(K, V) bool) {
	for k, v := range m {
		ov, ok := x.mp[k]
		if !ok {
			x.mp[k] = v
			continue
		}
		if f != nil && f(k, ov) {
			x.mp[k] = v
		}
	}
}

func (x *UnsafeAnyBMap[K, V]) MergeByBMap(m AnyBMap[K, V], f func(K, V) bool) {
	for k, v := range m.ToMetaMap() {
		ov, ok := x.mp[k]
		if !ok {
			x.mp[k] = v
			continue
		}
		if f != nil && f(k, ov) {
			x.mp[k] = v
		}
	}
}

func (x *UnsafeAnyBMap[K, V]) Replace(k K, ov, nv V) bool {
	v, ok := x.mp[k]
	flag := ok && reflect.DeepEqual(v, ov)
	if flag {
		x.mp[k] = nv
	}
	return flag
}

// =====================================================================================================================
// safe

func NewSafeAnyBMap[K comparable, V any]() *SafeAnyBMap[K, V] {
	return &SafeAnyBMap[K, V]{mp: NewUnsafeAnyBMap[K,V]()}
}

func NewSafeAnyBMapByMap[K comparable, V any](mp map[K]V) *SafeAnyBMap[K, V] {
	if mp == nil {
		return NewSafeAnyBMap[K, V]()
	}
	return &SafeAnyBMap[K, V]{mp: NewUnsafeAnyBMapByMap[K,V](mp)}
}

type SafeAnyBMap[K comparable, V any] struct {
	mp  *UnsafeAnyBMap[K, V]
	rwl sync.RWMutex
}

func (x *SafeAnyBMap[K, V]) ToMetaMap() map[K]V {
	return x.mp.mp
}

func (x *SafeAnyBMap[K, V]) Keys() []K {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.mp.Keys()
}

func (x *SafeAnyBMap[K, V]) Values() []V {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.mp.Values()
}

func (x *SafeAnyBMap[K, V]) EqualFuncByMap(m map[K]V, eq func(V1 V, V2 V) bool) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.EqualFuncByMap(m, eq)
}

func (x *SafeAnyBMap[K, V]) EqualFuncByBMap(m AnyBMap[K, V], eq func(V1 V, V2 V) bool) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.EqualFuncByBMap(m, eq)
}

func (x *SafeAnyBMap[K, V]) Clear() {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.mp.Clear()
}

func (x *SafeAnyBMap[K, V]) CloneToMap() map[K]V {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.CloneToMap()
}

func (x *SafeAnyBMap[K, V]) CloneToBMap() AnyBMap[K, V] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.CloneToBMap()
}

func (x *SafeAnyBMap[K, V]) CopyByMap(dst map[K]V) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	x.mp.CopyByMap(dst)
}

func (x *SafeAnyBMap[K, V]) CopyByBMap(dst AnyBMap[K, V]) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	x.mp.CopyByBMap(dst)
}

func (x *SafeAnyBMap[K, V]) DeleteFunc(del func(K, V) bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.mp.DeleteFunc(del)
}

func (x *SafeAnyBMap[K, V]) Marshal() ([]byte, error) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.Marshal()
}

func (x *SafeAnyBMap[K, V]) Unmarshal(data []byte) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.mp.Unmarshal(data)
}

func (x *SafeAnyBMap[K, V]) Size() int {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.Size()
}

func (x *SafeAnyBMap[K, V]) IsEmpty() bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.IsEmpty()
}

func (x *SafeAnyBMap[K, V]) IsExist(k K) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.IsExist(k)
}

func (x *SafeAnyBMap[K, V]) ContainsKey(k K) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.ContainsKey(k)
}

func (x *SafeAnyBMap[K, V]) ContainsValue(v V) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.ContainsValue(v)
}

func (x *SafeAnyBMap[K, V]) ForEach(f func(K, V)) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	x.mp.ForEach(f)
}

func (x *SafeAnyBMap[K, V]) Get(k K) (V, bool) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.Get(k)
}

func (x *SafeAnyBMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.mp.GetOrDefault(k, defaultValue)
}

func (x *SafeAnyBMap[K, V]) Put(k K, v V) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.mp.Put(k, v)
}

func (x *SafeAnyBMap[K, V]) PuTIfAbsent(k K, v V) bool {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.mp.PuTIfAbsent(k, v)
}

func (x *SafeAnyBMap[K, V]) Delete(k K) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.mp.Delete(k)
}

func (x *SafeAnyBMap[K, V]) DeleteIfPresent(k K) (V, bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.mp.DeleteIfPresent(k)
}

func (x *SafeAnyBMap[K, V]) MergeByMap(m map[K]V, f func(K, V) bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.mp.MergeByMap(m, f)
}

func (x *SafeAnyBMap[K, V]) MergeByBMap(m AnyBMap[K, V], f func(K, V) bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.mp.MergeByBMap(m, f)
}

func (x *SafeAnyBMap[K, V]) Replace(k K, ov, nv V) bool {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.mp.Replace(k, ov, nv)
}
