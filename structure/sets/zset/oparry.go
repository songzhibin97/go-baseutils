package zset

import (
	"unsafe"
)

const (
	op1 = 4
	op2 = maxLevel - op1 // TODO: not sure that whether 4 is the best number for op1([28]Pointer for op2).
)

type listLevel struct {
	next unsafe.Pointer // the forward pointer
	span int            // span is count of level 0 element to next element in current level
}

type optionalArray struct {
	base  [op1]listLevel
	extra *([op2]listLevel)
}

func (a *optionalArray) init(level int) {
	if level > op1 {
		a.extra = new([op2]listLevel)
	}
}

func (a *optionalArray) loadNext(i int) unsafe.Pointer {
	if i < op1 {
		return a.base[i].next
	}
	return a.extra[i-op1].next
}

func (a *optionalArray) storeNext(i int, p unsafe.Pointer) {
	if i < op1 {
		a.base[i].next = p
		return
	}
	a.extra[i-op1].next = p
}

func (a *optionalArray) loadSpan(i int) int {
	if i < op1 {
		return a.base[i].span
	}
	return a.extra[i-op1].span
}

func (a *optionalArray) storeSpan(i int, s int) {
	if i < op1 {
		a.base[i].span = s
		return
	}
	a.extra[i-op1].span = s
}
