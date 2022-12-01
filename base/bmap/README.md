# bmap

针对map的方法封装

## API

### AnyMap
- ToMetaMap 获取底层最原始map 无论是否是衍生的safe map都不是并发安全的
- Keys 返回map中所有的key
- Values 返回map中所有的value
- EqualFuncByMap 传入一个原始map以及一个比较函数判断anymap与原始map是否相同
- EqualFuncByBMap 传入一个anymap以及一个比较函数判断anymap与传入的anymap是否相同
- Clear 清除map的kv
- CloneToMap 克隆一个完全相同的原始map
- CloneToBMap 克隆一个完全相同的anymap
- CopyByMap 传入一个原始map将原始map的kv复制到anymap中
- CopyByBMap 传入一个anymap将anymap的kv复制到anymap中
- DeleteFunc 传入一个删除函数删除map中符合条件的kv
- Marshal 
- Unmarshal
- Size 返回键值对数量
- IsEmpty 判断是否为空
- IsExist 传入k判断是否存在
- ContainsKey 传入k判断是否存在
- ContainsValue 传入v判断是否存在
- ForEach 传入一个函数遍历map,最好不要在遍历过程中操作map
- Get 获取 k对应的value, 如果不存在返回false
- GetOrDefault  获取 k对应的value, 如果不存在返回defaultValue
- Put 设置 k对应的value,如果有会覆盖
- PuTIfAbsent 设置如果k对应value不存在的情况下设置成功
- Delete 删除k对应的value
- DeleteIfPresent 删除 k对应的value, 如果存在返回true
- MergeByMap 传入普通map根据返回进行合并 func(k, ov) 传入key和当前nmap的对应key的value值进行冲突处理 return true 进行替换 false则跳过
- MergeByBMap 传入anymap根据返回进行合并 func(k, ov) 传入key和当前nmap的对应key的value值进行冲突处理 return true 进行替换 false则跳过
- Replace 替换 k对应的value等于ov则设置为nv

### ComparableBMap
- AnyBMap[K, V] 继承anymap所有api
- EqualByMap 传入一个原始map以及一个比较函数判断anymap与原始map是否相同
- EqualByBMap 传入一个anymap以及一个比较函数判断anymap与传入的anymap是否相同

## EXAMPLE  
```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bmap"
)

func main() {
	amap := bmap.NewUnsafeAnyBMap[int, int]()
	amap.Put(1, 1)
	fmt.Println(amap.Get(1))             // 1, true
	fmt.Println(amap.Get(2))             // 0, false
	fmt.Println(amap.GetOrDefault(2, 2)) // 2
}
```