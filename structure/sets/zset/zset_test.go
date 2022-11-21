package zset

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/sys/fastrand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(prefix string) string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[fastrand.Intn(len(letterRunes))]
	}
	return prefix + string(b)
}

func TestSet(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	assert.Zero(t, z.Len())
}

func TestSetAdd(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	v := randString("")
	assert.True(t, z.AddB(1, v))
	assert.False(t, z.AddB(1, v))
}

func TestSetContains(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	v := randString("")
	z.AddB(1, v)
	assert.True(t, z.Contains(v))
	assert.False(t, z.Contains("no-such-"+v))
}

func TestSetScore(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	v := randString("")
	s := rand.Float64()
	z.AddB(s, v)
	as, ok := z.Score(v)
	assert.True(t, ok)
	assert.Equal(t, s, as)
	_, ok = z.Score("no-such-" + v)
	assert.False(t, ok)
}

func TestSetIncr(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	_, ok := z.Score("t")
	assert.False(t, ok)

	// test first insert
	s, ok := z.IncrBy(1, "t")
	assert.False(t, ok)
	assert.Equal(t, 1.0, s)

	// test regular incr
	s, ok = z.IncrBy(2, "t")
	assert.True(t, ok)
	assert.Equal(t, 3.0, s)
}

func TestSetRemove(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	// test first insert
	ok := z.AddB(1, "t")
	assert.True(t, ok)
	_, ok = z.RemoveB("t")
	assert.True(t, ok)
}

func TestSetRank(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	v := randString("")
	z.AddB(1, v)
	// test rank of exist value
	assert.Equal(t, 0, z.Rank(v))
	// test rank of non-exist value
	assert.Equal(t, -1, z.Rank("no-such-"+v))
}

func TestSetRank_Many(t *testing.T) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	rand.Seed(time.Now().Unix())

	var vs []string
	for i := 0; i < N; i++ {
		v := randString("")
		z.AddB(rand.Float64(), v)
		vs = append(vs, v)
	}
	for _, v := range vs {
		r := z.Rank(v)
		assert.NotEqual(t, -1, r)

		// verify rank by traversing level 0
		actualRank := 0
		x := z.list.header
		for x != nil {
			x = x.loadNext(0)
			if x.value == v {
				break
			}
			actualRank++
		}
		assert.Equal(t, v, x.value)
		assert.Equal(t, r, actualRank)
	}
}

func TestSetRank_UpdateScore(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	rand.Seed(time.Now().Unix())

	var vs []string
	for i := 0; i < 100; i++ {
		v := fmt.Sprint(i)
		z.AddB(rand.Float64(), v)
		vs = append(vs, v)
	}
	// Randomly update score
	for i := 0; i < 100; i++ {
		// 1/2
		if rand.Float64() > 0.5 {
			continue
		}
		z.AddB(float64(i), fmt.Sprint(i))
	}

	for _, v := range vs {
		r := z.Rank(v)
		assert.NotEqual(t, -1, r)
		assert.Greater(t, z.Len(), r)

		// verify rank by traversing level 0
		actualRank := 0
		x := z.list.header
		for x != nil {
			x = x.loadNext(0)
			if x.value == v {
				break
			}
			actualRank++
		}
		assert.Equal(t, v, x.value)
		assert.Equal(t, r, actualRank)
	}
}

// Test whether the ramdom inserted values sorted
func TestSetIsSorted(t *testing.T) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	rand.Seed(time.Now().Unix())

	// Test whether the ramdom inserted values sorted
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}
	testIsSorted(t, z)
	testInternalSpan(t, z)

	// Randomly update score
	for i := 0; i < N; i++ {
		// 1/2
		if rand.Float64() > 0.5 {
			continue
		}
		z.AddB(float64(i), fmt.Sprint(i))
	}

	testIsSorted(t, z)
	testInternalSpan(t, z)

	// Randomly add or delete value
	for i := 0; i < N; i++ {
		// 1/2
		if rand.Float64() > 0.5 {
			continue
		}
		z.Remove(fmt.Sprint(i))
	}
	testIsSorted(t, z)
	testInternalSpan(t, z)
}

func testIsSorted(t *testing.T, z *Set[string]) {
	var scores []float64
	for _, n := range z.Range(0, z.Len()-1) {
		scores = append(scores, n.Score)
	}
	assert.True(t, sort.Float64sAreSorted(scores))
}

func testInternalSpan(t *testing.T, z *Set[string]) {
	l := z.list
	for i := l.highestLevel - 1; i >= 0; i-- {
		x := l.header
		for x.loadNext(i) != nil {
			x = x.loadNext(i)
			span := x.loadSpan(i)
			from := x.value
			fromScore := x.score
			fromRank := l.Rank(fromScore, from)
			assert.NotEqual(t, -1, fromRank)

			if x.loadNext(i) != nil { // from -> to
				to := x.loadNext(i).value
				toScore := x.loadNext(i).score
				toRank := l.Rank(toScore, to)
				assert.NotEqual(t, -1, toRank)

				// span = to.rank - from.rank
				assert.Equalf(t, span, toRank-fromRank, "from %q (score: , rank: %d) to %q (score: %d, rank: %d), expect span: %d, actual: %d",
					from, fromScore, fromRank, to, toScore, toRank, span, toRank-fromRank)
			} else { // from -> nil
				// span = skiplist.len - from.rank
				assert.Equalf(t, l.length-fromRank, x.loadSpan(i), "%q (score: , rank: %d)", from, fromScore, fromRank)
			}
		}
	}
}

func TestSetRange(t *testing.T) {
	testFloat64SetRange(t, false)
}

func TestSetRevRange(t *testing.T) {
	testFloat64SetRange(t, true)
}

func testFloat64SetRange(t *testing.T, rev bool) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}

	start, stop := func(a, b int) (int, int) {
		if a < b {
			return a, b
		} else {
			return b, a
		}
	}(fastrand.Intn(N), fastrand.Intn(N))
	var ns []Node[string]
	if rev {
		ns = z.RevRange(start, stop)
	} else {
		ns = z.Range(start, stop)
	}
	assert.Equal(t, stop-start+1, len(ns))
	for i, n := range ns {
		if rev {
			assert.Equal(t, z.Len()-1-(start+i), z.Rank(n.Value))
		} else {
			assert.Equal(t, start+i, z.Rank(n.Value))
		}
	}
}

func TestSetRange_Negative(t *testing.T) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}
	ns := z.Range(-1, -1)
	assert.Equal(t, 1, len(ns))
	assert.Equal(t, z.Len()-1, z.Rank(ns[0].Value))
}

func TestSetRevRange_Negative(t *testing.T) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}
	ns := z.RevRange(-1, -1)
	assert.Equal(t, 1, len(ns))
	assert.Equal(t, 0, z.Rank(ns[0].Value))
}

func TestSetRangeByScore(t *testing.T) {
	testFloat64SetRangeByScore(t, false)
}

func TestSetRangeByScoreWithOpt(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	z.AddB(1.0, "1")
	z.AddB(1.1, "2")
	z.AddB(2.0, "3")

	ns := z.RangeByScoreWithOpt(1.0, 2.0, RangeOpt{ExcludeMin: true})
	assert.Equal(t, 2, len(ns))
	assert.Equal(t, 1.1, ns[0].Score)
	assert.Equal(t, 2.0, ns[1].Score)

	ns = z.RangeByScoreWithOpt(1.0, 2.0, RangeOpt{ExcludeMin: true, ExcludeMax: true})
	assert.Equal(t, 1, len(ns))
	assert.Equal(t, 1.1, ns[0].Score)

	ns = z.RangeByScoreWithOpt(1.0, 2.0, RangeOpt{ExcludeMax: true})
	assert.Equal(t, 2, len(ns))
	assert.Equal(t, 1.0, ns[0].Score)
	assert.Equal(t, 1.1, ns[1].Score)

	ns = z.RangeByScoreWithOpt(2.0, 1.0, RangeOpt{})
	assert.Equal(t, 0, len(ns))
	ns = z.RangeByScoreWithOpt(2.0, 1.0, RangeOpt{ExcludeMin: true})
	assert.Equal(t, 0, len(ns))
	ns = z.RangeByScoreWithOpt(2.0, 1.0, RangeOpt{ExcludeMax: true})
	assert.Equal(t, 0, len(ns))

	ns = z.RangeByScoreWithOpt(1.0, 1.0, RangeOpt{ExcludeMax: true})
	assert.Equal(t, 0, len(ns))
	ns = z.RangeByScoreWithOpt(1.0, 1.0, RangeOpt{ExcludeMin: true})
	assert.Equal(t, 0, len(ns))
	ns = z.RangeByScoreWithOpt(1.0, 1.0, RangeOpt{})
	assert.Equal(t, 1, len(ns))
}

func TestSetRevRangeByScoreWithOpt(t *testing.T) {
	z := New[string](bcomparator.StringComparator())
	z.AddB(1.0, "1")
	z.AddB(1.1, "2")
	z.AddB(2.0, "3")

	ns := z.RevRangeByScoreWithOpt(2.0, 1.0, RangeOpt{ExcludeMax: true})
	assert.Equal(t, 2, len(ns))
	assert.Equal(t, 1.1, ns[0].Score)
	assert.Equal(t, 1.0, ns[1].Score)

	ns = z.RevRangeByScoreWithOpt(2.0, 1.0, RangeOpt{ExcludeMax: true, ExcludeMin: true})
	assert.Equal(t, 1, len(ns))
	assert.Equal(t, 1.1, ns[0].Score)

	ns = z.RevRangeByScoreWithOpt(2.0, 1.0, RangeOpt{ExcludeMin: true})
	assert.Equal(t, 2, len(ns))
	assert.Equal(t, 2.0, ns[0].Score)
	assert.Equal(t, 1.1, ns[1].Score)

	ns = z.RevRangeByScoreWithOpt(1.0, 2.0, RangeOpt{})
	assert.Equal(t, 0, len(ns))
	ns = z.RevRangeByScoreWithOpt(1.0, 2.0, RangeOpt{ExcludeMin: true})
	assert.Equal(t, 0, len(ns))
	ns = z.RevRangeByScoreWithOpt(1.0, 2.0, RangeOpt{ExcludeMax: true})
	assert.Equal(t, 0, len(ns))

	ns = z.RevRangeByScoreWithOpt(1.0, 1.0, RangeOpt{ExcludeMax: true})
	assert.Equal(t, 0, len(ns))
	ns = z.RevRangeByScoreWithOpt(1.0, 1.0, RangeOpt{ExcludeMin: true})
	assert.Equal(t, 0, len(ns))
	ns = z.RevRangeByScoreWithOpt(1.0, 1.0, RangeOpt{})
	assert.Equal(t, 1, len(ns))
}

func TestSetRevRangeByScore(t *testing.T) {
	testFloat64SetRangeByScore(t, true)
}

func testFloat64SetRangeByScore(t *testing.T, rev bool) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}

	min, max := func(a, b float64) (float64, float64) {
		if a < b {
			return a, b
		} else {
			return b, a
		}
	}(fastrand.Float64(), fastrand.Float64())

	var ns []Node[string]
	if rev {
		ns = z.RevRangeByScore(max, min)
	} else {
		ns = z.RangeByScore(min, max)
	}
	var prev *float64
	for _, n := range ns {
		assert.LessOrEqual(t, min, n.Score)
		assert.GreaterOrEqual(t, max, n.Score)
		if prev != nil {
			if rev {
				assert.True(t, *prev >= n.Score)
			} else {
				assert.True(t, *prev <= n.Score)
			}
		}
		prev = &n.Score
	}
}

func TestSetCountWithOpt(t *testing.T) {
	testFloat64SetCountWithOpt(t, RangeOpt{})
	testFloat64SetCountWithOpt(t, RangeOpt{true, true})
	testFloat64SetCountWithOpt(t, RangeOpt{true, false})
	testFloat64SetCountWithOpt(t, RangeOpt{false, true})
}

func testFloat64SetCountWithOpt(t *testing.T, opt RangeOpt) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}

	min, max := func(a, b float64) (float64, float64) {
		if a < b {
			return a, b
		} else {
			return b, a
		}
	}(fastrand.Float64(), fastrand.Float64())

	n := z.CountWithOpt(min, max, opt)
	actualN := 0
	for _, n := range z.Range(0, -1) {
		if opt.ExcludeMin {
			if n.Score <= min {
				continue
			}
		} else {
			if n.Score < min {
				continue
			}
		}
		if opt.ExcludeMax {
			if n.Score >= max {
				continue
			}
		} else {
			if n.Score > max {
				continue
			}
		}
		actualN++
	}
	assert.Equal(t, actualN, n)
}

func TestSetRemoveRangeByRank(t *testing.T) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}

	start, stop := func(a, b int) (int, int) {
		if a < b {
			return a, b
		} else {
			return b, a
		}
	}(fastrand.Intn(N), fastrand.Intn(N))

	expectNs := z.Range(start, stop)
	actualNs := z.RemoveRangeByRank(start, stop)
	assert.Equal(t, expectNs, actualNs)

	// test whether removed
	for _, n := range actualNs {
		assert.False(t, z.Contains(n.Value))
	}
	assert.Equal(t, N, z.Len()+len(actualNs))
}

func TestSetRemoveRangeByScoreWithOpt(t *testing.T) {
	testFloat64SetRemoveRangeByScoreWithOpt(t, RangeOpt{})
	testFloat64SetRemoveRangeByScoreWithOpt(t, RangeOpt{true, true})
	testFloat64SetRemoveRangeByScoreWithOpt(t, RangeOpt{true, false})
	testFloat64SetRemoveRangeByScoreWithOpt(t, RangeOpt{false, false})
}

func testFloat64SetRemoveRangeByScoreWithOpt(t *testing.T, opt RangeOpt) {
	const N = 1000
	z := New[string](bcomparator.StringComparator())
	for i := 0; i < N; i++ {
		z.AddB(fastrand.Float64(), fmt.Sprint(i))
	}

	min, max := func(a, b float64) (float64, float64) {
		if a < b {
			return a, b
		} else {
			return b, a
		}
	}(fastrand.Float64(), fastrand.Float64())

	expectNs := z.RangeByScoreWithOpt(min, max, opt)
	actualNs := z.RemoveRangeByScoreWithOpt(min, max, opt)
	assert.Equal(t, expectNs, actualNs)

	// test whether removed
	for _, n := range actualNs {
		assert.False(t, z.Contains(n.Value))
	}
	assert.Equal(t, N, z.Len()+len(actualNs))
}

func TestUnionFloat64(t *testing.T) {
	var zs []*Set[string]
	for i := 0; i < 10; i++ {
		z := New[string](bcomparator.StringComparator())
		for j := 0; j < 100; j++ {
			if fastrand.Float64() > 0.8 {
				z.AddB(fastrand.Float64(), fmt.Sprint(i))
			}
		}
		zs = append(zs, z)
	}
	z := Union(bcomparator.StringComparator(), zs...)
	for _, n := range z.Range(0, z.Len()-1) {
		var expectScore float64
		for i := 0; i < 10; i++ {
			s, _ := zs[i].Score(n.Value)
			expectScore += s
		}
		assert.Equal(t, expectScore, n.Score)
	}
}

func TestUnionFloat64_Empty(t *testing.T) {
	z := Union(bcomparator.StringComparator())
	assert.Zero(t, z.Len())
}

func TestInterFloat64(t *testing.T) {
	var zs []*Set[string]
	for i := 0; i < 10; i++ {
		z := New[string](bcomparator.StringComparator())
		for j := 0; j < 10; j++ {
			if fastrand.Float64() > 0.8 {
				z.AddB(fastrand.Float64(), fmt.Sprint(i))
			}
		}
		zs = append(zs, z)
	}
	z := Inter(bcomparator.StringComparator(), zs...)
	for _, n := range z.Range(0, z.Len()-1) {
		var expectScore float64
		for i := 0; i < 10; i++ {
			s, ok := zs[i].Score(n.Value)
			assert.True(t, ok)
			expectScore += s
		}
		assert.Equal(t, expectScore, n.Score)
	}
}

func TestInterFloat64_Empty(t *testing.T) {
	z := Inter(bcomparator.StringComparator())
	assert.Zero(t, z.Len())
}

func TestInterFloat64_Simple(t *testing.T) {
	z1 := New[string](bcomparator.StringComparator())
	z1.AddB(0, "1")
	z2 := New[string](bcomparator.StringComparator())
	z2.AddB(0, "1")
	z3 := New[string](bcomparator.StringComparator())
	z3.AddB(0, "2")

	z := Inter(bcomparator.StringComparator(), z1, z2, z3)
	assert.Zero(t, z.Len())
}
