package bcomparator

import "sort"

func Sort[E any](values []E, comparator Comparator[E]) {
	sort.Sort(sortable[E]{values, comparator})
}

type sortable[E any] struct {
	values     []E
	comparator Comparator[E]
}

func (s sortable[E]) Len() int {
	return len(s.values)
}
func (s sortable[E]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortable[E]) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
