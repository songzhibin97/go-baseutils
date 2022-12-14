package skipmap

import "sync/atomic"

const (
	fullyLinked = 1 << iota
	marked
)

// concurrent-safe bitflag.
type bitflag struct {
	data uint32
}

func (f *bitflag) SetTrue(flags uint32) {
	for {
		old := atomic.LoadUint32(&f.data)
		if old&flags == flags {
			return
		}
		// Flag is 0, need set it to 1.
		n := old | flags
		if atomic.CompareAndSwapUint32(&f.data, old, n) {
			return
		}
		continue
	}
}

func (f *bitflag) SetFalse(flags uint32) {
	for {
		old := atomic.LoadUint32(&f.data)
		check := old & flags
		if check == 0 {
			return
		}
		// Flag is 1, need set it to 0.
		n := old ^ check
		if atomic.CompareAndSwapUint32(&f.data, old, n) {
			return
		}
		continue
	}
}

func (f *bitflag) Get(flag uint32) bool {
	return (atomic.LoadUint32(&f.data) & flag) != 0
}

func (f *bitflag) MGet(check, expect uint32) bool {
	return (atomic.LoadUint32(&f.data) & check) == expect
}
