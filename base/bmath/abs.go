package bmath

import "github.com/songzhibin97/go-baseutils/base/btype"

func Abs[T btype.Integer | btype.Float](a T) T {
	if a > 0 {
		return a
	}
	return -a
}
