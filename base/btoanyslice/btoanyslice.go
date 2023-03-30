package btoanyslice

func ToAnySlice[V any](s []V) []interface{} {
	r := make([]interface{}, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}
