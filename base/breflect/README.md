# breflect

## API

- IsNil 判断是否为nil

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/breflect"
)

func main() {
	var a int
	fmt.Println(breflect.IsNil(a)) // false
	var b int
	fmt.Println(breflect.IsNil(&b)) // false
	var c *int
	fmt.Println(breflect.IsNil(c))  // true
	fmt.Println(breflect.IsNil(&c)) // true
}
```