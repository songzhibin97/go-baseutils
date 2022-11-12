package bmap

type ComparableBMap[K comparable, V comparable] interface {
	AnyBMap[K, V]

	EqualByMap(m map[K]V) bool
	EqualByBMap(m AnyBMap[K, V]) bool
}

type AnyBMap[K comparable, V any] interface {
	ToMetaMap() map[K]V // no concurrency safe!

	Keys() []K
	Values() []V
	EqualFuncByMap(m map[K]V, eq func(V1, V2 V) bool) bool
	EqualFuncByBMap(m AnyBMap[K, V], eq func(V1, V2 V) bool) bool
	Clear()
	CloneToMap() map[K]V
	CloneToBMap() AnyBMap[K, V]
	CopyByMap(dst map[K]V)
	CopyByBMap(dst AnyBMap[K, V])
	DeleteFunc(del func(K, V) bool)

	Marshal() ([]byte, error)
	Unmarshal(data []byte) error

	Size() int
	IsEmpty() bool
	IsExist(K) bool // k有对应value 返回true
	ContainsKey(K) bool
	ContainsValue(V) bool
	ForEach(func(K, V))
	Get(K) (V, bool)                    // 获取 k对应的value, 如果不存在返回false
	GetOrDefault(k K, defaultValue V) V // 获取 k对应的value, 如果不存在返回defaultValue
	Put(K, V)
	PuTIfAbsent(K, V) bool                      // 设置如果k对应value不存在的情况下设置成功
	Delete(K)                                   // 删除 k对应的value, 如果存在返回旧值
	DeleteIfPresent(K) (V, bool)                // 删除 k对应的value, 如果存在返回true
	MergeByMap(map[K]V, func(K, V) bool)        // 根据返回进行合并 func(k, ov) 传入key和当前nmap的对应key的value值进行冲突处理 return true 进行替换 false则跳过
	MergeByBMap(AnyBMap[K, V], func(K, V) bool) // 根据返回进行合并 func(k, ov) 传入key和当前nmap的对应key的value值进行冲突处理 return true 进行替换 false则跳过
	Replace(k K, ov, nv V) bool                 // 替换 k对应的value等于ov则设置为nv
}
