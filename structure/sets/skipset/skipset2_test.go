package skipset

import (
	"strings"
	"testing"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
)

func TestSetNew(t *testing.T) {
	set := New[int](bcomparator.IntComparator())
	set.Add(2, 1)

	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetAdd(t *testing.T) {
	set := New[int](bcomparator.IntComparator())
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestSetContains(t *testing.T) {
	set := New[int](bcomparator.IntComparator())
	set.Add(3, 1, 2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.Contains(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetRemove(t *testing.T) {
	set := New[int](bcomparator.IntComparator())
	set.Add(3, 1, 2)
	set.Remove()
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	set.Remove(1)
	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	set.Remove(3)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if actualValue := set.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSetString(t *testing.T) {
	c := New[int](bcomparator.IntComparator())
	c.Add(1)
	if !strings.HasPrefix(c.String(), "SkipSet") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkContains[E int](b *testing.B, set *Set[E], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(E(n))
		}
	}
}

func benchmarkAdd[E int](b *testing.B, set *Set[E], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(E(n))
		}
	}
}

func benchmarkRemove[E int](b *testing.B, set *Set[E], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(E(n))
		}
	}
}

func BenchmarkSkipSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkSkipSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkSkipSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkSkipSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkSkipSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New[int](bcomparator.IntComparator())
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkSkipSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkSkipSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkSkipSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkSkipSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkSkipSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkSkipSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkSkipSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New[int](bcomparator.IntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}
