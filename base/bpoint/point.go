package bpoint

import "github.com/songzhibin97/go-baseutils/base/breflect"

func ToPoint[T any](v T) *T {
	if breflect.IsNil(v) {
		return nil
	}
	return &v
}

func FromPoint[T any](v *T) T {
	var zero T
	return FromPointOrDefaultIfNil(v, zero)
}

func FromPointOrDefaultIfNil[T any](v *T, defaultValue T) T {
	if v == nil || breflect.IsNil(v) {
		return defaultValue
	}
	return *v
}
