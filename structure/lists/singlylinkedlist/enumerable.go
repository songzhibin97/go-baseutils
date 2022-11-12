package singlylinkedlist

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Assert Enumerable implementation
var _ containers.EnumerableWithIndex[any] = (*List[any])(nil)

// Each calls the given function once for each element, passing that element's index and value.
func (list *List[E]) Each(f func(index int, value E)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (list *List[E]) Map(f func(index int, value E) E) *List[E] {
	newList := &List[E]{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (list *List[E]) Select(f func(index int, value E) bool) *List[E] {
	newList := &List[E]{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (list *List[E]) Any(f func(index int, value E) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (list *List[E]) All(f func(index int, value E) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (list *List[E]) Find(f func(index int, value E) bool) (index int, value E) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	var zero E
	return -1, zero
}
