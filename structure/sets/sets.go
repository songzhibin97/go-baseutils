// Package sets provides an abstract Set interface.
//
// In computer science, a set is an abstract data type that can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.
//
// Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package sets

import "github.com/songzhibin97/go-baseutils/structure/containers"

// Set interface that all sets implement
type Set[E any] interface {
	Add(elements ...E)
	Remove(elements ...E)
	Contains(elements ...E) bool
	// Intersection(another *Set) *Set
	// Union(another *Set) *Set
	// Difference(another *Set) *Set

	containers.Container[E]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []E{}
	// String() string
}
