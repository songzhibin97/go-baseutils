package bmath

import "github.com/songzhibin97/go-baseutils/base/btype"

func Max[T btype.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
