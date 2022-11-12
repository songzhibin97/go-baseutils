package lists

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/containers"
)

// List interface that all lists implement
type List[E any] interface {
	Get(index int) (E, bool)
	Remove(index int)
	Add(values ...E)
	Contains(values ...E) bool
	Sort(comparator bcomparator.Comparator[E])
	Swap(index1, index2 int)
	Insert(index int, values ...E)
	Set(index int, value E)

	containers.Container[E]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []E
	// String() string
}
