# bslice

## API

### AnyBSlice

- EqualFunc 传入一个slice以及比较函数判断两个slice是否相同
- CompareFunc 传入一个slice以及一个比较函数判断传入的slice与anyslice的大小关系
- IndexFunc 传入一个函数,返回第一个满足条件的下标
- Insert 指定index位置后插入若干元素
- InsertE 指定index位置后插入若干元素,如果插入失败返回错误
- Delete 指定start,end位置删除元素 s[i:j]
- DeleteE 指定start,end位置删除元素 s[i:j],如果删除失败返回错误
- DeleteToSlice 拷贝副本,指定start,end位置删除元素 s[i:j], 返回副本
- DeleteToSliceE 拷贝副本,指定start,end位置删除元素 s[i:j], 返回副本,如果删除失败返回错误
- DeleteToBSlice 拷贝副本,指定start,end位置删除元素 s[i:j], 返回副本初始化的anyslice
- DeleteToBSliceE 拷贝副本,指定start,end位置删除元素 s[i:j], 返回副本初始化的anyslice,如果删除失败返回错误
- Replace 指定start,end位置替换元素 s[i:j]
- ReplaceE 指定start,end位置替换元素 s[i:j],如果替换失败返回错误
- CloneToSlice 拷贝副本,返回副本
- CloneToBSlice 拷贝副本,返回副本初始化的anyslice
- CompactFunc 传入一个函数,返回true的元素会被删除,收缩slice
- Grow 增长底层数组容量
- GrowE 增长底层数组容量,如果增长失败返回错误
- Clip 删除数组中未使用的长度
- ForEach 传入一个函数,对slice中的每个元素执行函数
- SortFunc 传入一个函数,对slice进行排序
- SortFuncToSlice 拷贝副本,传入一个函数,对slice进行排序,返回副本
- SortFuncToBSlice 拷贝副本,传入一个函数,对slice进行排序,返回副本初始化的anyslice
- SortComparator 传入一个比较器,对slice进行排序
- SortComparatorToSlice 拷贝副本,传入一个比较器,对slice进行排序,返回副本
- SortComparatorToBSlice 拷贝副本,传入一个比较器,对slice进行排序,返回副本初始化的anyslice
- SortStableFunc 传入一个函数,对slice进行稳定排序
- SortStableFuncToSlice 拷贝副本,传入一个函数,对slice进行稳定排序,返回副本
- SortStableFuncToBSlice 拷贝副本,传入一个函数,对slice进行稳定排序,返回副本初始化的anyslice
- IsSortedFunc 传入一个函数,判断slice是否已经排序
- BinarySearchFunc 传入一个目标值以及一个比较函数进行二分查找
- Filter 传入一个函数,将元素传入其中,如果返回true保留,否则删除
- FilterToSlice 拷贝副本,传入一个函数,将元素传入其中,如果返回true保留,否则删除,返回副本
- FilterToBSlice 拷贝副本,传入一个函数,将元素传入其中,如果返回true保留,否则删除,返回副本初始化的anyslice
- Reverse 反转slice
- ReverseToSlice 拷贝副本,反转slice,返回副本
- ReverseToBSlice 拷贝副本,反转slice,返回副本初始化的anyslice
- Marshal
- Unmarshal
- Len 返回slice长度
- Cap 返回slice容量
- ToInterfaceSlice 将slice转换为interface slice
- ToMetaSlice 返回底层原始slice
- Swap 交换两个元素
- Clear 清空slice
- Append 添加元素
- AppendToSlice 拷贝副本,添加元素,返回副本
- AppendToBSlice 拷贝副本,添加元素,返回副本初始化的anyslice
- CopyToSlice 拷贝副本,返回副本
- CopyToBSlice 拷贝副本,返回副本初始化的anyslice
- GetByIndex 根据索引返回元素
- GetByIndexE 根据索引返回元素,如果索引越界返回错误
- GetByIndexOrDefault 根据索引返回元素,如果索引越界返回默认值
- GetByRange 拷贝一个副本,根据索引范围返回副本元素
- GetByRangeE 拷贝一个副本,根据索引范围返回副本,如果索引越界返回错误
- SetByIndex 根据索引设置元素
- SetByIndexE 根据索引设置元素,如果索引越界返回错误
- SetByRange 根据索引范围设置元素
- SetByRangeE 根据索引范围设置元素,如果索引越界返回错误

### ComparableBSlice

- AnyBSlice 继承AnySlice
- Contains 判断是否包含某个元素
- Equal 判断两个slice是否相等
- Compact 相同元素进行收缩

### OrderedBSlice

- ComparableBSlice 继承ComparableSlice
- Compare 比较两个slice
- Sort 对slice进行排序
- IsSorted 判断slice是否已经排序
- BinarySearch 二分查找

### CalculableBSlice

- OrderedBSlice 继承OrderedBSlice
- Sum 求和
- Avg 求平均值
- Max 求最大值
- Min 求最小值

## EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/songzhibin97/go-baseutils/base/bslice"
)

func main() {
	bs := bslice.NewUnsafeAnyBSlice[int]()
	bs.Append(1)
	bs.Append(2)
	fmt.Println(bs.Len()) // 2
	bs.Reverse()          // []int{2,1}
	bs.ForEach(func(index int, v int) {
		fmt.Println(index, v)
	})
}
```