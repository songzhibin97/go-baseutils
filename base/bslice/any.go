package bslice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"sync"
)

// =====================================================================================================================
// unsafe

func NewUnsafeAnyBSlice[E any]() AnyBSlice[E] {
	return &UnsafeAnyBSlice[E]{
		e: []E{},
	}
}

func NewUnsafeAnyBSliceBySlice[E any](s []E) AnyBSlice[E] {
	return &UnsafeAnyBSlice[E]{
		e: s,
	}
}

type UnsafeAnyBSlice[E any] struct {
	e    []E
	zero E
}

func (x *UnsafeAnyBSlice[E]) EqualFunc(es []E, f func(E, E) bool) bool {
	return EqualFunc(x.e, es, f)
}

func (x *UnsafeAnyBSlice[E]) CompareFunc(es []E, f func(E, E) int) int {
	return CompareFunc(x.e, es, f)
}

func (x *UnsafeAnyBSlice[E]) IndexFunc(f func(E) bool) int {
	return IndexFunc(x.e, f)
}

func (x *UnsafeAnyBSlice[E]) Insert(i int, e ...E) {
	_ = x.InsertE(i, e...)
}

func (x *UnsafeAnyBSlice[E]) InsertE(i int, e ...E) error {
	if i < 0 || i > len(x.e) {
		return errors.New("insert index out of range")
	}
	x.e = Insert(x.e, i, e...)
	return nil
}

func (x *UnsafeAnyBSlice[E]) Delete(i int, j int) {
	_ = x.DeleteE(i, j)
}

func (x *UnsafeAnyBSlice[E]) DeleteE(i int, j int) error {
	ln := len(x.e)
	if ln == 0 {
		return nil
	}
	if i < 0 || j > ln || i > j {
		return errors.New(fmt.Sprintf("invalid range:%d-%d, ln:%d", i, j, ln))
	}
	x.e = Delete(x.e, i, j)
	return nil
}

func (x *UnsafeAnyBSlice[E]) DeleteToSlice(i int, j int) []E {
	v, _ := x.DeleteToSliceE(i, j)
	return v
}

func (x *UnsafeAnyBSlice[E]) DeleteToSliceE(i int, j int) ([]E, error) {
	ln := len(x.e)
	if ln == 0 {
		return nil, nil
	}
	if i < 0 || j > ln || i > j {
		return nil, errors.New(fmt.Sprintf("invalid range:%d-%d, ln:%d", i, j, ln))
	}
	return Delete(x.e, i, j), nil
}

func (x *UnsafeAnyBSlice[E]) DeleteToBSlice(i int, j int) AnyBSlice[E] {
	v, err := x.DeleteToBSliceE(i, j)
	if err != nil {
		return v
	}
	return v
}

func (x *UnsafeAnyBSlice[E]) DeleteToBSliceE(i int, j int) (AnyBSlice[E], error) {
	ln := len(x.e)
	if ln == 0 {
		return NewUnsafeAnyBSlice[E](), nil
	}
	if i < 0 || j > ln || i > j {
		return NewUnsafeAnyBSlice[E](), errors.New(fmt.Sprintf("invalid range:%d-%d, ln:%d", i, j, ln))
	}
	return NewUnsafeAnyBSliceBySlice(Delete(x.e, i, j)), nil
}

func (x *UnsafeAnyBSlice[E]) Replace(i int, j int, e ...E) {
	_ = x.ReplaceE(i, j, e...)
}

func (x *UnsafeAnyBSlice[E]) ReplaceE(i int, j int, e ...E) error {
	ln := len(x.e)
	if i < 0 || j > ln || i > j {
		return errors.New(fmt.Sprintf("invalid range:%d-%d, ln:%d", i, j, ln))
	}
	x.e = Replace(x.e, i, j, e...)
	return nil
}

func (x *UnsafeAnyBSlice[E]) CloneToSlice() []E {
	return Clone(x.e)
}

func (x *UnsafeAnyBSlice[E]) CloneToBSlice() AnyBSlice[E] {
	return NewUnsafeAnyBSliceBySlice(Clone(x.e))
}

func (x *UnsafeAnyBSlice[E]) CompactFunc(f func(E, E) bool) {
	x.e = CompactFunc(x.e, f)
}

func (x *UnsafeAnyBSlice[E]) Grow(i int) {
	_ = x.GrowE(i)
}

func (x *UnsafeAnyBSlice[E]) GrowE(i int) error {
	if i < 0 {
		return errors.New("insert index out of range")
	}
	x.e = Grow(x.e, i)
	return nil
}

func (x *UnsafeAnyBSlice[E]) Clip() {
	x.e = Clip(x.e)
}

func (x *UnsafeAnyBSlice[E]) ForEach(f func(int, E)) {
	for i, e := range x.e {
		f(i, e)
	}
}

func (x *UnsafeAnyBSlice[E]) SortFunc(f func(i E, j E) bool) {
	SortFunc(x.e, f)
}

func (x *UnsafeAnyBSlice[E]) SortFuncToSlice(f func(i E, j E) bool) []E {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	SortFunc(cp, f)
	return cp
}

func (x *UnsafeAnyBSlice[E]) SortFuncToBSlice(f func(i E, j E) bool) AnyBSlice[E] {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	SortFunc(cp, f)
	return NewUnsafeAnyBSliceBySlice(cp)
}

func (x *UnsafeAnyBSlice[E]) SortComparator(f bcomparator.Comparator[E]) {
	bcomparator.Sort(x.e, f)
}

func (x *UnsafeAnyBSlice[E]) SortComparatorToSlice(f bcomparator.Comparator[E]) []E {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	bcomparator.Sort(cp, f)
	return cp
}

func (x *UnsafeAnyBSlice[E]) SortComparatorToBSlice(f bcomparator.Comparator[E]) AnyBSlice[E] {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	bcomparator.Sort(cp, f)
	return NewUnsafeAnyBSliceBySlice(cp)
}

func (x *UnsafeAnyBSlice[E]) SortStableFunc(f func(i E, j E) bool) {
	SortStableFunc(x.e, f)
}

func (x *UnsafeAnyBSlice[E]) SortStableFuncToSlice(f func(i E, j E) bool) []E {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	SortStableFunc(cp, f)
	return cp
}

func (x *UnsafeAnyBSlice[E]) SortStableFuncToBSlice(f func(i E, j E) bool) AnyBSlice[E] {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	SortStableFunc(cp, f)
	return NewUnsafeAnyBSliceBySlice(cp)
}

func (x *UnsafeAnyBSlice[E]) IsSortedFunc(f func(i E, j E) bool) bool {
	return IsSortedFunc(x.e, f)
}

func (x *UnsafeAnyBSlice[E]) BinarySearchFunc(e E, f func(E, E) int) (int, bool) {
	return BinarySearchFunc(x.e, e, f)
}

func (x *UnsafeAnyBSlice[E]) Filter(f func(E) bool) {
	var res []E
	for _, e := range x.e {
		if f(e) {
			res = append(res, e)
		}
	}
	x.e = res
}

func (x *UnsafeAnyBSlice[E]) FilterToSlice(f func(E) bool) []E {
	var res []E
	for _, e := range x.e {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}

func (x *UnsafeAnyBSlice[E]) FilterToBSlice(f func(E) bool) AnyBSlice[E] {
	var res []E
	for _, e := range x.e {
		if f(e) {
			res = append(res, e)
		}
	}
	return NewUnsafeAnyBSliceBySlice(res)
}

func (x *UnsafeAnyBSlice[E]) Reverse() {
	l, r := 0, len(x.e)-1
	for l < r {
		x.e[l], x.e[r] = x.e[r], x.e[l]
		l++
		r--
	}
}

func (x *UnsafeAnyBSlice[E]) ReverseToSlice() []E {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	l, r := 0, len(cp)-1
	for l < r {
		cp[l], cp[r] = cp[r], cp[l]
		l++
		r--
	}
	return cp
}

func (x *UnsafeAnyBSlice[E]) ReverseToBSlice() AnyBSlice[E] {
	es := x.e
	cp := make([]E, len(es))
	copy(cp, es)
	l, r := 0, len(cp)-1
	for l < r {
		cp[l], cp[r] = cp[r], cp[l]
		l++
		r--
	}
	return NewUnsafeAnyBSliceBySlice(cp)
}

func (x *UnsafeAnyBSlice[E]) Marshal() ([]byte, error) {
	if x.e == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(x.e)
}

func (x *UnsafeAnyBSlice[E]) Unmarshal(data []byte) error {
	if x.e == nil {
		x.e = []E{}
	}
	return json.Unmarshal(data, &x.e)
}

func (x *UnsafeAnyBSlice[E]) Len() int {
	return len(x.e)
}

func (x *UnsafeAnyBSlice[E]) Cap() int {
	return cap(x.e)
}

func (x *UnsafeAnyBSlice[E]) ToInterfaceSlice() []interface{} {
	res := make([]interface{}, len(x.e))
	for i, e := range x.e {
		res[i] = e
	}
	return res
}

func (x *UnsafeAnyBSlice[E]) ToMetaSlice() []E {
	return x.e
}

func (x *UnsafeAnyBSlice[E]) Swap(i, j int) {
	ln := len(x.e)
	if i < 0 || j < 0 || i >= ln || j >= ln || i == j {
		return
	}
	x.e[i], x.e[j] = x.e[j], x.e[i]
}

func (x *UnsafeAnyBSlice[E]) Clear() {
	x.e = []E{}
}

func (x *UnsafeAnyBSlice[E]) Append(es ...E) {
	x.e = append(x.e, es...)
}

func (x *UnsafeAnyBSlice[E]) AppendToSlice(es ...E) []E {
	nes := x.e
	cp := make([]E, len(nes))
	copy(cp, nes)
	return append(cp, es...)
}

func (x *UnsafeAnyBSlice[E]) AppendToBSlice(es ...E) AnyBSlice[E] {
	nes := x.e
	cp := make([]E, len(nes))
	copy(cp, nes)
	return NewUnsafeAnyBSliceBySlice(append(cp, es...))
}

func (x *UnsafeAnyBSlice[E]) CopyToSlice() []E {
	nes := x.e
	cp := make([]E, len(nes))
	copy(cp, nes)
	return cp
}

func (x *UnsafeAnyBSlice[E]) CopyToBSlice() AnyBSlice[E] {
	nes := x.e
	cp := make([]E, len(nes))
	copy(cp, nes)
	return NewUnsafeAnyBSliceBySlice(cp)
}

func (x *UnsafeAnyBSlice[E]) GetByIndex(index int) E {
	v, _ := x.GetByIndexE(index)
	return v
}

func (x *UnsafeAnyBSlice[E]) GetByIndexE(index int) (E, error) {
	if index < 0 || index >= len(x.e) {
		return x.zero, errors.New("index out of range")
	}
	return x.e[index], nil
}

func (x *UnsafeAnyBSlice[E]) GetByIndexOrDefault(index int, defaultE E) E {
	if index < 0 || index >= len(x.e) {
		return defaultE
	}
	return x.e[index]
}

func (x *UnsafeAnyBSlice[E]) GetByRange(start int, end int) []E {
	v, _ := x.GetByRangeE(start, end)
	return v
}

func (x *UnsafeAnyBSlice[E]) GetByRangeE(start int, end int) ([]E, error) {
	ln := len(x.e)
	if start < 0 || end > ln || start > end {
		return nil, errors.New(fmt.Sprintf("invalid range:%d-%d, ln:%d", start, end, ln))
	}
	return x.e[start:end], nil
}

func (x *UnsafeAnyBSlice[E]) SetByIndex(index int, e E) {
	_ = x.SetByIndexE(index, e)

}

func (x *UnsafeAnyBSlice[E]) SetByIndexE(index int, e E) error {
	if index < 0 || index >= len(x.e) {
		return errors.New("index out of range")
	}
	x.e[index] = e
	return nil
}

func (x *UnsafeAnyBSlice[E]) SetByRange(index int, es []E) {
	_ = x.SetByRangeE(index, es)
}

func (x *UnsafeAnyBSlice[E]) SetByRangeE(index int, es []E) error {
	if index < 0 {
		return errors.New("insert index out of range")
	}
	total := index + len(es)
	if total > cap(x.e) {
		x.Append(es...)
		return nil
	}
	copy(x.e[index:total], es)
	return nil
}

// =====================================================================================================================
// safe

func NewSafeAnyBSlice[E any]() AnyBSlice[E] {
	return &UnsafeAnyBSlice[E]{
		e: []E{},
	}
}

func NewSafeAnyBSliceBySlice[E any](e []E) AnyBSlice[E] {
	return &SafeAnyBSlice[E]{
		es: &UnsafeAnyBSlice[E]{
			e: e,
		},
	}
}

type SafeAnyBSlice[E any] struct {
	es  *UnsafeAnyBSlice[E]
	rwl sync.RWMutex
}

func (x *SafeAnyBSlice[E]) Marshal() ([]byte, error) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.Marshal()
}

func (x *SafeAnyBSlice[E]) Unmarshal(data []byte) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.Unmarshal(data)
}

func (x *SafeAnyBSlice[E]) Len() int {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.Len()
}

func (x *SafeAnyBSlice[E]) Cap() int {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.Cap()
}

func (x *SafeAnyBSlice[E]) ToInterfaceSlice() []interface{} {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.ToInterfaceSlice()
}

func (x *SafeAnyBSlice[E]) Swap(i, j int) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Swap(i, j)
}

func (x *SafeAnyBSlice[E]) Clear() {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Clear()
}

func (x *SafeAnyBSlice[E]) EqualFunc(es []E, f func(E, E) bool) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.EqualFunc(es, f)
}

func (x *SafeAnyBSlice[E]) CompareFunc(es []E, f func(E, E) int) int {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.CompareFunc(es, f)
}

func (x *SafeAnyBSlice[E]) IndexFunc(f func(E) bool) int {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.IndexFunc(f)
}

func (x *SafeAnyBSlice[E]) Insert(i int, e ...E) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Insert(i, e...)
}

func (x *SafeAnyBSlice[E]) InsertE(i int, e ...E) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.InsertE(i, e...)
}

func (x *SafeAnyBSlice[E]) Delete(i int, i2 int) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Delete(i, i2)
}

func (x *SafeAnyBSlice[E]) DeleteE(i int, i2 int) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.DeleteE(i, i2)
}

func (x *SafeAnyBSlice[E]) DeleteToSlice(i int, i2 int) []E {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.DeleteToSlice(i, i2)
}

func (x *SafeAnyBSlice[E]) DeleteToSliceE(i int, i2 int) ([]E, error) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.DeleteToSliceE(i, i2)
}

func (x *SafeAnyBSlice[E]) DeleteToBSlice(i int, i2 int) AnyBSlice[E] {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.DeleteToBSlice(i, i2)
}

func (x *SafeAnyBSlice[E]) DeleteToBSliceE(i int, i2 int) (AnyBSlice[E], error) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.DeleteToBSliceE(i, i2)
}

func (x *SafeAnyBSlice[E]) Replace(i int, i2 int, e ...E) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Replace(i, i2, e...)
}

func (x *SafeAnyBSlice[E]) ReplaceE(i int, i2 int, e ...E) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.ReplaceE(i, i2, e...)
}

func (x *SafeAnyBSlice[E]) CloneToSlice() []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.CloneToSlice()
}

func (x *SafeAnyBSlice[E]) CloneToBSlice() AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.CloneToBSlice()
}

func (x *SafeAnyBSlice[E]) CompactFunc(f func(E, E) bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.CompactFunc(f)
}

func (x *SafeAnyBSlice[E]) Grow(i int) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Grow(i)
}

func (x *SafeAnyBSlice[E]) GrowE(i int) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.GrowE(i)
}

func (x *SafeAnyBSlice[E]) Clip() {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Clip()
}

func (x *SafeAnyBSlice[E]) ForEach(f func(int, E)) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	x.es.ForEach(f)
}

func (x *SafeAnyBSlice[E]) SortFunc(f func(i E, j E) bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.SortFunc(f)
}

func (x *SafeAnyBSlice[E]) SortFuncToSlice(f func(i E, j E) bool) []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.SortFuncToSlice(f)
}

func (x *SafeAnyBSlice[E]) SortFuncToBSlice(f func(i E, j E) bool) AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.SortFuncToBSlice(f)
}

func (x *SafeAnyBSlice[E]) SortComparator(f bcomparator.Comparator[E]) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.SortComparator(f)
}

func (x *SafeAnyBSlice[E]) SortComparatorToSlice(f bcomparator.Comparator[E]) []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.SortComparatorToSlice(f)
}

func (x *SafeAnyBSlice[E]) SortComparatorToBSlice(f bcomparator.Comparator[E]) AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.SortComparatorToBSlice(f)
}

func (x *SafeAnyBSlice[E]) SortStableFunc(f func(i E, j E) bool) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.SortStableFunc(f)
}

func (x *SafeAnyBSlice[E]) SortStableFuncToSlice(f func(i E, j E) bool) []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.SortStableFuncToSlice(f)
}

func (x *SafeAnyBSlice[E]) SortStableFuncToBSlice(f func(i E, j E) bool) AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.SortStableFuncToBSlice(f)
}

func (x *SafeAnyBSlice[E]) IsSortedFunc(f func(i E, j E) bool) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.IsSortedFunc(f)
}

func (x *SafeAnyBSlice[E]) BinarySearchFunc(e E, f func(E, E) int) (int, bool) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.BinarySearchFunc(e, f)
}

func (x *SafeAnyBSlice[E]) Filter(f func(E) bool) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	x.es.Filter(f)
}

func (x *SafeAnyBSlice[E]) FilterToSlice(f func(E) bool) []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.FilterToSlice(f)
}

func (x *SafeAnyBSlice[E]) FilterToBSlice(f func(E) bool) AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.FilterToBSlice(f)
}

func (x *SafeAnyBSlice[E]) Reverse() {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Reverse()
}

func (x *SafeAnyBSlice[E]) ReverseToSlice() []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.ReverseToSlice()
}

func (x *SafeAnyBSlice[E]) ReverseToBSlice() AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.ReverseToBSlice()
}

func (x *SafeAnyBSlice[E]) ToMetaSlice() []E {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.ToMetaSlice()
}

func (x *SafeAnyBSlice[E]) Append(e ...E) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.Append(e...)
}

func (x *SafeAnyBSlice[E]) AppendToSlice(e ...E) []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.AppendToSlice(e...)
}

func (x *SafeAnyBSlice[E]) AppendToBSlice(e ...E) AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.AppendToBSlice(e...)
}

func (x *SafeAnyBSlice[E]) CopyToSlice() []E {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.CopyToSlice()
}

func (x *SafeAnyBSlice[E]) CopyToBSlice() AnyBSlice[E] {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.CopyToBSlice()
}

func (x *SafeAnyBSlice[E]) GetByIndex(index int) E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.GetByIndex(index)
}

func (x *SafeAnyBSlice[E]) GetByIndexE(index int) (E, error) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.GetByIndexE(index)
}

func (x *SafeAnyBSlice[E]) GetByIndexOrDefault(index int, defaultE E) E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.GetByIndexOrDefault(index, defaultE)
}

func (x *SafeAnyBSlice[E]) GetByRange(start int, end int) []E {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.GetByRange(start, end)
}

func (x *SafeAnyBSlice[E]) GetByRangeE(start int, end int) ([]E, error) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return x.es.GetByRangeE(start, end)
}

func (x *SafeAnyBSlice[E]) SetByIndex(index int, e E) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.SetByIndex(index, e)

}

func (x *SafeAnyBSlice[E]) SetByIndexE(index int, e E) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.SetByIndexE(index, e)
}

func (x *SafeAnyBSlice[E]) SetByRange(index int, es []E) {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.SetByRange(index, es)
}

func (x *SafeAnyBSlice[E]) SetByRangeE(index int, es []E) error {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	return x.es.SetByRangeE(index, es)
}
