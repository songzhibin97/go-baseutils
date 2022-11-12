package bcomparator

import (
	"encoding/json"
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bmath"
	"github.com/songzhibin97/go-baseutils/base/bternaryexpr"
	"github.com/songzhibin97/go-baseutils/base/btype"
	"strings"
)

type Comparator[T any] func(a, b T) int

func (c Comparator[T]) Marshal(t T) ([]byte, error) {
	return json.Marshal(t)
}

type temporary[T any] struct {
	Key T `json:"key"`
}

func (c Comparator[T]) Unmarshal(data []byte, v *T) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			var temp temporary[T]
			data = []byte(fmt.Sprintf(`{"key": "%s"}`, string(data)))
			err = json.Unmarshal(data, &temp)
			if err != nil {
				return err
			}
			*v = temp.Key
		default:
			return err
		}
	}
	return nil
}

func OrderedComparator[T btype.Ordered]() Comparator[T] {
	return func(a T, b T) int {
		return bternaryexpr.TernaryExpr[int](a == b, 0, bternaryexpr.TernaryExpr[int](a > b, 1, -1))
	}
}

func ReverseComparator[T any](comparator Comparator[T]) Comparator[T] {
	return func(a T, b T) int {
		return comparator(a, b) * -1
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// StringComparator 比较字符串
func StringComparator() Comparator[string] {
	return strings.Compare
}

// ------------------------------------------------ ---------------------------------------------------------------------

func IntComparator() Comparator[int] {
	return OrderedComparator[int]()
}

func Int8Comparator() Comparator[int8] {
	return OrderedComparator[int8]()
}

func Int16Comparator() Comparator[int16] {
	return OrderedComparator[int16]()
}

func Int32Comparator() Comparator[int32] {
	return OrderedComparator[int32]()
}

func Int64Comparator() Comparator[int64] {
	return OrderedComparator[int64]()
}

// ------------------------------------------------ ---------------------------------------------------------------------

func UintComparator() Comparator[uint] {
	return OrderedComparator[uint]()
}

func Uint8Comparator() Comparator[uint8] {
	return OrderedComparator[uint8]()
}

func Uint16Comparator() Comparator[uint16] {
	return OrderedComparator[uint16]()
}

func Uint32Comparator() Comparator[uint32] {
	return OrderedComparator[uint32]()
}

func Uint64Comparator() Comparator[uint64] {
	return OrderedComparator[uint64]()
}

// ------------------------------------------------ ---------------------------------------------------------------------

func Float64Comparator() Comparator[float64] {
	return func(a float64, b float64) int {
		return bternaryexpr.TernaryExpr[int](bmath.Abs(a-b) < 0.0000001, 0, bternaryexpr.TernaryExpr[int](a > b, 1, -1))
	}
}

func Float32Comparator() Comparator[float32] {
	return func(a float32, b float32) int {
		return bternaryexpr.TernaryExpr[int](bmath.Abs(float64(a-b)) < 0.0000001, 0, bternaryexpr.TernaryExpr[int](a > b, 1, -1))
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// BoolComparator 布尔类型的比较器，默认是让false排在前面
func BoolComparator() Comparator[bool] {
	return func(a bool, b bool) int {
		return bternaryexpr.TernaryExpr[int](a == b, 0, bternaryexpr.TernaryExpr[int](!a && b, -1, 1))
	}
}
