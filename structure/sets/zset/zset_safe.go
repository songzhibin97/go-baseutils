package zset

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/structure/sets"
	"sync"
)

var _ sets.Set[int] = (*SetSafe[int])(nil)

func NewSafe[K comparable](comparator bcomparator.Comparator[K]) *SetSafe[K] {
	return &SetSafe[K]{
		unsafe: New[K](comparator),
	}
}

type SetSafe[K comparable] struct {
	unsafe *Set[K]
	lock   sync.Mutex
}

func (s *SetSafe[K]) Add(elements ...K) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Add(elements...)

}

func (s *SetSafe[K]) Remove(elements ...K) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Remove(elements...)

}

func (s *SetSafe[K]) Contains(elements ...K) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Contains(elements...)
}

func (s *SetSafe[K]) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Empty()
}

func (s *SetSafe[K]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Size()
}

func (s *SetSafe[K]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unsafe.Clear()

}

func (s *SetSafe[K]) Values() []K {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Values()
}

func (s *SetSafe[K]) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.String()
}

func (s *SetSafe[K]) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Len()
}

func (s *SetSafe[K]) AddB(score float64, value K) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.AddB(score, value)
}

func (s *SetSafe[K]) RemoveB(value K) (float64, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RemoveB(value)
}

func (s *SetSafe[K]) IncrBy(incr float64, value K) (float64, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.IncrBy(incr, value)
}

func (s *SetSafe[K]) ContainsB(value K) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.ContainsB(value)
}

func (s *SetSafe[K]) Score(value K) (float64, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Score(value)
}

func (s *SetSafe[K]) Rank(value K) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Rank(value)
}

func (s *SetSafe[K]) RevRank(value K) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RevRank(value)
}

func (s *SetSafe[K]) Count(min, max float64) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Count(min, max)
}

func (s *SetSafe[K]) CountWithOpt(min, max float64, opt RangeOpt) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.CountWithOpt(min, max, opt)
}

func (s *SetSafe[K]) Range(start, stop int) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.Range(start, stop)
}

func (s *SetSafe[K]) RangeByScore(min, max float64) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RangeByScore(min, max)
}

func (s *SetSafe[K]) RangeByScoreWithOpt(min, max float64, opt RangeOpt) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RangeByScoreWithOpt(min, max, opt)
}

func (s *SetSafe[K]) RevRange(start, stop int) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RevRange(start, stop)
}

func (s *SetSafe[K]) RevRangeByScore(max, min float64) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RevRangeByScore(max, min)
}

func (s *SetSafe[K]) RevRangeByScoreWithOpt(max, min float64, opt RangeOpt) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RevRangeByScoreWithOpt(max, min, opt)
}

func (s *SetSafe[K]) RemoveRangeByRank(start, stop int) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RemoveRangeByRank(start, stop)
}

func (s *SetSafe[K]) RemoveRangeByScore(min, max float64) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RemoveRangeByScore(min, max)
}

func (s *SetSafe[K]) RemoveRangeByScoreWithOpt(min, max float64, opt RangeOpt) []Node[K] {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.unsafe.RemoveRangeByScoreWithOpt(min, max, opt)
}
