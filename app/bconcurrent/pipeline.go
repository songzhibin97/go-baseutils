package bconcurrent

// Pipeline 串联模式
func Pipeline[T any](in chan T) <-chan T {
	out := make(chan T, 1)
	go func() {
		defer close(out)
		for v := range in {
			out <- v
		}
	}()
	return out
}
