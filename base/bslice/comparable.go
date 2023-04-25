package bslice

// =====================================================================================================================
// unsafe

func NewUnsafeComparableBSlice[E comparable]() *UnsafeComparableBSlice[E] {
	return &UnsafeComparableBSlice[E]{
		UnsafeAnyBSlice: NewUnsafeAnyBSlice[E](),
	}
}

func NewUnsafeComparableBSliceBySlice[E comparable](s []E) *UnsafeComparableBSlice[E] {
	return &UnsafeComparableBSlice[E]{
		UnsafeAnyBSlice: NewUnsafeAnyBSliceBySlice[E](s),
	}
}

type UnsafeComparableBSlice[E comparable] struct {
	*UnsafeAnyBSlice[E]
}

func (x *UnsafeComparableBSlice[E]) Contains(e E) bool {
	return Contains(x.e, e)
}

func (x *UnsafeComparableBSlice[E]) Equal(es []E) bool {
	return Equal(x.e, es)
}

func (x *UnsafeComparableBSlice[E]) Compact() {
	x.e = Compact(x.e)
}

// =====================================================================================================================
// safe

func NewSafeComparableBSlice[E comparable]() *SafeComparableBSlice[E] {
	return &SafeComparableBSlice[E]{
		SafeAnyBSlice: NewSafeAnyBSlice[E](),
	}
}

func NewSafeComparableBSliceBySlice[E comparable](s []E) *SafeComparableBSlice[E] {
	return &SafeComparableBSlice[E]{
		SafeAnyBSlice: NewSafeAnyBSliceBySlice[E](s),
	}
}

type SafeComparableBSlice[E comparable] struct {
	*SafeAnyBSlice[E]
}

func (x *SafeComparableBSlice[E]) Contains(e E) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return Contains(x.es.e, e)
}

func (x *SafeComparableBSlice[E]) Equal(es []E) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return Equal(x.es.e, es)
}

func (x *SafeComparableBSlice[E]) Compact() {
	x.rwl.Lock()
	defer x.rwl.Unlock()
	x.es.e = Compact(x.es.e)
}
