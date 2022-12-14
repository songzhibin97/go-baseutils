// Package skipset is a high-performance, scalable, concurrent-safe set based on skip-list.
// In the typical pattern(100000 operations, 90%CONTAINS 9%Add 1%Remove, 8C16T), the skipset
// up to 15x faster than the built-in sync.Map.
package skipset

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/sets"
)

// Assert Set implementation
var _ sets.Set[any] = (*Set[any])(nil)

type Set[E any] struct {
	header       *node[E]
	length       int64
	highestLevel int64 // highest level for now
	comparator   bcomparator.Comparator[E]
}

type node[E any] struct {
	value E
	next  optionalArray // [level]*node[E]
	mu    sync.Mutex
	flags bitflag
	level uint32
}

func newNode[E any](e E, level int) *node[E] {
	node := &node[E]{
		value: e,
		level: uint32(level),
	}
	if level > op1 {
		node.next.extra = new([op2]unsafe.Pointer)
	}
	return node
}

func (n *node[E]) loadNext(i int) *node[E] {
	return (*node[E])(n.next.load(i))
}

func (n *node[E]) storeNext(i int, node *node[E]) {
	n.next.store(i, unsafe.Pointer(node))
}

func (n *node[E]) atomicLoadNext(i int) *node[E] {
	return (*node[E])(n.next.atomicLoad(i))
}

func (n *node[E]) atomicStoreNext(i int, node *node[E]) {
	n.next.atomicStore(i, unsafe.Pointer(node))
}

func (n *node[E]) lessthan(value E, comparable bcomparator.Comparator[E]) bool {
	return comparable(n.value, value) < 0
}

func (n *node[E]) equal(value E, comparable bcomparator.Comparator[E]) bool {
	return comparable(n.value, value) == 0
}

func New[E any](comparator bcomparator.Comparator[E]) *Set[E] {
	var zero E
	h := newNode[E](zero, maxLevel)
	h.flags.SetTrue(fullyLinked)
	return &Set[E]{
		header:       h,
		highestLevel: defaultHighestLevel,
		comparator:   comparator,
	}
}

// findNodeRemove takes a value and two maximal-height arrays then searches exactly as in a sequential skip-list.
// The returned preds and succs always satisfy preds[i] > value >= succs[i].
func (s *Set[E]) findNodeRemove(value E, preds *[maxLevel]*node[E], succs *[maxLevel]*node[E]) int {
	// lFound represents the index of the first layer at which it found a node.
	lFound, x := -1, s.header
	for i := int(atomic.LoadInt64(&s.highestLevel)) - 1; i >= 0; i-- {
		succ := x.atomicLoadNext(i)
		for succ != nil && succ.lessthan(value, s.comparator) {
			x = succ
			succ = x.atomicLoadNext(i)
		}
		preds[i] = x
		succs[i] = succ

		// Check if the value already in the skip list.
		if lFound == -1 && succ != nil && succ.equal(value, s.comparator) {
			lFound = i
		}
	}
	return lFound
}

// findNodeAdd takes a value and two maximal-height arrays then searches exactly as in a sequential skip-set.
// The returned preds and succs always satisfy preds[i] > value >= succs[i].
func (s *Set[E]) findNodeAdd(value E, preds *[maxLevel]*node[E], succs *[maxLevel]*node[E]) int {
	x := s.header
	for i := int(atomic.LoadInt64(&s.highestLevel)) - 1; i >= 0; i-- {
		succ := x.atomicLoadNext(i)
		for succ != nil && succ.lessthan(value, s.comparator) {
			x = succ
			succ = x.atomicLoadNext(i)
		}
		preds[i] = x
		succs[i] = succ

		// Check if the value already in the skip list.
		if succ != nil && succ.equal(value, s.comparator) {
			return i
		}
	}
	return -1
}

func unlockInt64[E any](preds [maxLevel]*node[E], highestLevel int) {
	var prevPred *node[E]
	for i := highestLevel; i >= 0; i-- {
		if preds[i] != prevPred { // the node could be unlocked by previous loop
			preds[i].mu.Unlock()
			prevPred = preds[i]
		}
	}
}

// AddB add the value into skip set, return true if this process insert the value into skip set,
// return false if this process can't insert this value, because another process has insert the same value.
//
// If the value is in the skip set but not fully linked, this process will wait until it is.
func (s *Set[E]) AddB(value E) bool {
	level := s.randomLevel()
	var preds, succs [maxLevel]*node[E]
	for {
		lFound := s.findNodeAdd(value, &preds, &succs)
		if lFound != -1 { // indicating the value is already in the skip-list
			nodeFound := succs[lFound]
			if !nodeFound.flags.Get(marked) {
				for !nodeFound.flags.Get(fullyLinked) {
					// The node is not yet fully linked, just waits until it is.
				}
				return false
			}
			// If the node is marked, represents some other thread is in the process of deleting this node,
			// we need to add this node in next loop.
			continue
		}
		// Add this node into skip list.
		var (
			highestLocked        = -1 // the highest level being locked by this process
			valid                = true
			pred, succ, prevPred *node[E]
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
			unlockInt64(preds, highestLocked)
			continue
		}

		nn := newNode[E](value, level)
		for layer := 0; layer < level; layer++ {
			nn.storeNext(layer, succs[layer])
			preds[layer].atomicStoreNext(layer, nn)
		}
		nn.flags.SetTrue(fullyLinked)
		unlockInt64(preds, highestLocked)
		atomic.AddInt64(&s.length, 1)
		return true
	}
}

func (s *Set[E]) Add(values ...E) {
	for _, value := range values {
		s.AddB(value)
	}
}

func (s *Set[E]) randomLevel() int {
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

// ContainsB check if the value is in the skip set.
func (s *Set[E]) ContainsB(value E) bool {
	x := s.header
	for i := int(atomic.LoadInt64(&s.highestLevel)) - 1; i >= 0; i-- {
		nex := x.atomicLoadNext(i)
		for nex != nil && nex.lessthan(value, s.comparator) {
			x = nex
			nex = x.atomicLoadNext(i)
		}

		// Check if the value already in the skip list.
		if nex != nil && nex.equal(value, s.comparator) {
			return nex.flags.MGet(fullyLinked|marked, fullyLinked)
		}
	}
	return false
}

func (s *Set[E]) Contains(values ...E) bool {
	for _, value := range values {
		if !s.ContainsB(value) {
			return false
		}
	}
	return true
}

// RemoveB a node from the skip set.
func (s *Set[E]) RemoveB(value E) bool {
	var (
		nodeToRemove *node[E]
		isMarked     bool // represents if this operation mark the node
		topLayer     = -1
		preds, succs [maxLevel]*node[E]
	)
	for {
		lFound := s.findNodeRemove(value, &preds, &succs)
		if isMarked || // this process mark this node or we can find this node in the skip list
			lFound != -1 && succs[lFound].flags.MGet(fullyLinked|marked, fullyLinked) && (int(succs[lFound].level)-1) == lFound {
			if !isMarked { // we don't mark this node for now
				nodeToRemove = succs[lFound]
				topLayer = lFound
				nodeToRemove.mu.Lock()
				if nodeToRemove.flags.Get(marked) {
					// The node is marked by another process,
					// the physical deletion will be accomplished by another process.
					nodeToRemove.mu.Unlock()
					return false
				}
				nodeToRemove.flags.SetTrue(marked)
				isMarked = true
			}
			// Accomplish the physical deletion.
			var (
				highestLocked        = -1 // the highest level being locked by this process
				valid                = true
				pred, succ, prevPred *node[E]
			)
			for layer := 0; valid && (layer <= topLayer); layer++ {
				pred, succ = preds[layer], succs[layer]
				if pred != prevPred { // the node in this layer could be locked by previous loop
					pred.mu.Lock()
					highestLocked = layer
					prevPred = pred
				}
				// valid check if there is another node has inserted into the skip list in this layer
				// during this process, or the previous is removed by another process.
				// It is valid if:
				// 1. the previous node exists.
				// 2. no another node has inserted into the skip list in this layer.
				valid = !pred.flags.Get(marked) && pred.loadNext(layer) == succ
			}
			if !valid {
				unlockInt64(preds, highestLocked)
				continue
			}
			for i := topLayer; i >= 0; i-- {
				// Now we own the `nodeToRemove`, no other goroutine will modify it.
				// So we don't need `nodeToRemove.loadNext`
				preds[i].atomicStoreNext(i, nodeToRemove.loadNext(i))
			}
			nodeToRemove.mu.Unlock()
			unlockInt64(preds, highestLocked)
			atomic.AddInt64(&s.length, -1)
			return true
		}
		return false
	}
}

func (s *Set[E]) Remove(values ...E) {
	for _, value := range values {
		s.RemoveB(value)
	}
}

// Range calls f sequentially for each value present in the skip set.
// If f returns false, range stops the iteration.
func (s *Set[E]) Range(f func(value E) bool) {
	x := s.header.atomicLoadNext(0)
	for x != nil {
		if !x.flags.MGet(fullyLinked|marked, fullyLinked) {
			x = x.atomicLoadNext(0)
			continue
		}
		if !f(x.value) {
			break
		}
		x = x.atomicLoadNext(0)
	}
}

// Len return the length of this skip set.
func (s *Set[E]) Len() int {
	return int(atomic.LoadInt64(&s.length))
}

func (s *Set[E]) Empty() bool {
	return s.Len() == 0
}

func (s *Set[E]) Size() int {
	return s.Len()
}

func (s *Set[E]) Clear() {
	var zero E
	h := newNode[E](zero, maxLevel)
	h.flags.SetTrue(fullyLinked)
	s.header = h
	s.highestLevel = defaultHighestLevel
}

func (s *Set[E]) Values() []E {
	ln := s.Len()
	vals := make([]E, 0, ln)
	s.Range(func(e E) bool {
		vals = append(vals, e)
		return true
	})
	return vals
}

func (s *Set[E]) String() string {
	b := strings.Builder{}
	b.WriteString("SkipSet\n")
	for val := range s.Values() {
		b.WriteString(fmt.Sprintf("(key:%v) ", val))
	}
	return b.String()
}
