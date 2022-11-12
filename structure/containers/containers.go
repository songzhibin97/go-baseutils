package containers

import "github.com/songzhibin97/go-baseutils/base/bcomparator"

// Container is base interface that all data structures implement.
type Container[E any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []E
	String() string
}

// GetSortedValues returns sorted container's elements with respect to the passed comparator.
// Does not affect the ordering of elements within the container.
func GetSortedValues[E any](container Container[E], comparator bcomparator.Comparator[E]) []E {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	bcomparator.Sort(values, comparator)
	return values
}
