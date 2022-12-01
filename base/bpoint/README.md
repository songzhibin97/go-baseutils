# bpoint

## API

- ToPoint 传入v返回*v
- FromPoint 传入*v返回v
- FromPointOrDefaultIfNil 传入*v如果*v为nil返回default value

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bpoint"
)

func main() {
	fmt.Println(*bpoint.ToPoint(1)) // 1

	v := 1
	fmt.Println(bpoint.FromPoint(&v)) // 1

	var nv *int
	fmt.Println(bpoint.FromPointOrDefaultIfNil(nv, 2)) // 2
}

```