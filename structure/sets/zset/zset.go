// Package zset provides a concurrent-safety sorted set, can be used as a local
// replacement of Redis' zset (https://redis.com/ebook/part-2-core-concepts/chapter-3-commands-in-redis/3-5-sorted-sets/).
//
// The main different to other sets is, every value of set is associated with a score,
// that is used in order to take the sorted set ordered, from the smallest to the greatest score.
//
// The sorted set has O(log(N)) time complexity when doing Add(ZADD) and
// Remove(ZREM) operations and O(1) time complexity when doing Contains operations.
package zset

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/sets"
)

var _ sets.Set[int] = (*Set[int])(nil)

// Node represents an element of Set.
type Node[K comparable] struct {
	Value K
	Score float64
}

// Set is a sorted set implementation with string value and float64 score.
type Set[K comparable] struct {
	mu   sync.RWMutex
	dict map[K]float64
	list *list[K]
}

func (z *Set[K]) Add(elements ...K) {
	for _, element := range elements {
		z.AddB(0, element)
	}
}

func (z *Set[K]) Remove(elements ...K) {
	for _, element := range elements {
		z.RemoveB(element)
	}
}

func (z *Set[K]) Contains(elements ...K) bool {
	for _, element := range elements {
		if !z.ContainsB(element) {
			return false
		}
	}
	return true
}

func (z *Set[K]) Empty() bool {
	return z.Size() == 0
}

func (z *Set[K]) Size() int {
	return z.Count(0, math.MaxFloat64)
}

func (z *Set[K]) Clear() {
	z.dict = make(map[K]float64)
	z.list = newList(z.list.comparator)
}

func (z *Set[K]) Values() []K {
	list := z.Range(0, -1)
	ret := make([]K, 0, len(list))
	for _, n := range list {
		ret = append(ret, n.Value)
	}
	return ret
}

func (z *Set[K]) String() string {
	b := strings.Builder{}
	b.WriteString("ZSet\n")
	for _, v := range z.Values() {
		b.WriteString(fmt.Sprintf("(key:%v) ", v))
	}
	return b.String()
}

// New returns an empty string sorted set with int score.
// strings are sorted in ascending order.
func New[K comparable](comparator bcomparator.Comparator[K]) *Set[K] {
	return &Set[K]{
		dict: make(map[K]float64),
		list: newList(comparator),
	}
}

// Union returns the union of given sorted sets, the resulting score of
// a value is the sum of its scores in the sorted sets where it exists.
//
// Union is the replacement of UNIONSTORE command of redis.
func Union[K comparable](comparator bcomparator.Comparator[K], zs ...*Set[K]) *Set[K] {
	dest := New(comparator)
	for _, z := range zs {
		for _, n := range z.Range(0, -1) {
			dest.AddB(n.Score, n.Value)
		}
	}
	return dest
}

// Inter returns the intersection of given sorted sets, the resulting
// score of a value is the sum of its scores in the sorted sets where it exists.
//
// Inter is the replacement of INTERSTORE command of redis.
func Inter[K comparable](comparator bcomparator.Comparator[K], zs ...*Set[K]) *Set[K] {
	dest := New(comparator)
	if len(zs) == 0 {
		return dest
	}
	for _, n := range zs[0].Range(0, -1) {
		ok := true
		for _, z := range zs[1:] {
			if !z.Contains(n.Value) {
				ok = false
				break
			}
		}
		if ok {
			dest.AddB(n.Score, n.Value)
		}
	}
	return dest
}

// Len returns the length of Set.
//
// Len is the replacement of ZCARD command of redis.
func (z *Set[K]) Len() int {
	z.mu.RLock()
	defer z.mu.RUnlock()

	return z.list.length
}

// Add adds a new value or update the score of an existing value.
// Returns true if the value is newly created.
//
// Add is the replacement of ZADD command of redis.
func (z *Set[K]) AddB(score float64, value K) bool {
	z.mu.Lock()
	defer z.mu.Unlock()

	oldScore, ok := z.dict[value]
	if ok {
		// Update score if need.
		if score != oldScore {
			_ = z.list.UpdateScore(oldScore, value, score)
			z.dict[value] = score
		}
		return false
	}

	// Insert a new element.
	z.list.Insert(score, value)
	z.dict[value] = score
	return true
}

// Remove removes a value from the sorted set.
// Returns score of the removed value and true if the node was found and deleted,
// otherwise returns (0.0, false).
//
// Remove is the replacement of ZREM command of redis.
func (z *Set[K]) RemoveB(value K) (float64, bool) {
	z.mu.Lock()
	defer z.mu.Unlock()

	score, ok := z.dict[value]
	if !ok {
		return 0, false
	}
	delete(z.dict, value)
	z.list.Delete(score, value)
	return score, true
}

// IncrBy increments the score of value in the sorted set by incr.
// If value does not exist in the sorted set, it is added with incr as its score
// (as if its previous score was zero).
//
// IncrBy is the replacement of ZINCRBY command of redis.
func (z *Set[K]) IncrBy(incr float64, value K) (float64, bool) {
	z.mu.Lock()
	defer z.mu.Unlock()

	oldScore, ok := z.dict[value]
	if !ok {
		// Insert a new element.
		z.list.Insert(incr, value)
		z.dict[value] = incr
		return incr, false
	}
	// Update score.
	newScore := oldScore + incr
	_ = z.list.UpdateScore(oldScore, value, newScore)
	z.dict[value] = newScore
	return newScore, true
}

// Contains returns whether the value exists in sorted set.
func (z *Set[K]) ContainsB(value K) bool {
	_, ok := z.Score(value)
	return ok
}

// Score returns the score of the value in the sorted set.
//
// Score is the replacement of ZSCORE command of redis.
func (z *Set[K]) Score(value K) (float64, bool) {
	z.mu.RLock()
	defer z.mu.RUnlock()

	score, ok := z.dict[value]
	return score, ok
}

// Rank returns the rank of element in the sorted set, with the scores
// ordered from low to high.
// The rank (or index) is 0-based, which means that the member with the lowest
// score has rank 0.
// -1 is returned when value is not found.
//
// Rank is the replacement of ZRANK command of redis.
func (z *Set[K]) Rank(value K) int {
	z.mu.RLock()
	defer z.mu.RUnlock()

	score, ok := z.dict[value]
	if !ok {
		return -1
	}
	// NOTE: list.Rank returns 1-based rank.
	return z.list.Rank(score, value) - 1
}

// RevRank returns the rank of element in the sorted set, with the scores
// ordered from high to low.
// The rank (or index) is 0-based, which means that the member with the highest
// score has rank 0.
// -1 is returned when value is not found.
//
// RevRank is the replacement of ZREVRANK command of redis.
func (z *Set[K]) RevRank(value K) int {
	z.mu.RLock()
	defer z.mu.RUnlock()

	score, ok := z.dict[value]
	if !ok {
		return -1
	}
	// NOTE: list.Rank returns 1-based rank.
	return z.list.Rank(score, value) - 1
}

// Count returns the number of elements in the sorted set at element with a score
// between min and max (including elements with score equal to min or max).
//
// Count is the replacement of ZCOUNT command of redis.
func (z *Set[K]) Count(min, max float64) int {
	return z.CountWithOpt(min, max, RangeOpt{})
}

func (z *Set[K]) CountWithOpt(min, max float64, opt RangeOpt) int {
	z.mu.RLock()
	defer z.mu.RUnlock()

	first := z.list.FirstInRange(min, max, opt)
	if first == nil {
		return 0
	}
	// Sub 1 for 1-based rank.
	firstRank := z.list.Rank(first.score, first.value) - 1
	last := z.list.LastInRange(min, max, opt)
	if last == nil {
		return z.list.length - firstRank
	}
	// Sub 1 for 1-based rank.
	lastRank := z.list.Rank(last.score, last.value) - 1
	return lastRank - firstRank + 1
}

// Range returns the specified inclusive range of elements in the sorted set by rank(index).
// Both start and stop are 0-based, they can also be negative numbers indicating
// offsets from the end of the sorted set, with -1 being the last element of the sorted set,
// and so on.
//
// The returned elements are ordered by score, from lowest to highest.
// Elements with the same score are ordered lexicographically.
//
// This function won't panic even when the given rank out of range.
//
// NOTE: Please always use z.Range(0, -1) for iterating the whole sorted set.
// z.Range(0, z.Len()-1) has 2 method calls, the sorted set may changes during
// the gap of calls.
//
// Range is the replacement of ZRANGE command of redis.
func (z *Set[K]) Range(start, stop int) []Node[K] {
	z.mu.RLock()
	defer z.mu.RUnlock()

	// Convert negative rank to positive.
	if start < 0 {
		start = z.list.length + start
	}
	if stop < 0 {
		stop = z.list.length + stop
	}

	var res []Node[K]
	x := z.list.GetNodeByRank(start + 1) // 0-based rank -> 1-based rank
	for x != nil && start <= stop {
		start++
		res = append(res, Node[K]{
			Score: x.score,
			Value: x.value,
		})
		x = x.loadNext(0)
	}
	return res
}

// RangeByScore returns all the elements in the sorted set with a score
// between min and max (including elements with score equal to min or max).
// The elements are considered to be ordered from low to high scores.
//
// RangeByScore is the replacement of ZRANGEBYSCORE command of redis.
func (z *Set[K]) RangeByScore(min, max float64) []Node[K] {
	return z.RangeByScoreWithOpt(min, max, RangeOpt{})
}

func (z *Set[K]) RangeByScoreWithOpt(min, max float64, opt RangeOpt) []Node[K] {
	z.mu.RLock()
	defer z.mu.RUnlock()

	var res []Node[K]
	x := z.list.FirstInRange(min, max, opt)
	for x != nil && (x.score < max || (!opt.ExcludeMax && x.score == max)) {
		res = append(res, Node[K]{
			Score: x.score,
			Value: x.value,
		})
		x = x.loadNext(0)
	}
	return res
}

// RevRange returns the specified inclusive range of elements in the sorted set by rank(index).
// Both start and stop are 0-based, they can also be negative numbers indicating
// offsets from the end of the sorted set, with -1 being the first element of the sorted set,
// and so on.
//
// The returned elements are ordered by score, from highest to lowest.
// Elements with the same score are ordered in reverse lexicographical ordering.
//
// This function won't panic even when the given rank out of range.
//
// NOTE: Please always use z.RevRange(0, -1) for iterating the whole sorted set.
// z.RevRange(0, z.Len()-1) has 2 method calls, the sorted set may changes during
// the gap of calls.
//
// RevRange is the replacement of ZREVRANGE command of redis.
func (z *Set[K]) RevRange(start, stop int) []Node[K] {
	z.mu.RLock()
	defer z.mu.RUnlock()

	// Convert negative rank to positive.
	if start < 0 {
		start = z.list.length + start
	}
	if stop < 0 {
		stop = z.list.length + stop
	}

	var res []Node[K]
	x := z.list.GetNodeByRank(z.list.length - start) // 0-based rank -> 1-based rank
	for x != nil && start <= stop {
		start++
		res = append(res, Node[K]{
			Score: x.score,
			Value: x.value,
		})
		x = x.prev
	}
	return res
}

// RevRangeByScore returns all the elements in the sorted set with a
// score between max and min (including elements with score equal to max or min).
// The elements are considered to be ordered from high to low scores.
//
// RevRangeByScore is the replacement of ZREVRANGEBYSCORE command of redis.
func (z *Set[K]) RevRangeByScore(max, min float64) []Node[K] {
	return z.RevRangeByScoreWithOpt(max, min, RangeOpt{})
}

func (z *Set[K]) RevRangeByScoreWithOpt(max, min float64, opt RangeOpt) []Node[K] {
	z.mu.RLock()
	defer z.mu.RUnlock()

	var res []Node[K]
	x := z.list.LastInRange(min, max, opt)
	for x != nil && (x.score > min || (!opt.ExcludeMin && x.score == min)) {
		res = append(res, Node[K]{
			Score: x.score,
			Value: x.value,
		})
		x = x.prev
	}
	return res
}

// RemoveRangeByRank removes all elements in the sorted set stored with rank
// between start and stop.
// Both start and stop are 0-based, they can also be negative numbers indicating
// offsets from the end of the sorted set, with -1 being the last element of the sorted set,
// and so on.
//
// RemoveRangeByRank is the replacement of ZREMRANGEBYRANK command of redis.
func (z *Set[K]) RemoveRangeByRank(start, stop int) []Node[K] {
	z.mu.RLock()
	defer z.mu.RUnlock()

	// Convert negative rank to positive.
	if start < 0 {
		start = z.list.length + start
	}
	if stop < 0 {
		stop = z.list.length + stop
	}

	return z.list.DeleteRangeByRank(start+1, stop+1, z.dict) // 0-based rank -> 1-based rank
}

// RemoveRangeByScore removes all elements in the sorted set stored with a score
// between min and max (including elements with score equal to min or max).
//
// RemoveRangeByScore is the replacement of ZREMRANGEBYSCORE command of redis.
func (z *Set[K]) RemoveRangeByScore(min, max float64) []Node[K] {
	return z.RevRangeByScoreWithOpt(min, max, RangeOpt{})
}

func (z *Set[K]) RemoveRangeByScoreWithOpt(min, max float64, opt RangeOpt) []Node[K] {
	z.mu.RLock()
	defer z.mu.RUnlock()

	return z.list.DeleteRangeByScore(min, max, opt, z.dict)
}
