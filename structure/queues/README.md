# queues

## API
- Enqueue 入队
- Dequeue 出队
- Peek 查看队头
- Empty 判断是否为空
- Size 获取map长度
- Clear 清空列表
- Values 获取列表中的所有元素
- String 返回列表的字符串表示

## Realize
- arrayqueue
- circularbuffer
- linkedlistqueue
- [lscq](structure/queues/lscq/README.md)(只实现Enqueue、Dequeue,由于泛型问题,不推荐直接使用)
- priorityqueue