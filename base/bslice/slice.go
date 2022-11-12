package bslice

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/base/btype"
)

type CalculableBSlice[E btype.Integer | btype.Float] interface {
	OrderedBSlice[E]
	Sum() E
	Avg() E
	Max() E
	Min() E
}

type OrderedBSlice[E btype.Ordered] interface {
	ComparableBSlice[E]

	Compare([]E) int

	Sort()
	IsSorted() bool
	BinarySearch(E) (int, bool)
}

type ComparableBSlice[E comparable] interface {
	AnyBSlice[E]

	Contains(E) bool
	Equal([]E) bool
	Compact()
}

type AnyBSlice[E any] interface {
	EqualFunc([]E, func(E, E) bool) bool
	CompareFunc([]E, func(E, E) int) int
	IndexFunc(func(E) bool) int

	Insert(int, ...E)
	InsertE(int, ...E) error

	Delete(int, int)
	DeleteE(int, int) error
	DeleteToSlice(int, int) []E
	DeleteToSliceE(int, int) ([]E, error)
	DeleteToBSlice(int, int) AnyBSlice[E]
	DeleteToBSliceE(int, int) (AnyBSlice[E], error)

	Replace(int, int, ...E)
	ReplaceE(int, int, ...E) error

	CloneToSlice() []E
	CloneToBSlice() AnyBSlice[E]

	CompactFunc(func(E, E) bool)

	Grow(int)
	GrowE(int) error

	Clip()

	ForEach(func(int, E))

	SortFunc(func(i, j E) bool)
	SortFuncToSlice(func(i, j E) bool) []E
	SortFuncToBSlice(func(i, j E) bool) AnyBSlice[E]

	SortComparator(comparator bcomparator.Comparator[E])
	SortComparatorToSlice(comparator bcomparator.Comparator[E]) []E
	SortComparatorToBSlice(comparator bcomparator.Comparator[E]) AnyBSlice[E]

	SortStableFunc(func(i, j E) bool)
	SortStableFuncToSlice(func(i, j E) bool) []E
	SortStableFuncToBSlice(func(i, j E) bool) AnyBSlice[E]

	IsSortedFunc(func(i, j E) bool) bool
	BinarySearchFunc(E, func(E, E) int) (int, bool)

	Filter(func(E) bool)
	FilterToSlice(func(E) bool) []E
	FilterToBSlice(func(E) bool) AnyBSlice[E]

	Reverse()
	ReverseToSlice() []E
	ReverseToBSlice() AnyBSlice[E]

	Marshal() ([]byte, error)
	Unmarshal(data []byte) error

	Len() int
	Cap() int
	ToInterfaceSlice() []interface{}
	ToMetaSlice() []E // no concurrency safe!

	Swap(i, j int)
	Clear()

	Append(...E)
	AppendToSlice(...E) []E
	AppendToBSlice(...E) AnyBSlice[E]

	CopyToSlice() []E
	CopyToBSlice() AnyBSlice[E]

	GetByIndex(int) E
	GetByIndexE(int) (E, error)
	GetByIndexOrDefault(int, E) E

	GetByRange(int, int) []E           // no concurrency safe!
	GetByRangeE(int, int) ([]E, error) // no concurrency safe!

	SetByIndex(int, E)
	SetByIndexE(int, E) error

	SetByRange(int, []E)
	SetByRangeE(int, []E) error
}
