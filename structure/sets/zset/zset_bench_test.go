package zset

import (
	"math"
	"strconv"
	"testing"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/songzhibin97/go-baseutils/sys/fastrand"
)

const initSize = 1 << 10

const randN = math.MaxUint32

func BenchmarkContains100Hits(b *testing.B) {
	benchmarkContainsNHits(b, 100)
}

func BenchmarkContains50Hits(b *testing.B) {
	benchmarkContainsNHits(b, 50)
}

func BenchmarkContainsNoHits(b *testing.B) {
	benchmarkContainsNHits(b, 0)
}

func benchmarkContainsNHits(b *testing.B, n int) {
	b.Run("sortedset", func(b *testing.B) {
		z := New[string](bcomparator.StringComparator())
		var vals []string
		for i := 0; i < initSize; i++ {
			val := strconv.Itoa(i)
			vals = append(vals, val)
			if fastrand.Intn(100)+1 <= n {
				z.Add(fastrand.Float64(), val)
			}
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = z.Contains(vals[fastrand.Intn(initSize)])
			}
		})
	})
}

func BenchmarkAdd(b *testing.B) {
	benchmarkNAddNIncrNRemoveNContains(b, 100, 0, 0, 0)
}

func Benchmark1Add99Contains(b *testing.B) {
	benchmarkNAddNIncrNRemoveNContains(b, 1, 0, 0, 99)
}

func Benchmark10Add90Contains(b *testing.B) {
	benchmarkNAddNIncrNRemoveNContains(b, 10, 0, 0, 90)
}

func Benchmark50Add50Contains(b *testing.B) {
	benchmarkNAddNIncrNRemoveNContains(b, 50, 0, 0, 50)
}

func Benchmark1Add3Incr6Remove90Contains(b *testing.B) {
	benchmarkNAddNIncrNRemoveNContains(b, 1, 3, 6, 90)
}

func benchmarkNAddNIncrNRemoveNContains(b *testing.B, nAdd, nIncr, nRemove, nContains int) {
	// var anAdd, anIncr, anRemove, anContains int

	b.Run("sortedset", func(b *testing.B) {
		z := New[string](bcomparator.StringComparator())
		var vals []string
		var scores []float64
		var ops []int
		for i := 0; i < initSize; i++ {
			vals = append(vals, strconv.Itoa(fastrand.Intn(randN)))
			scores = append(scores, fastrand.Float64())
			ops = append(ops, fastrand.Intn(100))
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				r := fastrand.Intn(initSize)
				val := vals[r]
				if u := ops[r] + 1; u <= nAdd {
					// anAdd++
					z.Add(scores[r], val)
				} else if u-nAdd <= nIncr {
					// anIncr++
					z.IncrBy(scores[r], val)
				} else if u-nAdd-nIncr <= nRemove {
					// anRemove++
					z.Remove(val)
				} else if u-nAdd-nIncr-nRemove <= nContains {
					// anContains++
					z.Contains(val)
				}
			}
		})
		// b.Logf("N: %d, Add: %f, Incr: %f, Remove: %f, Contains: %f", b.N, float64(anAdd)/float64(b.N), float64(anIncr)/float64(b.N), float64(anRemove)/float64(b.N), float64(anContains)/float64(b.N))
	})
}
