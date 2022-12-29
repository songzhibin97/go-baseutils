# bobjectstorage

## API

- Set
- GetSafeAssertion
- Get

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bobjectstorage"
)

type mock struct {
}

func (t mock) Test() string {
	return "test"
}

func main() {
	mc := mock{}
	bobjectstorage.Set("key1", "string")
	bobjectstorage.Set("key2", 1)
	bobjectstorage.Set("key3", mc)
	fmt.Println(bobjectstorage.Get[string]("key1")) // "string"
	fmt.Println(bobjectstorage.Get[string]("key2")) // 1
	fmt.Println(bobjectstorage.Get[mock]("key3"))   // mock obj
	fmt.Println(bobjectstorage.Get[*mock]("key4"))  // mock point obj
}

```
