# zset

## Introduction

zset provides a concurrent-safety sorted set, can be used as a local replacement
of [Redis' zset](https://redis.com/ebook/part-2-core-concepts/chapter-3-commands-in-redis/3-5-sorted-sets/).

The main difference to other sets is, every Value of set is associated with a score, that is used to take the sorted set
ordered, from the smallest to the greatest score.

The zset has `O(log(N))` time complexity when doing Add(ZADD) and Remove(ZREM) operations and `O(1)` time complexity
when doing Contains operations.

## Features

- Concurrent safe API
- Values are sorted with score
- Implementation equivalent to redis
- Fast skiplist level randomization

## Comparison

| Redis command         | Go function         |
|-----------------------|---------------------|
| ZADD                  | Add                 |
| ZINCRBY               | IncrBy              |
| ZREM                  | Remove              |
| ZREMRANGEBYSCORE      | RemoveRangeByScore  |
| ZREMRANGEBYRANK       | RemoveRangeByRank   |
| ZUNION                | Union               |
| ZINTER                | Inter               |
| ZINTERCARD            | *TODO*              |
| ZDIFF                 | *TODO*              |
| ZRANGE                | Range               |
| ZRANGEBYSCORE         | IncrBy              |
| ZREVRANGEBYSCORE      | RevRangeByScore     |
| ZCOUNT                | Count               |
| ZREVRANGE             | RevRange            |
| ZCARD                 | Len                 |
| ZSCORE                | Score               |
| ZRANK                 | Rank                |
| ZREVRANK              | RevRank             |
| ZPOPMIN               | *TODO*              |
| ZPOPMAX               | *TODO*              |
| ZRANDMEMBER           | *TODO*              |

List of redis commands are generated from the following command:

```bash
cat redis/src/server.c | grep -o '"z.*",z.*Command' | grep -o '".*"' | cut -d '"' -f2
```

You may find that not all redis commands have corresponding go implementations,
the reason is as follows:

### Unsupported Commands

Redis' zset can operates elements in lexicographic order, which is not commonly
used function, so zset does not support commands like ZREMRANGEBYLEX, ZLEXCOUNT
and so on.

| Redis command         |
|-----------------------|
| ZREMRANGEBYLEX        |
| ZRANGEBYLEX           |
| ZREVRANGEBYLEX        |
| ZLEXCOUNT             |

In redis, user accesses zset via a string key. We do not need such string key
because we have variable. so the following commands are not implemented:

| Redis command         |
|-----------------------|
| ZUNIONSTORE           |
| ZINTERSTORE           |
| ZDIFFSTORE            |
| ZRANGESTORE           |
| ZMSCORE               |
| ZSCAN                 |


