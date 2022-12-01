package bmath

import "fmt"

func ExampleAbs() {
	fmt.Println(Abs[int](1))  // 1
	fmt.Println(Abs[int](0))  // 0
	fmt.Println(Abs[int](-1)) // 1
}

func ExampleMax() {
	fmt.Println(Max[int](1, 2)) // 2
}

func ExampleMin() {
	fmt.Println(Min[int](1, 2)) // 1
}
