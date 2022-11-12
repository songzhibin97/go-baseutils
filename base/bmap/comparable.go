package bmap

// =====================================================================================================================
// unsafe

func NewUnsafeComparableBMap[K comparable, V comparable]() ComparableBMap[K, V] {
	return &UnsafeComparableBMap[K, V]{
		UnsafeAnyBMap: &UnsafeAnyBMap[K, V]{
			mp: make(map[K]V),
		},
	}
}

func NewUnsafeComparableBMapByMap[K comparable, V comparable](mp map[K]V) ComparableBMap[K, V] {
	return &UnsafeComparableBMap[K, V]{
		UnsafeAnyBMap: &UnsafeAnyBMap[K, V]{
			mp: mp,
		},
	}
}

type UnsafeComparableBMap[K comparable, V comparable] struct {
	*UnsafeAnyBMap[K, V]
}

func (x *UnsafeComparableBMap[K, V]) EqualByMap(m map[K]V) bool {
	return Equal(x.mp, m)
}

func (x *UnsafeComparableBMap[K, V]) EqualByBMap(m AnyBMap[K, V]) bool {
	return Equal(x.mp, m.ToMetaMap())
}

// =====================================================================================================================
// safe

func NewSafeComparableBMap[K comparable, V comparable]() ComparableBMap[K, V] {
	return &SafeComparableBMap[K, V]{
		SafeAnyBMap: &SafeAnyBMap[K, V]{
			mp: &UnsafeAnyBMap[K, V]{
				mp: map[K]V{},
			},
		},
	}
}

func NewSafeComparableBMapByMap[K comparable, V comparable](mp map[K]V) ComparableBMap[K, V] {
	return &SafeComparableBMap[K, V]{
		SafeAnyBMap: &SafeAnyBMap[K, V]{
			mp: &UnsafeAnyBMap[K, V]{
				mp: mp,
			},
		},
	}
}

type SafeComparableBMap[K comparable, V comparable] struct {
	*SafeAnyBMap[K, V]
}

func (x *SafeComparableBMap[K, V]) EqualByMap(m map[K]V) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return Equal(x.mp.mp, m)
}

func (x *SafeComparableBMap[K, V]) EqualByBMap(m AnyBMap[K, V]) bool {
	x.rwl.RLock()
	defer x.rwl.RUnlock()
	return Equal(x.mp.mp, m.ToMetaMap())
}
