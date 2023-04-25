package bslice

import (
	"github.com/songzhibin97/go-baseutils/base/btype"
)

// =====================================================================================================================
// unsafe

func NewUnsafeOrderedBSlice[E btype.Ordered]() *UnsafeOrderedBSlice[E] {
	return &UnsafeOrderedBSlice[E]{
		UnsafeComparableBSlice: NewUnsafeComparableBSlice[E](),
	}
}

func NewUnsafeOrderedBSliceBySlice[E btype.Ordered](s []E) *UnsafeOrderedBSlice[E] {
	return &UnsafeOrderedBSlice[E]{
		UnsafeComparableBSlice: NewUnsafeComparableBSliceBySlice[E](s),
	}
}

type UnsafeOrderedBSlice[E btype.Ordered] struct {
	*UnsafeComparableBSlice[E]
}

func (x *UnsafeOrderedBSlice[E]) Compare(s []E) int {
	return Compare(x.UnsafeAnyBSlice.e, s)
}

func (x *UnsafeOrderedBSlice[E]) Sort() {
	Sort(x.UnsafeAnyBSlice.e)
}

func (x *UnsafeOrderedBSlice[E]) IsSorted() bool {
	return IsSorted(x.UnsafeAnyBSlice.e)
}

func (x *UnsafeOrderedBSlice[E]) BinarySearch(target E) (int, bool) {
	return BinarySearch(x.UnsafeAnyBSlice.e, target)
}

// =====================================================================================================================
// safe

func NewSafeOrderedBSlice[E btype.Integer | btype.Float]() *SafeOrderedBSlice[E] {
	return &SafeOrderedBSlice[E]{
		SafeComparableBSlice: NewSafeComparableBSlice[E](),
	}
}

func NewSafeOrderedBSliceBySlice[E btype.Integer | btype.Float](s []E) *SafeOrderedBSlice[E] {
	return &SafeOrderedBSlice[E]{
		SafeComparableBSlice: NewSafeComparableBSliceBySlice[E](s),
	}
}

type SafeOrderedBSlice[E btype.Integer | btype.Float] struct {
	*SafeComparableBSlice[E]
}

func (x *SafeOrderedBSlice[E]) Compare(s []E) int {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return Compare(x.es.e, s)
}

func (x *SafeOrderedBSlice[E]) Sort() {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	Sort(x.es.e)
}

func (x *SafeOrderedBSlice[E]) IsSorted() bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return IsSorted(x.es.e)
}

func (x *SafeOrderedBSlice[E]) BinarySearch(target E) (int, bool) {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return BinarySearch(x.es.e, target)
}
