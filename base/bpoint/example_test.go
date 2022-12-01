package bpoint

import "fmt"

func ExampleToPoint() {
	fmt.Println(*ToPoint(1)) // 1
}

func ExampleFromPoint() {
	v := 1
	fmt.Println(FromPoint(&v)) // 1
}

func ExampleFromPointOrDefaultIfNil() {
	var v *int
	fmt.Println(FromPointOrDefaultIfNil(v, 2)) // 2
}
