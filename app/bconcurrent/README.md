# bconcurrent

channel一些最佳实践

## API
- FanInRec 扇入模式
- MergeChannel 合并channel
- FanOut 扇出模式
- MapChan 对channel中的元素进行map操作
- ReduceChan 对channel中的元素进行reduce操作
- OrDone 任意channel完成后返回
- Orderly 顺序并发执行
- Pipeline 串联执行
- Stream 流式操作
- TaskN 只取流中的前N个数据
- TaskFn 筛选流中的数据,只保留满足条件的数据
- TaskWhile 只取满足条件的数据,一旦不满足就不再取
- SkipN 跳过流中的前N个数据
- SkipFn 跳过满足条件的数据
- SkipWhile 跳过满足条件的数据,一旦不满足,当前这个元素以后的元素都会输出