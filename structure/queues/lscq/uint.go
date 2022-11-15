package lscq

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

var uint64SCQPool = sync.Pool{
	New: func() interface{} {
		return newUint64SCQ()
	},
}

type Uint64Queue struct {
	head *uint64SCQ
	_    [lscqcacheLineSize - unsafe.Sizeof(new(uintptr))]byte
	tail *uint64SCQ
}

func NewUint64() *Uint64Queue {
	q := newUint64SCQ()
	return &Uint64Queue{head: q, tail: q}
}

func (q *Uint64Queue) Dequeue() (data uint64, ok bool) {
	for {
		cq := (*uint64SCQ)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head))))
		data, ok = cq.Dequeue()
		if ok {
			return
		}
		// cq does not have enough entries.
		nex := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&cq.next)))
		if nex == nil {
			// We don't have next SCQ.
			return
		}
		// cq.next is not empty, subsequent entry will be insert into cq.next instead of cq.
		// So if cq is empty, we can move it into ncqpool.
		atomic.StoreInt64(&cq.threshold, int64(scqsize*2)-1)
		data, ok = cq.Dequeue()
		if ok {
			return
		}
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), (unsafe.Pointer(cq)), nex) {
			// We can't ensure no other goroutines will access cq.
			// The cq can still be previous dequeue's cq.
			cq = nil
		}
	}
}

func (q *Uint64Queue) Enqueue(data uint64) bool {
	for {
		cq := (*uint64SCQ)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail))))
		nex := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&cq.next)))
		if nex != nil {
			// Help move cq.next into tail.
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), (unsafe.Pointer(cq)), nex)
			continue
		}
		if cq.Enqueue(data) {
			return true
		}
		// Concurrent cq is full.
		atomicTestAndSetFirstBit(&cq.tail) // close cq, subsequent enqueue will fail
		cq.mu.Lock()
		if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&cq.next))) != nil {
			cq.mu.Unlock()
			continue
		}
		ncq := uint64SCQPool.Get().(*uint64SCQ) // create a new queue
		ncq.Enqueue(data)
		// Try Add this queue into cq.next.
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&cq.next)), nil, unsafe.Pointer(ncq)) {
			// Success.
			// Try move cq.next into tail (we don't need to recheck since other enqueuer will help).
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), unsafe.Pointer(cq), unsafe.Pointer(ncq))
			cq.mu.Unlock()
			return true
		}
		// CAS failed, put this new SCQ into scqpool.
		// No other goroutines will access this queue.
		ncq.Dequeue()
		uint64SCQPool.Put(ncq)
		cq.mu.Unlock()
	}
}

func newUint64SCQ() *uint64SCQ {
	ring := new([scqsize]scqNodeUint64)
	for i := range ring {
		ring[i].flags = 1<<63 + 1<<62 // newSCQFlags(true, true, 0)
	}
	return &uint64SCQ{
		head:      scqsize,
		tail:      scqsize,
		threshold: -1,
		ring:      ring,
	}
}

type uint64SCQ struct {
	_         [lscqcacheLineSize]byte
	head      uint64
	_         [lscqcacheLineSize - unsafe.Sizeof(new(uint64))]byte
	tail      uint64 // 1-bit finalize + 63-bit tail
	_         [lscqcacheLineSize - unsafe.Sizeof(new(uint64))]byte
	threshold int64
	_         [lscqcacheLineSize - unsafe.Sizeof(new(uint64))]byte
	next      *uint64SCQ
	ring      *[scqsize]scqNodeUint64
	mu        sync.Mutex
}

type scqNodeUint64 struct {
	flags uint64 // isSafe 1-bit + isEmpty 1-bit + cycle 62-bit
	data  uint64
}

func (q *uint64SCQ) Enqueue(data uint64) bool {
	for {
		// Increment the TAIL, try to occupy an entry.
		tailvalue := atomic.AddUint64(&q.tail, 1)
		tailvalue -= 1 // we need previous value
		T := uint64Get63(tailvalue)
		if uint64Get1(tailvalue) {
			// The queue is closed, return false, so following enqueuer
			// will insert this data into next SCQ.
			return false
		}
		entAddr := &q.ring[cacheRemap16Byte(T)]
		cycleT := T / scqsize
	eqretry:
		// Enqueue do not need data, if this entry is empty, we can assume the data is also empty.
		entFlags := atomic.LoadUint64(&entAddr.flags)
		isSafe, isEmpty, cycleEnt := loadSCQFlags(entFlags)
		if cycleEnt < cycleT && isEmpty && (isSafe || atomic.LoadUint64(&q.head) <= T) {
			// We can use this entry for adding new data if
			// 1. Tail's cycle is bigger than entry's cycle.
			// 2. It is empty.
			// 3. It is safe or tail >= head (There is enough space for this data)
			ent := scqNodeUint64{flags: entFlags}
			newEnt := scqNodeUint64{flags: newSCQFlags(true, false, cycleT), data: data}
			// Save input data into this entry.
			if !compareAndSwapSCQNodeUint64(entAddr, ent, newEnt) {
				// Failed, do next retry.
				goto eqretry
			}
			// Success.
			if atomic.LoadInt64(&q.threshold) != (int64(scqsize)*2)-1 {
				atomic.StoreInt64(&q.threshold, (int64(scqsize)*2)-1)
			}
			return true
		}
		// Add a full queue check in the loop(CAS2).
		if T+1 >= atomic.LoadUint64(&q.head)+scqsize {
			// T is tail's value before FAA(1), latest tail is T+1.
			return false
		}
	}
}

func (q *uint64SCQ) Dequeue() (data uint64, ok bool) {
	if atomic.LoadInt64(&q.threshold) < 0 {
		// Empty queue.
		return
	}

	for {
		// Decrement HEAD, try to release an entry.
		H := atomic.AddUint64(&q.head, 1)
		H -= 1 // we need previous value
		entAddr := &q.ring[cacheRemap16Byte(H)]
		cycleH := H / scqsize
	dqretry:
		ent := loadSCQNodeUint64(unsafe.Pointer(entAddr))
		isSafe, isEmpty, cycleEnt := loadSCQFlags(ent.flags)
		if cycleEnt == cycleH { // same cycle, return this entry directly
			// 1. Clear the data in this slot.
			// 2. Set `isEmpty` to 1

			resetNode(unsafe.Pointer(entAddr))
			return ent.data, true
		}
		if cycleEnt < cycleH {
			var newEnt scqNodeUint64
			if isEmpty {
				newEnt = scqNodeUint64{flags: newSCQFlags(isSafe, true, cycleH)}
			} else {
				newEnt = scqNodeUint64{flags: newSCQFlags(false, false, cycleEnt), data: ent.data}
			}
			if !compareAndSwapSCQNodeUint64(entAddr, ent, newEnt) {
				goto dqretry
			}
		}
		// Check if the queue is empty.
		tailvalue := atomic.LoadUint64(&q.tail)
		T := uint64Get63(tailvalue)
		if T <= H+1 {
			// Invalid state.
			q.fixstate(H + 1)
			atomic.AddInt64(&q.threshold, -1)
			return
		}
		if atomic.AddInt64(&q.threshold, -1)+1 <= 0 {
			return
		}
	}
}

func (q *uint64SCQ) fixstate(originalHead uint64) {
	for {
		head := atomic.LoadUint64(&q.head)
		if originalHead < head {
			// The last dequeuer will be responsible for fixstate.
			return
		}
		tailvalue := atomic.LoadUint64(&q.tail)
		if tailvalue >= head {
			// The queue has been closed, or in normal state.
			return
		}
		if atomic.CompareAndSwapUint64(&q.tail, tailvalue, head) {
			return
		}
	}
}