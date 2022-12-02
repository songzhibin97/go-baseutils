## Introduction

skipset is a high-performance concurrent set based on skip list. In typical pattern(100000 operations, 90%CONTAINS 9%ADD
1%REMOVE), the skipset up to 3x ~ 15x faster than the built-in sync.Map.

The main idea behind the skipset
is [A Simple Optimistic Skiplist Algorithm](<https://people.csail.mit.edu/shanir/publications/LazySkipList.pdf>).

Different from the sync.Map, the items in the skipset are always sorted, and the `Contains` and `Range` operations are
wait-free (A goroutine is guaranteed to complete an operation as long as it keeps taking steps, regardless of the
activity of other goroutines).

## Features

- Concurrent safe API with high-performance.
- Wait-free Contains and Range operations.
- Sorted items.

## When should you use skipset

In these situations, `skipset` is better

- **Sorted elements is needed**.
- **Concurrent calls multiple operations**. such as use both `Contains` and `Add` at the same time.
- **Memory intensive**. The skipset save at least 50% memory in the benchmark.

In these situations, `sync.Map` is better

- Only one goroutine access the set for most of the time, such as insert a batch of elements and then use
  only `Contains` (use built-in map is even better).
