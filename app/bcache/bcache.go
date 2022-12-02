package bcache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/base/bmap"
	"github.com/songzhibin97/go-baseutils/base/options"
	"github.com/songzhibin97/go-baseutils/structure/sets/zset"
)

const (
	DefaultExpire time.Duration = 0
	NoExpire      time.Duration = -1
)

var _ Cache[int, any] = (*BCache[int, any])(nil)

type BCache[K comparable, V any] struct {
	*bCache[K, V]
}

func New[K comparable, V any](comparator bcomparator.Comparator[K], opts ...options.Option[*Config[K, V]]) *BCache[K, V] {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Config[K, V]{
		defaultExpire: 0,
		interval:      0,
		capture: func(k K, v V) {
			fmt.Printf("delete k:%v v:%v\n", k, v)
		},
		comparator: comparator,
	}
	for _, option := range opts {
		option(c)
	}
	obj := &bCache[K, V]{
		config:        c,
		defaultExpire: c.defaultExpire,
		capture:       c.capture,
		cancel:        cancel,
	}
	if c.setSentinelFn == nil {
		c.setSentinelFn = obj.deleteExpire
	}
	obj.member = bmap.NewUnsafeAnyBMap[K, Iterator[V]]()
	obj.visit = zset.New[K](comparator)
	go NewSentinel(ctx, c.interval, c.setSentinelFn).Start()
	return &BCache[K, V]{obj}
}

func (c *BCache[K, V]) Set(k K, v V, d time.Duration) {
	c.set(k, v, d)
}

func (c *BCache[K, V]) SetDefault(k K, v V) {
	c.set(k, v, c.defaultExpire)
}

func (c *BCache[K, V]) SetNoExpire(k K, v V) {
	c.set(k, v, NoExpire)
}

func (c *BCache[K, V]) SetIfAbsent(k K, v V, d time.Duration) bool {
	return c.setIfAbsent(k, v, d)
}

func (c *BCache[K, V]) Replace(k K, v V, d time.Duration) bool {
	return c.replace(k, v, d)
}

func (c *BCache[K, V]) Delete(k K) {
	c.bCache.Delete(k)
}

func (c *BCache[K, V]) Get(k K) (V, bool) {
	v, _, ok := c.get(k)
	return v, ok
}

func (c *BCache[K, V]) GetWithExpire(k K) (V, time.Time, bool) {
	v, t, ok := c.get(k)
	return v, t, ok
}

func (c *BCache[K, V]) Count() int {
	return c.count()
}

func (c *BCache[K, V]) Clear() {
	c.clear()
}

func (c *BCache[K, V]) Load(data []byte) error {
	return c.Unmarshal(data)
}

func (c *BCache[K, V]) Export() ([]byte, error) {
	return c.Marshal()
}

type bCache[K comparable, V any] struct {
	config *Config[K, V]

	// defaultExpire 默认超时时间
	defaultExpire time.Duration

	sync.RWMutex // Protection member&visit

	// member 维护map存储kv关系
	member bmap.AnyBMap[K, Iterator[V]]

	// visit 维护具有超时的key
	visit *zset.Set[K]

	// capture 捕获删除对象时间 会返回kv值用于用户自定义处理
	capture func(k K, v V)

	cancel context.CancelFunc

	zero     V
	zeroTime time.Time
}

func (c *bCache[K, V]) newIterator(v V, d time.Duration) Iterator[V] {
	var expire int64
	switch d {
	case NoExpire:
	case DefaultExpire:
		if c.defaultExpire > 0 {
			expire = time.Now().Add(c.defaultExpire).UnixNano()
		}
	default:
		if d > 0 {
			expire = time.Now().Add(d).UnixNano()
		}
		// 如果走到这里 默认是 NoExpire
	}
	return Iterator[V]{
		Value:  v,
		Expire: expire,
	}
}

func (c *bCache[K, V]) set(k K, v V, d time.Duration) {

	iter := c.newIterator(v, d)
	c.Lock()
	defer c.Unlock()

	if iter.Expire != 0 {
		c.visit.AddB(float64(iter.Expire), k)
	}
	c.member.Put(k, iter)
}

func (c *bCache[K, V]) setDeadline(k K, v V, d int64) {

	if d != 0 {
		c.visit.AddB(float64(d), k)
	}
	c.member.Put(k, Iterator[V]{
		Value:  v,
		Expire: d,
	})
}

func (c *bCache[K, V]) setIfAbsent(k K, v V, d time.Duration) bool {
	iter := c.newIterator(v, d)

	c.Lock()
	defer c.Unlock()

	ok := c.member.PuTIfAbsent(k, iter)
	if ok && iter.Expire != 0 {
		c.visit.AddB(float64(iter.Expire), k)
	}
	return ok
}

func (c *bCache[K, V]) get(k K) (V, time.Time, bool) {
	c.Lock()
	defer c.Unlock()
	v, ok := c.member.Get(k)
	if !ok {
		return c.zero, c.zeroTime, false
	}
	if v.expired() {
		c.delete(k)
		return c.zero, c.zeroTime, false
	}
	if !v.isVisit() {
		return v.Value, c.zeroTime, true
	}
	return v.Value, time.Unix(0, v.Expire), true
}

func (c *bCache[K, V]) replace(k K, v V, d time.Duration) bool {
	iter := c.newIterator(v, d)
	c.Lock()
	defer c.Unlock()
	ov, ok := c.member.Get(k)
	if !ok {
		return false
	}
	if ov.expired() {
		c.delete(k)
		return false
	}
	c.member.Put(k, iter)
	if iter.Expire != 0 {
		c.visit.AddB(float64(iter.Expire), k)
	}
	return true
}

func (c *bCache[K, V]) Delete(k K) {
	c.Lock()
	defer c.Unlock()
	v, ok := c.delete(k)
	if ok && c.capture != nil {
		c.capture(k, v)
	}
}

func (c *bCache[K, V]) delete(k K) (V, bool) {
	nv, ok := c.member.DeleteIfPresent(k)
	if !ok {
		return c.zero, false
	}
	c.visit.Remove(k)
	return nv.Value, true
}

func (c *bCache[K, V]) deleteExpire() {
	c.Lock()
	defer c.Unlock()
	nodes := c.visit.RemoveRangeByScore(0, float64(time.Now().UnixNano()))
	for _, n := range nodes {
		c.member.Delete(n.Value)
	}
}

func (c *bCache[K, V]) count() int {
	c.Lock()
	defer c.Unlock()
	return c.member.Size()
}

func (c *bCache[K, V]) clear() {
	c.Lock()
	defer c.Unlock()
	c.member = bmap.NewUnsafeAnyBMap[K, Iterator[V]]()
	c.visit = zset.New[K](c.config.comparator)
}

func (c *bCache[K, V]) Marshal() ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	return c.member.Marshal()
}

func (c *bCache[K, V]) Unmarshal(data []byte) error {
	c.Lock()
	defer c.Unlock()
	mp := bmap.NewUnsafeAnyBMap[K, Iterator[V]]()
	err := mp.Unmarshal(data)
	if err != nil {
		return err
	}
	mp.ForEach(func(k K, v Iterator[V]) {
		if !v.expired() {
			c.setDeadline(k, v.Value, v.Expire)
		}
	})
	return nil
}
