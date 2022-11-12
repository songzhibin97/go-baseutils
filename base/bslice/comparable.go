package bslice

// =====================================================================================================================
// unsafe

func NewUnsafeComparableBSlice[E comparable]() ComparableBSlice[E] {
	return &UnsafeComparableBSlice[E]{
		UnsafeAnyBSlice: &UnsafeAnyBSlice[E]{},
	}
}

func NewUnsafeComparableBSliceBySlice[E comparable](s []E) ComparableBSlice[E] {
	return &UnsafeComparableBSlice[E]{
		UnsafeAnyBSlice: &UnsafeAnyBSlice[E]{e: s},
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

func NewSafeComparableBSlice[E comparable]() ComparableBSlice[E] {
	return &SafeComparableBSlice[E]{
		SafeAnyBSlice: &SafeAnyBSlice[E]{},
	}
}

func NewSafeComparableBSliceBySlice[E comparable](s []E) ComparableBSlice[E] {
	return &SafeComparableBSlice[E]{
		SafeAnyBSlice: &SafeAnyBSlice[E]{es: &UnsafeAnyBSlice[E]{e: s}},
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
