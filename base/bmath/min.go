package bmath

import "github.com/songzhibin97/go-baseutils/base/btype"

func Min[T btype.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}
