package bcache

import (
	"time"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/base/bmap"
	"github.com/songzhibin97/go-baseutils/base/options"
)

type Config[K comparable, V any] struct {
	// defaultExpire 默认超时时间
	defaultExpire time.Duration

	// interval 间隔时间
	interval time.Duration
	// fn 哨兵周期执行的函数
	setSentinelFn func()

	// capture 捕获删除对象时间 会返回kv值用于用户自定义处理
	capture func(k K, v V)

	member bmap.AnyBMap[K, V]

	comparator bcomparator.Comparator[K]
}

// SetInternal 设置间隔时间
func SetInternal[K comparable, V any](interval time.Duration) options.Option[*Config[K, V]] {
	return func(c *Config[K, V]) {
		c.interval = interval
	}
}

// SetDefaultExpire 设置默认的超时时间
func SetDefaultExpire[K comparable, V any](expire time.Duration) options.Option[*Config[K, V]] {
	return func(c *Config[K, V]) {
		c.defaultExpire = expire
	}
}

// SetSentinelFn 设置周期的执行函数,默认(不设置)是扫描全局清除过期的k
func SetSentinelFn[K comparable, V any](fn func()) options.Option[*Config[K, V]] {
	return func(c *Config[K, V]) {
		c.setSentinelFn = fn
	}
}

// SetCapture 设置触发删除后的捕获函数, 数据删除后回调用设置的捕获函数
func SetCapture[K comparable, V any](capture func(k K, v V)) options.Option[*Config[K, V]] {
	return func(c *Config[K, V]) {
		c.capture = capture
	}
}
