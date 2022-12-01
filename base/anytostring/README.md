# banytostring

将任何输入转化成字符串

## API

- ToStringE

```go
ToStringE(i interface{}) (string, error) 
```

- ToString

```go
ToString(v any) string 
```

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/banytostring"
)

func main() {
	v := "test"
	nv, err := anytostring.ToStringE(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nv) // test
	
	v = "test"
	fmt.Println(anytostring.ToString(v)) // test
}
```