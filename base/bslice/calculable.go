package bslice

import (
	"github.com/songzhibin97/go-baseutils/base/bmath"
	"github.com/songzhibin97/go-baseutils/base/bternaryexpr"
	"github.com/songzhibin97/go-baseutils/base/btype"
)

// =====================================================================================================================
// unsafe

func NewUnsafeCalculableBSlice[E btype.Integer | btype.Float]() CalculableBSlice[E] {
	return &UnsafeCalculableBSlice[E]{
		UnsafeOrderedBSlice: &UnsafeOrderedBSlice[E]{
			UnsafeComparableBSlice: &UnsafeComparableBSlice[E]{
				UnsafeAnyBSlice: &UnsafeAnyBSlice[E]{
					e: nil,
				},
			},
		},
	}
}

func NewUnsafeCalculableBSliceBySlice[E btype.Integer | btype.Float](s []E) CalculableBSlice[E] {
	return &UnsafeCalculableBSlice[E]{
		UnsafeOrderedBSlice: &UnsafeOrderedBSlice[E]{
			UnsafeComparableBSlice: &UnsafeComparableBSlice[E]{
				UnsafeAnyBSlice: &UnsafeAnyBSlice[E]{
					e: s,
				},
			},
		},
	}
}

type UnsafeCalculableBSlice[E btype.Integer | btype.Float] struct {
	*UnsafeOrderedBSlice[E]
}

func (x *UnsafeCalculableBSlice[E]) Sum() E {
	var r E
	for _, e := range x.ToMetaSlice() {
		r += e
	}
	return r
}

func (x *UnsafeCalculableBSlice[E]) Avg() E {
	var r E
	list := x.ToMetaSlice()
	ln := len(list)
	for _, e := range list {
		r += e
	}
	return x.Sum() / bternaryexpr.TernaryExpr(ln == 0, 1, E(ln))
}

func (x *UnsafeCalculableBSlice[E]) Max() E {
	var r E
	list := x.ToMetaSlice()
	ln := len(list)
	if ln != 0 {
		r = list[0]
	}
	for i := 1; i < ln; i++ {
		r = bmath.Max(r, list[i])
	}
	return r
}

func (x *UnsafeCalculableBSlice[E]) Min() E {
	var r E
	list := x.ToMetaSlice()
	ln := len(list)
	if ln != 0 {
		r = list[0]
	}
	for i := 1; i < ln; i++ {
		r = bmath.Min(r, list[i])
	}
	return r
}

// =====================================================================================================================
// safe

func NewSafeCalculableBSlice[E btype.Integer | btype.Float]() CalculableBSlice[E] {
	return &SafeCalculableBSlice[E]{
		SafeOrderedBSlice: &SafeOrderedBSlice[E]{
			SafeComparableBSlice: &SafeComparableBSlice[E]{
				SafeAnyBSlice: &SafeAnyBSlice[E]{
					es: &UnsafeAnyBSlice[E]{
						e: []E{},
					},
				},
			},
		},
	}
}

func NewSafeCalculableBSliceBySlice[E btype.Integer | btype.Float](s []E) CalculableBSlice[E] {
	return &SafeCalculableBSlice[E]{
		SafeOrderedBSlice: &SafeOrderedBSlice[E]{
			SafeComparableBSlice: &SafeComparableBSlice[E]{
				SafeAnyBSlice: &SafeAnyBSlice[E]{
					es: &UnsafeAnyBSlice[E]{
						e: s,
					},
				},
			},
		},
	}
}

type SafeCalculableBSlice[E btype.Integer | btype.Float] struct {
	*SafeOrderedBSlice[E]
}

func (x *SafeCalculableBSlice[E]) Sum() E {
	list := x.ToMetaSlice()
	var r E
	for _, e := range list {
		r += e
	}
	return r
}

func (x *SafeCalculableBSlice[E]) Avg() E {
	var r E
	list := x.ToMetaSlice()
	ln := len(list)
	for _, e := range list {
		r += e
	}
	return x.Sum() / bternaryexpr.TernaryExpr(ln == 0, 1, E(ln))
}

func (x *SafeCalculableBSlice[E]) Max() E {
	var r E
	list := x.ToMetaSlice()
	ln := len(list)
	if ln != 0 {
		r = list[0]
	}
	for i := 1; i < ln; i++ {
		r = bmath.Max(r, list[i])
	}
	return r
}

func (x *SafeCalculableBSlice[E]) Min() E {
	var r E
	list := x.ToMetaSlice()
	ln := len(list)
	if ln != 0 {
		r = list[0]
	}
	for i := 1; i < ln; i++ {
		r = bmath.Min(r, list[i])
	}
	return r
}
