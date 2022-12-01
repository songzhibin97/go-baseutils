# bmath

## API

- Max
- Min
- Abs

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bmath"
)

func main() {
	fmt.Println(bmath.Abs[int](1))  // 1
	fmt.Println(bmath.Abs[int](0))  // 0
	fmt.Println(bmath.Abs[int](-1)) // 1

	fmt.Println(bmath.Max[int](1, 2)) // 2

	fmt.Println(bmath.Min[int](1, 2)) // 1
}
```
