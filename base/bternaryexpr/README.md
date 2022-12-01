# ternary

## API

- TernaryExpr

```go
TernaryExpr[T any](boolExpr bool, trueReturn T, falseReturn T) T
```

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bternaryexpr"
)

func main() {
	fmt.Println(bternaryexpr.TernaryExpr(true, 1, 2))  // 1
	fmt.Println(bternaryexpr.TernaryExpr(false, 1, 2)) // 2
}
```
