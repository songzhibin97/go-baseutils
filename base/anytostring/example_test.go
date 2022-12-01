package banytostring

import "fmt"

func ExampleToStringE() {
	v := "test"
	nv, err := ToStringE(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nv) // test
}

func ExampleToString() {
	v := "test"
	fmt.Println(ToString(v)) // test
}
