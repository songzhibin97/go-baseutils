# bcomparator

比较器,用于比较两个对象的大小。

## API

传入a,b两个相同类型的对象,返回0,1,-1来判断a,b大小,

1. a > b 返回1
2. a == b 返回0
3. a < b 返回-1

```go
Comparator[T any] func (a, b T) int
```

- OrderedComparator 返回类型[Integer | Float | ~string] 的默认比较器

```go
OrderedComparator[T btype.Ordered]() Comparator[T]
```

- ReverseComparator 传入比较器,将返回值反转

```go
ReverseComparator[T any](comparator Comparator[T]) Comparator[T]
```

- StringComparator 返回类型string的默认比较器

```go
StringComparator() Comparator[string]
```

- IntComparator 返回类型int的默认比较器

```go
IntComparator() Comparator[int]
```

- Int8Comparator 返回类型int8的默认比较器

```go
Int8Comparator() Comparator[int8]
```

- Int16Comparator 返回类型int16的默认比较器

```go
Int16Comparator() Comparator[int16]
```

- Int32Comparator 返回类型int32的默认比较器

```go
Int32Comparator() Comparator[int32]
```

- Int64Comparator 返回类型int64的默认比较器

```go
Int64Comparator() Comparator[int64]
```

- UintComparator 返回类型uint的默认比较器

```go
UintComparator() Comparator[uint]
```

- Uint8Comparator 返回类型uint8的默认比较器

```go
Uint8Comparator() Comparator[uint8]
```

- Uint16Comparator 返回类型uint16的默认比较器

```go
Uint16Comparator() Comparator[uint16]
```

- Uint32Comparator 返回类型uint32的默认比较器

```go
Uint32Comparator() Comparator[uint32]
```

- Uint64Comparator 返回类型uint64的默认比较器

```go
Uint64Comparator() Comparator[uint64]
```

- Float32Comparator 返回类型float32的默认比较器

```go
Float32Comparator() Comparator[float32]
```

- Float64Comparator 返回类型float64的默认比较器

```go
Float64Comparator() Comparator[float64]
```

- BoolComparator 返回类型bool的默认比较器

```go
BoolComparator() Comparator[bool]
```

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
)

func main() {
	fmt.Println(bcomparator.OrderedComparator[int]()(1, 2)) // -1
	fmt.Println(bcomparator.OrderedComparator[int]()(2, 2)) // 0
	fmt.Println(bcomparator.OrderedComparator[int]()(3, 2)) // 1
}
```