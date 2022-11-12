package concurrent

func MapChan[T any](in <-chan T, fn func(T) T) <-chan T {
	out := make(chan T, 1)
	if in == nil {
		close(out)
		return out
	}
	go func() {
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

func ReduceChan[T any](in <-chan T, fn func(r, v T) T) T {
	if in == nil {
		var zero T
		return zero
	}
	out := <-in
	for v := range in {
		out = fn(out, v)
	}
	return out
}
