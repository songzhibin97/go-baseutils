package breflect

import "fmt"

func ExampleIsNil() {
	var a int
	fmt.Println(IsNil(a)) // false
	var b int
	fmt.Println(IsNil(&b)) // false
	var c *int
	fmt.Println(IsNil(c))  // true
	fmt.Println(IsNil(&c)) // true
}
