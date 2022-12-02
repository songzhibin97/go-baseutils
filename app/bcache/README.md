# bcache

缓存组件

## API
- Set 设置缓存,带失效时间
- SetDefault 设置缓存,使用默认的缓存时间
- SetNoExpire 设置缓存,不过期
- SetIfAbsent 设置缓存,如果不存在设置成功返回bool
- Replace 替换缓存,如果存在设置成功返回bool
- Delete 删除缓存
- Get 获取缓存,返回对应V以及bool
- GetWithExpire 获取缓存,返回对应V以及bool以及过期时间
- Count 获取缓存数量
- Clear 清空缓存
- Load 从文件加载对象
- Export 导出到文件
- Marshal 
- Unmarshal

## EXAMPLE
```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/app/bcache"
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"time"
)

func main() {
	c := bcache.New[int, int](bcomparator.IntComparator())
	c.Set(1, 1, 5*time.Second)
	fmt.Println(c.Get(1)) // 1,true
	time.Sleep(5 * time.Second)
	fmt.Println(c.Get(1)) // 0,false
}

```