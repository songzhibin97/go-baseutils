package bternaryexpr

import "fmt"

func ExampleTernaryExpr() {
	fmt.Println(TernaryExpr(true, 1, 2))  // 1
	fmt.Println(TernaryExpr(false, 1, 2)) // 2
}
