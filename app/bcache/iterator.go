package bcache

import (
	"time"
)

type Iterator[E any] struct {
	// Value 实际存储的对象
	Value E
	// Expire 过期时间
	// 0 不设置过期时间
	Expire int64
}

// expired 判断是否过期,过期返回 true
func (i Iterator[E]) expired(v ...int64) bool {
	if !i.isVisit() {
		return false
	}
	if len(v) != 0 {
		return v[0] > i.Expire
	}
	return time.Now().UnixNano() > i.Expire
}

// IsVisit 根据expire判断是否需要监控
func (i Iterator[E]) isVisit() bool {
	return i.Expire > 0
}
