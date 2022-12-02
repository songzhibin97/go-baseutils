# LSCQ

LSCQ is a scalable, unbounded, multiple-producer and multiple-consumer FIFO queue in Go language.

In the benchmark(AMD 3700x, running at 3.6 GHZ, -cpu=16), the LSCQ outperforms lock-based linked queue **5x ~ 6x** in
most cases. Since built-in channel is a bounded queue, we can only compared it in EnqueueDequeuePair, the LSCQ
outperforms built-in channel **8x ~ 9x** in this case.

The ideas behind the LSCQ
are [A Scalable, Portable, and Memory-Efficient Lock-Free FIFO Queue](https://arxiv.org/abs/1908.04511)
and [Fast Concurrent Queues for x86 Processors](https://www.cs.tau.ac.il/~mad/publications/ppopp2013-x86queues.pdf).


