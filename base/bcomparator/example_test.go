package bcomparator

import "fmt"

func ExampleOrderedComparator() {
	fmt.Println(OrderedComparator[int]()(1, 2)) // -1
	fmt.Println(OrderedComparator[int]()(2, 2)) // 0
	fmt.Println(OrderedComparator[int]()(3, 2)) // 1
}
