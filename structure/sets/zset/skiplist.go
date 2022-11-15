package zset

import (
	"math"
	"unsafe"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/sys/fastrand"
)

//
// Skip list implementation.
//

const (
	maxLevel    = 32   // same to ZSKIPLIST_MAXLEVEL, should be enough for 2^64 elements
	probability = 0.25 // same to ZSKIPLIST_P, 1/4
)

// listNode is node of list.
type listNode[K comparable] struct {
	score float64 // key for sorting, which is allowed to be repeated
	value K
	prev  *listNode[K] // back pointer that only available at level 1
	level int          // the length of optionalArray
	oparr optionalArray
}

func newListNode[K comparable](score float64, value K, level int) *listNode[K] {
	node := &listNode[K]{
		score: score,
		value: value,
		level: level,
	}
	node.oparr.init(level)
	return node
}

func (n *listNode[K]) loadNext(i int) *listNode[K] {
	return (*listNode[K])(n.oparr.loadNext(i))
}

func (n *listNode[K]) storeNext(i int, node *listNode[K]) {
	n.oparr.storeNext(i, unsafe.Pointer(node))
}

func (n *listNode[K]) loadSpan(i int) int {
	return n.oparr.loadSpan(i)
}

func (n *listNode[K]) storeSpan(i int, span int) {
	n.oparr.storeSpan(i, span)
}

func (n *listNode[K]) loadNextAndSpan(i int) (*listNode[K], int) {
	return n.loadNext(i), n.loadSpan(i)
}

func (n *listNode[K]) storeNextAndSpan(i int, next *listNode[K], span int) {
	n.storeNext(i, next)
	n.storeSpan(i, span)
}

func (n *listNode[K]) lessThan(score float64, value K, comparator bcomparator.Comparator[K]) bool {
	if n.score < score {
		return true
	}
	if n.score == score {
		return comparator(n.value, value) < 0
	}
	return false
}

func (n *listNode[K]) lessEqual(score float64, value K, comparator bcomparator.Comparator[K]) bool {
	if n.score < score {
		return true
	}
	if n.score == score {
		return comparator(n.value, value) <= 0
	}
	return false
}

func (n *listNode[K]) equal(score float64, value K) bool {
	return n.value == value && n.score == score
}

// list is a specialized skip list implementation for sorted set.
//
// It is almost implement the original
// algorithm described by William Pugh in " Lists: A Probabilistic
// Alternative to Balanced Trees", modified in three ways:
// a) this implementation allows for repeated scores.
// b) the comparison is not just by key (our 'score') but by satellite data(?).
// c) there is a back pointer, so it's a doubly linked list with the back
// pointers being only at "level 1". This allows to traverse the list
// from tail to head, useful for RevRange.
type list[K comparable] struct {
	header       *listNode[K]
	tail         *listNode[K]
	length       int
	highestLevel int // highest level for now
	comparator   bcomparator.Comparator[K]
}

func newList[K comparable](comparator bcomparator.Comparator[K]) *list[K] {
	var zero K
	l := &list[K]{
		header:       newListNode[K](-math.MaxFloat64, zero, maxLevel), // FIXME:
		highestLevel: 1,
		comparator:   comparator,
	}
	return l
}

// Insert inserts a new node in the skiplist. Assumes the element does not already
// exist (up to the caller to enforce that).
func (l *list[K]) Insert(score float64, value K) *listNode[K] {
	var (
		update [maxLevel]*listNode[K]
		rank   [maxLevel + 1]int // +1 for eliminating a boundary judgment
	)

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		rank[i] = rank[i+1] // also fine when i == maxLevel - 1
		next := x.loadNext(i)
		for next != nil && next.lessThan(score, value, l.comparator) {
			rank[i] += x.loadSpan(i)
			x = next
			next = x.loadNext(i)
		}
		update[i] = x
	}

	// We assume the element is not already inside, since we allow duplicated
	// scores, reinserting the same element should never happen since the
	// caller of Add() should test in the hash table if the element is
	// already inside or not.
	level := l.randomLevel()
	if level > l.highestLevel {
		// Create higher levels.
		for i := l.highestLevel; i < level; i++ {
			rank[i] = 0
			update[i] = l.header
			update[i].storeSpan(i, l.length)
		}
		l.highestLevel = level
	}
	x = newListNode(score, value, level)
	for i := 0; i < level; i++ {
		// update --> x --> update.next
		x.storeNext(i, update[i].loadNext(i))
		update[i].storeNext(i, x)
		// update[i].span is splitted to: new update[i].span and x.span
		x.storeSpan(i, update[i].loadSpan(i)-(rank[0]-rank[i]))
		update[i].storeSpan(i, (rank[0]-rank[i])+1)
	}
	// Increment span for untouched levels.
	for i := level; i < l.highestLevel; i++ {
		update[i].storeSpan(i, update[i].loadSpan(i)+1)
	}

	// Update back pointer.
	if update[0] != l.header {
		x.prev = update[0]
	}

	if next := x.loadNext(0); next != nil { // not tail of skiplist
		next.prev = x
	} else {
		l.tail = x
	}
	l.length++

	return x
}

// randomLevel returns a level between [1, maxLevel] for insertion.
func (l *list[K]) randomLevel() int {
	level := 1
	for fastrand.Uint32n(1/probability) == 0 {
		level++
	}
	if level > maxLevel {
		return maxLevel
	}
	return level
}

// Rank finds the rank for an element by both score and value.
// Returns 0 when the element cannot be found, rank otherwise.
//
// NOTE: the rank is 1-based due to the span of l->header to the
// first element.
func (l *list[K]) Rank(score float64, value K) int {
	rank := 0
	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next := x.loadNext(i)
		for next != nil && next.lessEqual(score, value, l.comparator) {
			rank += x.loadSpan(i)
			x = next
			next = x.loadNext(i)
		}

		// x might be equal to l->header, so test if obj is non-nil
		// TODO: Why not use if x != l.header?
		if x.equal(score, value) {
			return rank
		}
	}
	return 0
}

// deleteNode is a internal function for deleting node x in O(1) time by giving a
// update position matrix.
func (l *list[K]) deleteNode(x *listNode[K], update *[maxLevel]*listNode[K]) {
	for i := 0; i < l.highestLevel; i++ {
		if update[i].loadNext(i) == x {
			// Remove x, updaet[i].span = updaet[i].span + x.span - 1 (x removed).
			next, span := x.loadNextAndSpan(i)
			span += update[i].loadSpan(i) - 1
			update[i].storeNextAndSpan(i, next, span)
		} else {
			// x does not appear on this level, just update span.
			update[i].storeSpan(i, update[i].loadSpan(i)-1)
		}
	}
	if next := x.loadNext(0); next != nil { // not tail of skiplist
		next.prev = x.prev
	} else {
		l.tail = x.prev
	}
	for l.highestLevel > 1 && l.header.loadNext(l.highestLevel-1) != nil {
		// Clear the pointer and span for safety.
		l.header.storeNextAndSpan(l.highestLevel-1, nil, 0)
		l.highestLevel--
	}
	l.length--
}

// Delete deletes an element with matching score/element from the skiplist.
// The deleted node is returned if the node was found, otherwise 0 is returned.
func (l *list[K]) Delete(score float64, value K) *listNode[K] {
	var update [maxLevel]*listNode[K]

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next := x.loadNext(i)
		for next != nil && next.lessThan(score, value, l.comparator) {
			x = next
			next = x.loadNext(i)
		}
		update[i] = x
	}
	x = x.loadNext(0)
	if x != nil && x.equal(score, value) {
		l.deleteNode(x, &update)
		return x
	}
	return nil // not found
}

// UpdateScore updates the score of an element inside the sorted set skiplist.
//
// NOTE: the element must exist and must match 'score'.
// This function does not update the score in the hash table side, the
// caller should take care of it.
//
// NOTE: this function attempts to just update the node, in case after
// the score update, the node would be exactly at the same position.
// Otherwise the skiplist is modified by removing and re-adding a new
// element, which is more costly.
//
// The function returns the updated element skiplist node pointer.
func (l *list[K]) UpdateScore(oldScore float64, value K, newScore float64) *listNode[K] {
	var update [maxLevel]*listNode[K]

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next := x.loadNext(i)
		for next != nil && next.lessThan(oldScore, value, l.comparator) {
			x = next
			next = x.loadNext(i)
		}
		update[i] = x
	}

	// Jump to our element: note that this function assumes that the
	// element with the matching score exists.
	x = x.loadNext(0)

	// Fastpath: If the node, after the score update, would be still exactly
	// at the same position, we can just update the score without
	// actually removing and re-inserting the element in the skiplist.
	if next := x.loadNext(0); (x.prev == nil || x.prev.score < newScore) &&
		(next == nil || next.score > newScore) {
		x.score = newScore
		return x
	}

	// No way to reuse the old node: we need to remove and insert a new
	// one at a different place.
	v := x.value
	l.deleteNode(x, &update)
	newNode := l.Insert(newScore, v)
	return newNode
}

func greaterThanMin(value float64, min float64, ex bool) bool {
	if ex {
		return value > min
	}
	return value >= min
}

func lessThanMax(value float64, max float64, ex bool) bool {
	if ex {
		return value < max
	}
	return value <= max
}

// DeleteRangeByScore deletes all the elements with score between min and max
// from the skiplist.
// Both min and max can be inclusive or exclusive (see RangeOpt).
// When inclusive a score >= min && score <= max is deleted.
//
// This function returns count of deleted elements.
func (l *list[K]) DeleteRangeByScore(min, max float64, opt RangeOpt, dict map[K]float64) []Node[K] {
	var (
		update  [maxLevel]*listNode[K]
		removed []Node[K]
	)

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next := x.loadNext(i)
		for next != nil && !greaterThanMin(next.score, min, opt.ExcludeMin) {
			x = next
			next = x.loadNext(i)
		}
		update[i] = x
	}

	// Current node is the last with score not greater than min.
	x = x.loadNext(0)

	// Delete nodes in range.
	for x != nil && lessThanMax(x.score, max, opt.ExcludeMax) {
		next := x.loadNext(0)
		l.deleteNode(x, &update)
		delete(dict, x.value)
		removed = append(removed, Node[K]{
			Value: x.value,
			Score: x.score,
		})
		x = next
	}

	return removed
}

// Delete all the elements with rank between start and end from the skiplist.
// Start and end are inclusive.
//
// NOTE: start and end need to be 1-based
func (l *list[K]) DeleteRangeByRank(start, end int, dict map[K]float64) []Node[K] {
	var (
		update    [maxLevel]*listNode[K]
		removed   []Node[K]
		traversed int
	)

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next, span := x.loadNextAndSpan(i)
		for next != nil && traversed+span < start {
			traversed += span
			x = next
			next, span = x.loadNextAndSpan(i)
		}
		update[i] = x
	}

	traversed++
	x = x.loadNext(0)
	// Delete nodes in range.
	for x != nil && traversed <= end {
		next := x.loadNext(0)
		l.deleteNode(x, &update)
		delete(dict, x.value)
		removed = append(removed, Node[K]{
			Value: x.value,
			Score: x.score,
		})
		traversed++
		x = next
	}
	return removed
}

// GetNodeByRank finds an element by its rank. The rank argument needs to be 1-based.
func (l *list[K]) GetNodeByRank(rank int) *listNode[K] {
	var traversed int

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next, span := x.loadNextAndSpan(i)
		for next != nil && traversed+span <= rank {
			traversed += span
			x = next
			next, span = x.loadNextAndSpan(i)
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}

// FirstInRange finds the first node that is contained in the specified range.
func (l *list[K]) FirstInRange(min, max float64, opt RangeOpt) *listNode[K] {
	if !l.IsInRange(min, max, opt) {
		return nil
	}

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next := x.loadNext(i)
		for next != nil && !greaterThanMin(next.score, min, opt.ExcludeMin) {
			x = next
			next = x.loadNext(i)
		}
	}

	// The next node MUST not be NULL (excluded by IsInRange).
	x = x.loadNext(0)
	if !lessThanMax(x.score, max, opt.ExcludeMax) {
		return nil
	}
	return x
}

// LastInRange finds the last node that is contained in the specified range.
func (l *list[K]) LastInRange(min, max float64, opt RangeOpt) *listNode[K] {
	if !l.IsInRange(min, max, opt) {
		return nil
	}

	x := l.header
	for i := l.highestLevel - 1; i >= 0; i-- {
		next := x.loadNext(i)
		for next != nil && lessThanMax(next.score, max, opt.ExcludeMax) {
			x = next
			next = x.loadNext(i)
		}
	}

	// The node x must not be NULL (excluded by IsInRange).
	if !greaterThanMin(x.score, min, opt.ExcludeMin) {
		return nil
	}
	return x
}

// IsInRange returns whether there is a port of sorted set in given range.
func (l *list[K]) IsInRange(min, max float64, opt RangeOpt) bool {
	// Test empty range.
	if min > max || (min == max && (opt.ExcludeMin || opt.ExcludeMax)) {
		return false
	}
	if l.tail == nil || !greaterThanMin(l.tail.score, min, opt.ExcludeMin) {
		return false
	}
	if next := l.header.loadNext(0); next == nil || !lessThanMax(next.score, max, opt.ExcludeMax) {
		return false
	}
	return true
}
