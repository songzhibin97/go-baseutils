// Package skipmap is a high-performance, scalable, concurrent-safe map based on skip-list.
// In the typical pattern(100000 operations, 90%LOAD 9%STORE 1%DELETE, 8C16T), the skipmap
// up to 10x faster than the built-in sync.Map.
package skipmap

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/maps"
)

// Assert Map implementation
var _ maps.Map[int, any] = (*Map[int, any])(nil)

// Map represents a map based on skip list in ascending order.
type Map[K, V any] struct {
	header       *node[K, V]
	length       int64
	highestLevel int64 // highest level for now
	comparator   bcomparator.Comparator[K]
	zeroV        V
}

func (s *Map[K, V]) Put(key K, value V) {
	s.Store(key, value)
}

func (s *Map[K, V]) Get(key K) (value V, found bool) {
	return s.Load(key)
}

func (s *Map[K, V]) Remove(key K) {
	s.Delete(key)
}

func (s *Map[K, V]) Keys() []K {
	ks := make([]K, 0, s.Len())
	s.Range(func(k K, v V) bool {
		ks = append(ks, k)
		return true
	})
	return ks
}

func (s *Map[K, V]) Empty() bool {
	return s.Len() == 0
}

func (s *Map[K, V]) Size() int {
	return s.Len()
}

func (s *Map[K, V]) Clear() {
	var zerok K
	var zerov V
	h := newNode[K, V](zerok, zerov, maxLevel, s.comparator)
	h.flags.SetTrue(fullyLinked)
	s.header = h
	s.highestLevel = defaultHighestLevel
}

func (s *Map[K, V]) Values() []V {
	vs := make([]V, 0, s.Len())
	s.Range(func(k K, v V) bool {
		vs = append(vs, v)
		return true
	})
	return vs
}

func (s *Map[K, V]) String() string {
	bf := strings.Builder{}
	bf.WriteString("SkipMap\nmap[")
	s.Range(func(k K, v V) bool {
		bf.WriteString(fmt.Sprintf("(%v:%v) ", k, v))
		return true
	})
	bf.WriteString("]")
	return bf.String()
}

type node[K, V any] struct {
	key        K
	value      unsafe.Pointer // unsafe.Pointer *V{}
	next       optionalArray  // [level]*node
	mu         sync.Mutex
	flags      bitflag
	level      uint32
	comparator bcomparator.Comparator[K]
}

func newNode[K, V any](key K, value V, level int, comparator bcomparator.Comparator[K]) *node[K, V] {
	node := &node[K, V]{
		key:        key,
		level:      uint32(level),
		comparator: comparator,
	}
	node.storeVal(value)
	if level > op1 {
		node.next.extra = new([op2]unsafe.Pointer)
	}
	return node
}

func (n *node[K, V]) storeVal(value V) {
	atomic.StorePointer(&n.value, unsafe.Pointer(&value))
}

func (n *node[K, V]) loadVal() V {
	return *(*V)(atomic.LoadPointer(&n.value))
}

func (n *node[K, V]) loadNext(i int) *node[K, V] {
	return (*node[K, V])(n.next.load(i))
}

func (n *node[K, V]) storeNext(i int, node *node[K, V]) {
	n.next.store(i, unsafe.Pointer(node))
}

func (n *node[K, V]) atomicLoadNext(i int) *node[K, V] {
	return (*node[K, V])(n.next.atomicLoad(i))
}

func (n *node[K, V]) atomicStoreNext(i int, node *node[K, V]) {
	n.next.atomicStore(i, unsafe.Pointer(node))
}

func (n *node[K, V]) lessthan(key K, comparator bcomparator.Comparator[K]) bool {
	return comparator(n.key, key) < 0
}

func (n *node[K, V]) equal(key K, comparator bcomparator.Comparator[K]) bool {
	return comparator(n.key, key) == 0
}

// New return an empty int64 skipmap.
func New[K, V any](comparator bcomparator.Comparator[K]) *Map[K, V] {
	var zerok K
	var zerov V
	h := newNode[K, V](zerok, zerov, maxLevel, comparator)
	h.flags.SetTrue(fullyLinked)
	return &Map[K, V]{
		header:       h,
		highestLevel: defaultHighestLevel,
		comparator:   comparator,
	}
}

// findNode takes a key and two maximal-height arrays then searches exactly as in a sequential skipmap.
// The returned preds and succs always satisfy preds[i] > key >= succs[i].
// (without fullpath, if find the node will return immediately)
func (s *Map[K, V]) findNode(key K, preds *[maxLevel]*node[K, V], succs *[maxLevel]*node[K, V]) *node[K, V] {
	x := s.header
	for i := int(atomic.LoadInt64(&s.highestLevel)) - 1; i >= 0; i-- {
		succ := x.atomicLoadNext(i)
		for succ != nil && succ.lessthan(key, s.comparator) {
			x = succ
			succ = x.atomicLoadNext(i)
		}
		preds[i] = x
		succs[i] = succ

		// Check if the key already in the skipmap.
		if succ != nil && succ.equal(key, s.comparator) {
			return succ
		}
	}
	return nil
}

// findNodeDelete takes a key and two maximal-height arrays then searches exactly as in a sequential skip-list.
// The returned preds and succs always satisfy preds[i] > key >= succs[i].
func (s *Map[K, V]) findNodeDelete(key K, preds *[maxLevel]*node[K, V], succs *[maxLevel]*node[K, V]) int {
	// lFound represents the index of the first layer at which it found a node.
	lFound, x := -1, s.header
	for i := int(atomic.LoadInt64(&s.highestLevel)) - 1; i >= 0; i-- {
		succ := x.atomicLoadNext(i)
		for succ != nil && succ.lessthan(key, s.comparator) {
			x = succ
			succ = x.atomicLoadNext(i)
		}
		preds[i] = x
		succs[i] = succ

		// Check if the key already in the skip list.
		if lFound == -1 && succ != nil && succ.equal(key, s.comparator) {
			lFound = i
		}
	}
	return lFound
}

func unlock[K, V any](preds [maxLevel]*node[K, V], highestLevel int) {
	var prevPred *node[K, V]
	for i := highestLevel; i >= 0; i-- {
		if preds[i] != prevPred { // the node could be unlocked by previous loop
			preds[i].mu.Unlock()
			prevPred = preds[i]
		}
	}
}

// Store sets the value for a key.
func (s *Map[K, V]) Store(key K, value V) {
	level := s.randomlevel()
	var preds, succs [maxLevel]*node[K, V]
	for {
		nodeFound := s.findNode(key, &preds, &succs)
		if nodeFound != nil { // indicating the key is already in the skip-list
			if !nodeFound.flags.Get(marked) {
				// We don't need to care about whether or not the node is fully linked,
				// just replace the value.
				nodeFound.storeVal(value)
				return
			}
			// If the node is marked, represents some other goroutines is in the process of deleting this node,
			// we need to add this node in next loop.
			continue
		}

		// Add this node into skip list.
		var (
			highestLocked        = -1 // the highest level being locked by this process
			valid                = true
			pred, succ, prevPred *node[K, V]
		)
		for layer := 0; valid && layer < level; layer++ {
			pred = preds[layer]   // target node's previous node
			succ = succs[layer]   // target node's next node
			if pred != prevPred { // the node in this layer could be locked by previous loop
				pred.mu.Lock()
				highestLocked = layer
				prevPred = pred
			}
			// valid check if there is another node has inserted into the skip list in this layer during this process.
			// It is valid if:
			// 1. The previous node and next node both are not marked.
			// 2. The previous node's next node is succ in this layer.
			valid = !pred.flags.Get(marked) && (succ == nil || !succ.flags.Get(marked)) && pred.loadNext(layer) == succ
		}
		if !valid {
			unlock(preds, highestLocked)
			continue
		}

		nn := newNode[K, V](key, value, level, s.comparator)
		for layer := 0; layer < level; layer++ {
			nn.storeNext(layer, succs[layer])
			preds[layer].atomicStoreNext(layer, nn)
		}
		nn.flags.SetTrue(fullyLinked)
		unlock(preds, highestLocked)
		atomic.AddInt64(&s.length, 1)
		return
	}
}

func (s *Map[K, V]) randomlevel() int {
	// Generate random level.
	level := randomLevel()
	// Update highest level if possible.
	for {
		hl := atomic.LoadInt64(&s.highestLevel)
		if int64(level) <= hl {
			break
		}
		if atomic.CompareAndSwapInt64(&s.highestLevel, hl, int64(level)) {
			break
		}
	}
	return level
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (s *Map[K, V]) Load(key K) (value V, ok bool) {
	x := s.header
	for i := int(atomic.LoadInt64(&s.highestLevel)) - 1; i >= 0; i-- {
		nex := x.atomicLoadNext(i)
		for nex != nil && nex.lessthan(key, s.comparator) {
			x = nex
			nex = x.atomicLoadNext(i)
		}

		// Check if the key already in the skip list.
		if nex != nil && nex.equal(key, s.comparator) {
			if nex.flags.MGet(fullyLinked|marked, fullyLinked) {
				return nex.loadVal(), true
			}
			return s.zeroV, false
		}
	}
	return s.zeroV, false
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
// (Modified from Delete)
func (s *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	var (
		nodeToDelete *node[K, V]
		isMarked     bool // represents if this operation mark the node
		topLayer     = -1
		preds, succs [maxLevel]*node[K, V]
	)
	for {
		lFound := s.findNodeDelete(key, &preds, &succs)
		if isMarked || // this process mark this node or we can find this node in the skip list
			lFound != -1 && succs[lFound].flags.MGet(fullyLinked|marked, fullyLinked) && (int(succs[lFound].level)-1) == lFound {
			if !isMarked { // we don't mark this node for now
				nodeToDelete = succs[lFound]
				topLayer = lFound
				nodeToDelete.mu.Lock()
				if nodeToDelete.flags.Get(marked) {
					// The node is marked by another process,
					// the physical deletion will be accomplished by another process.
					nodeToDelete.mu.Unlock()
					return s.zeroV, false
				}
				nodeToDelete.flags.SetTrue(marked)
				isMarked = true
			}
			// Accomplish the physical deletion.
			var (
				highestLocked        = -1 // the highest level being locked by this process
				valid                = true
				pred, succ, prevPred *node[K, V]
			)
			for layer := 0; valid && (layer <= topLayer); layer++ {
				pred, succ = preds[layer], succs[layer]
				if pred != prevPred { // the node in this layer could be locked by previous loop
					pred.mu.Lock()
					highestLocked = layer
					prevPred = pred
				}
				// valid check if there is another node has inserted into the skip list in this layer
				// during this process, or the previous is deleted by another process.
				// It is valid if:
				// 1. the previous node exists.
				// 2. no another node has inserted into the skip list in this layer.
				valid = !pred.flags.Get(marked) && pred.loadNext(layer) == succ
			}
			if !valid {
				unlock(preds, highestLocked)
				continue
			}
			for i := topLayer; i >= 0; i-- {
				// Now we own the `nodeToDelete`, no other goroutine will modify it.
				// So we don't need `nodeToDelete.loadNext`
				preds[i].atomicStoreNext(i, nodeToDelete.loadNext(i))
			}
			nodeToDelete.mu.Unlock()
			unlock(preds, highestLocked)
			atomic.AddInt64(&s.length, -1)
			return nodeToDelete.loadVal(), true
		}
		return s.zeroV, false
	}
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
// (Modified from Store)
func (s *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	var (
		level        int
		preds, succs [maxLevel]*node[K, V]
		hl           = int(atomic.LoadInt64(&s.highestLevel))
	)
	for {
		nodeFound := s.findNode(key, &preds, &succs)
		if nodeFound != nil { // indicating the key is already in the skip-list
			if !nodeFound.flags.Get(marked) {
				// We don't need to care about whether or not the node is fully linked,
				// just return the value.
				return nodeFound.loadVal(), true
			}
			// If the node is marked, represents some other goroutines is in the process of deleting this node,
			// we need to add this node in next loop.
			continue
		}

		// Add this node into skip list.
		var (
			highestLocked        = -1 // the highest level being locked by this process
			valid                = true
			pred, succ, prevPred *node[K, V]
		)
		if level == 0 {
			level = s.randomlevel()
			if level > hl {
				// If the highest level is updated, usually means that many goroutines
				// are inserting items. Hopefully we can find a better path in next loop.
				// TODO(zyh): consider filling the preds if s.header[level].next == nil,
				// but this strategy's performance is almost the same as the existing method.
				continue
			}
		}
		for layer := 0; valid && layer < level; layer++ {
			pred = preds[layer]   // target node's previous node
			succ = succs[layer]   // target node's next node
			if pred != prevPred { // the node in this layer could be locked by previous loop
				pred.mu.Lock()
				highestLocked = layer
				prevPred = pred
			}
			// valid check if there is another node has inserted into the skip list in this layer during this process.
			// It is valid if:
			// 1. The previous node and next node both are not marked.
			// 2. The previous node's next node is succ in this layer.
			valid = !pred.flags.Get(marked) && (succ == nil || !succ.flags.Get(marked)) && pred.loadNext(layer) == succ
		}
		if !valid {
			unlock(preds, highestLocked)
			continue
		}

		nn := newNode(key, value, level, s.comparator)
		for layer := 0; layer < level; layer++ {
			nn.storeNext(layer, succs[layer])
			preds[layer].atomicStoreNext(layer, nn)
		}
		nn.flags.SetTrue(fullyLinked)
		unlock(preds, highestLocked)
		atomic.AddInt64(&s.length, 1)
		return value, false
	}
}

// LoadOrStoreLazy returns the existing value for the key if present.
// Otherwise, it stores and returns the given value from f, f will only be called once.
// The loaded result is true if the value was loaded, false if stored.
// (Modified from LoadOrStore)
func (s *Map[K, V]) LoadOrStoreLazy(key K, f func() V) (actual V, loaded bool) {
	var (
		level        int
		preds, succs [maxLevel]*node[K, V]
		hl           = int(atomic.LoadInt64(&s.highestLevel))
	)
	for {
		nodeFound := s.findNode(key, &preds, &succs)
		if nodeFound != nil { // indicating the key is already in the skip-list
			if !nodeFound.flags.Get(marked) {
				// We don't need to care about whether or not the node is fully linked,
				// just return the value.
				return nodeFound.loadVal(), true
			}
			// If the node is marked, represents some other goroutines is in the process of deleting this node,
			// we need to add this node in next loop.
			continue
		}

		// Add this node into skip list.
		var (
			highestLocked        = -1 // the highest level being locked by this process
			valid                = true
			pred, succ, prevPred *node[K, V]
		)
		if level == 0 {
			level = s.randomlevel()
			if level > hl {
				// If the highest level is updated, usually means that many goroutines
				// are inserting items. Hopefully we can find a better path in next loop.
				// TODO(zyh): consider filling the preds if s.header[level].next == nil,
				// but this strategy's performance is almost the same as the existing method.
				continue
			}
		}
		for layer := 0; valid && layer < level; layer++ {
			pred = preds[layer]   // target node's previous node
			succ = succs[layer]   // target node's next node
			if pred != prevPred { // the node in this layer could be locked by previous loop
				pred.mu.Lock()
				highestLocked = layer
				prevPred = pred
			}
			// valid check if there is another node has inserted into the skip list in this layer during this process.
			// It is valid if:
			// 1. The previous node and next node both are not marked.
			// 2. The previous node's next node is succ in this layer.
			valid = !pred.flags.Get(marked) && pred.loadNext(layer) == succ && (succ == nil || !succ.flags.Get(marked))
		}
		if !valid {
			unlock(preds, highestLocked)
			continue
		}
		value := f()
		nn := newNode(key, value, level, s.comparator)
		for layer := 0; layer < level; layer++ {
			nn.storeNext(layer, succs[layer])
			preds[layer].atomicStoreNext(layer, nn)
		}
		nn.flags.SetTrue(fullyLinked)
		unlock(preds, highestLocked)
		atomic.AddInt64(&s.length, 1)
		return value, false
	}
}

// Delete deletes the value for a key.
func (s *Map[K, V]) Delete(key K) bool {
	var (
		nodeToDelete *node[K, V]
		isMarked     bool // represents if this operation mark the node
		topLayer     = -1
		preds, succs [maxLevel]*node[K, V]
	)
	for {
		lFound := s.findNodeDelete(key, &preds, &succs)
		if isMarked || // this process mark this node or we can find this node in the skip list
			lFound != -1 && succs[lFound].flags.MGet(fullyLinked|marked, fullyLinked) && (int(succs[lFound].level)-1) == lFound {
			if !isMarked { // we don't mark this node for now
				nodeToDelete = succs[lFound]
				topLayer = lFound
				nodeToDelete.mu.Lock()
				if nodeToDelete.flags.Get(marked) {
					// The node is marked by another process,
					// the physical deletion will be accomplished by another process.
					nodeToDelete.mu.Unlock()
					return false
				}
				nodeToDelete.flags.SetTrue(marked)
				isMarked = true
			}
			// Accomplish the physical deletion.
			var (
				highestLocked        = -1 // the highest level being locked by this process
				valid                = true
				pred, succ, prevPred *node[K, V]
			)
			for layer := 0; valid && (layer <= topLayer); layer++ {
				pred, succ = preds[layer], succs[layer]
				if pred != prevPred { // the node in this layer could be locked by previous loop
					pred.mu.Lock()
					highestLocked = layer
					prevPred = pred
				}
				// valid check if there is another node has inserted into the skip list in this layer
				// during this process, or the previous is deleted by another process.
				// It is valid if:
				// 1. the previous node exists.
				// 2. no another node has inserted into the skip list in this layer.
				valid = !pred.flags.Get(marked) && pred.atomicLoadNext(layer) == succ
			}
			if !valid {
				unlock(preds, highestLocked)
				continue
			}
			for i := topLayer; i >= 0; i-- {
				// Now we own the `nodeToDelete`, no other goroutine will modify it.
				// So we don't need `nodeToDelete.loadNext`
				preds[i].atomicStoreNext(i, nodeToDelete.loadNext(i))
			}
			nodeToDelete.mu.Unlock()
			unlock(preds, highestLocked)
			atomic.AddInt64(&s.length, -1)
			return true
		}
		return false
	}
}

// Range calls f sequentially for each key and value present in the skipmap.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
func (s *Map[K, V]) Range(f func(key K, value V) bool) {
	x := s.header.atomicLoadNext(0)
	for x != nil {
		if !x.flags.MGet(fullyLinked|marked, fullyLinked) {
			x = x.atomicLoadNext(0)
			continue
		}
		if !f(x.key, x.loadVal()) {
			break
		}
		x = x.atomicLoadNext(0)
	}
}

// Len return the length of this skipmap.
func (s *Map[K, V]) Len() int {
	return int(atomic.LoadInt64(&s.length))
}
