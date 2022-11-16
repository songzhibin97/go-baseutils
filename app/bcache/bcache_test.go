package bcache

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/songzhibin97/go-baseutils/base/bcomparator"
)

type TestStruct struct {
	Num      int
	Children []*TestStruct
}

func TestCache(t *testing.T) {
	tc := New[string, any](bcomparator.StringComparator())

	a, found := tc.Get("a")
	if found || a != nil {
		t.Error("a exist:", a)
	}

	b, found := tc.Get("b")
	if found || b != nil {
		t.Error("b exist:", b)
	}

	c, found := tc.Get("c")
	if found || c != nil {
		t.Error("c exist::", c)
	}

	tc.Set("a", 1, DefaultExpire)
	tc.Set("b", "b", DefaultExpire)
	tc.Set("c", 3.5, DefaultExpire)

	v, found := tc.Get("a")
	if !found {
		t.Error("a not exist")
	}
	if v == nil {
		t.Error("a == nil")
	} else if vv := v.(int); vv+2 != 3 {
		t.Error("vv != 3", vv)
	}

	v, found = tc.Get("b")
	if !found {
		t.Error("b not exist")
	}
	if v == nil {
		t.Error("b == nil")
	} else if vv := v.(string); vv+"B" != "bB" {
		t.Error("bb != bB:", vv)
	}

	v, found = tc.Get("c")
	if !found {
		t.Error("c not exist")
	}
	if v == nil {
		t.Error("x for c is nil")
	} else if vv := v.(float64); vv+1.2 != 4.7 {
		t.Error("vv != 4,7:", vv)
	}
}

func TestCacheTimes(t *testing.T) {
	var found bool
	tc := New[string, int](bcomparator.StringComparator(), SetDefaultExpire[string, int](50*time.Millisecond), SetInternal[string, int](time.Millisecond))
	tc.Set("a", 1, DefaultExpire)
	tc.Set("b", 2, NoExpire)
	tc.Set("c", 3, 20*time.Millisecond)
	tc.Set("d", 4, 70*time.Millisecond)

	<-time.After(25 * time.Millisecond)
	_, found = tc.Get("c")
	if found {
		t.Error("Found c when it should have been automatically deleted")
	}

	<-time.After(30 * time.Millisecond)
	_, found = tc.Get("a")
	if found {
		t.Error("Found a when it should have been automatically deleted")
	}

	_, found = tc.Get("b")
	if !found {
		t.Error("Did not find b even though it was set to never expire")
	}

	_, found = tc.Get("d")
	if !found {
		t.Error("Did not find d even though it was set to expire later than the default")
	}

	<-time.After(20 * time.Millisecond)
	_, found = tc.Get("d")
	if found {
		t.Error("Found d when it should have been automatically deleted (later than the default)")
	}
}

func TestStorePointerToStruct(t *testing.T) {
	tc := New[string, any](bcomparator.StringComparator())
	tc.Set("foo", &TestStruct{Num: 1}, DefaultExpire)
	x, found := tc.Get("foo")
	if !found {
		t.Fatal("*TestStruct was not found for foo")
	}
	foo := x.(*TestStruct)
	foo.Num++

	y, found := tc.Get("foo")
	if !found {
		t.Fatal("*TestStruct was not found for foo (second time)")
	}
	bar := y.(*TestStruct)
	if bar.Num != 2 {
		t.Fatal("TestStruct.Num is not 2")
	}
}

func TestAdd(t *testing.T) {
	tc := New[string, string](bcomparator.StringComparator())
	ok := tc.SetIfAbsent("foo", "bar", DefaultExpire)
	if !ok {
		t.Error("Couldn't add foo even though it shouldn't exist")
	}
	ok = tc.SetIfAbsent("foo", "baz", DefaultExpire)
	if ok {
		t.Error("Successfully added another foo when it should have returned an error")
	}
}

func TestReplace(t *testing.T) {
	tc := New[string, string](bcomparator.StringComparator())
	ok := tc.Replace("foo", "bar", DefaultExpire)
	if ok {
		t.Error("Replaced foo when it shouldn't exist")
	}
	tc.Set("foo", "bar", DefaultExpire)
	ok = tc.Replace("foo", "bar", DefaultExpire)
	if !ok {
		t.Error("Couldn't replace existing key foo")
	}
}

func TestDelete(t *testing.T) {
	tc := New[string, string](bcomparator.StringComparator())
	tc.Set("foo", "bar", DefaultExpire)
	tc.Delete("foo")
	x, found := tc.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != "" {
		t.Error("x is not nil:", x)
	}
}

func TestItemCount(t *testing.T) {
	tc := New[string, string](bcomparator.StringComparator())
	tc.Set("foo", "1", DefaultExpire)
	tc.Set("bar", "2", DefaultExpire)
	tc.Set("baz", "3", DefaultExpire)
	if n := tc.Count(); n != 3 {
		t.Errorf("Item count is not 3: %d", n)
	}
}

func TestFlush(t *testing.T) {
	tc := New[string, string](bcomparator.StringComparator())
	tc.Set("foo", "bar", DefaultExpire)
	tc.Set("baz", "yes", DefaultExpire)
	tc.Clear()
	x, found := tc.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != "" {
		t.Error("x is not nil:", x)
	}
	x, found = tc.Get("baz")
	if found {
		t.Error("baz was found, but it should have been deleted")
	}
	if x != "" {
		t.Error("x is not nil:", x)
	}
}

func TestCacheSerialization(t *testing.T) {
	tc := New[string, any](bcomparator.StringComparator())
	testFillAndSerialize(t, tc)

	// Check if gob.Register behaves properly even after multiple gob.Register
	// on c.Items (many of which will be the same type)
	testFillAndSerialize(t, tc)
}

func testFillAndSerialize(t *testing.T, tc *BCache[string, any]) {
	tc.Set("a", "a", DefaultExpire)
	tc.Set("b", "b", DefaultExpire)
	tc.Set("c", "c", DefaultExpire)
	tc.Set("expired", "foo", 1*time.Millisecond)
	tc.Set("*struct", &TestStruct{Num: 1}, DefaultExpire)
	tc.Set("[]struct", []TestStruct{
		{Num: 2},
		{Num: 3},
	}, DefaultExpire)
	tc.Set("[]*struct", []*TestStruct{
		{Num: 4},
		{Num: 5},
	}, DefaultExpire)
	tc.Set("structuration", &TestStruct{
		Num: 42,
		Children: []*TestStruct{
			{Num: 6174},
			{Num: 4716},
		},
	}, DefaultExpire)

	data, err := tc.Marshal()
	if err != nil {
		t.Fatal("Couldn't save cache to fp:", err)
	}

	oc := New[string, any](bcomparator.StringComparator())
	err = oc.Load(data)
	if err != nil {
		t.Fatal("Couldn't load cache from fp:", err)
	}

	a, found := oc.Get("a")
	if !found {
		t.Error("a was not found")
	}
	if a.(string) != "a" {
		t.Error("a is not a")
	}

	b, found := oc.Get("b")
	if !found {
		t.Error("b was not found")
	}
	if b.(string) != "b" {
		t.Error("b is not b")
	}

	c, found := oc.Get("c")
	if !found {
		t.Error("c was not found")
	}
	if c.(string) != "c" {
		t.Error("c is not c")
	}

	<-time.After(5 * time.Millisecond)
	_, found = oc.Get("expired")
	if found {
		t.Error("expired was found")
	}

	s1, found := oc.Get("*struct")
	if !found {
		t.Error("*struct was not found")
	}
	if s1.(map[string]interface{})["Num"].(float64) != 1 {
		t.Error("*struct.Num is not 1")
	}

	s2, found := oc.Get("[]struct")
	if !found {
		t.Error("[]struct was not found")
	}
	s2r := s2.([]interface{})
	if len(s2r) != 2 {
		t.Error("Length of s2r is not 2")
	}
	if s2r[0].(map[string]interface{})["Num"].(float64) != 2 {
		t.Error("s2r[0].Num is not 2")
	}
	if s2r[1].(map[string]interface{})["Num"].(float64) != 3 {
		t.Error("s2r[1].Num is not 3")
	}

	s3, found := oc.Get("[]*struct")
	if !found {
		t.Error("[]*struct was not found")
	}
	s3r := s3.([]interface{})
	if len(s3r) != 2 {
		t.Error("Length of s3r is not 2")
	}
	if s3r[0].(map[string]interface{})["Num"].(float64) != 4 {
		t.Error("s3r[0].Num is not 4")
	}
	if s3r[1].(map[string]interface{})["Num"].(float64) != 5 {
		t.Error("s3r[1].Num is not 5")
	}

	s4, found := oc.Get("structuration")
	if !found {
		t.Error("structuration was not found")
	}
	s4r := s4.(map[string]interface{})
	if len(s4r["Children"].([]interface{})) != 2 {
		t.Error("Length of s4r.Children is not 2")
	}

	if s4r["Children"].([]interface{})[0].(map[string]interface{})["Num"].(float64) != 6174 {
		t.Error("s4r.Children[0].Num is not 6174")
	}
	if s4r["Children"].([]interface{})[1].(map[string]interface{})["Num"].(float64) != 4716 {
		t.Error("s4r.Children[1].Num is not 4716")
	}
}

func BenchmarkCacheGetExpiring(b *testing.B) {
	benchmarkCacheGet(b)
}

func BenchmarkCacheGetNotExpiring(b *testing.B) {
	benchmarkCacheGet(b)
}

func benchmarkCacheGet(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](bcomparator.StringComparator())
	tc.Set("foo", "bar", DefaultExpire)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Get("foo")
	}
}

func BenchmarkRWMutexMapGet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"]
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetStruct(b *testing.B) {
	b.StopTimer()
	s := struct{ name string }{name: "foo"}
	m := map[interface{}]string{
		s: "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m[s]
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetString(b *testing.B) {
	b.StopTimer()
	m := map[interface{}]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"]
		mu.RUnlock()
	}
}

func BenchmarkCacheGetConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b)
}

func BenchmarkCacheGetConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b)
}

func benchmarkCacheGetConcurrent(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](bcomparator.StringComparator())
	tc.Set("foo", "bar", DefaultExpire)
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				tc.Get("foo")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkRWMutexMapGetConcurrent(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				mu.RLock()
				_, _ = m["foo"]
				mu.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkCacheGetManyConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b)
}

func BenchmarkCacheGetManyConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b)
}

func benchmarkCacheGetManyConcurrent(b *testing.B) {
	// This is the same as BenchmarkCacheGetConcurrent, but its result
	// can be compared against BenchmarkShardedCacheGetManyConcurrent
	// in sharded_test.go.
	b.StopTimer()
	n := 10000
	tc := New[string, string](bcomparator.StringComparator())
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		k := "foo" + strconv.Itoa(i)
		keys[i] = k
		tc.Set(k, "bar", DefaultExpire)
	}
	each := b.N / n
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for _, v := range keys {
		go func(k string) {
			for j := 0; j < each; j++ {
				tc.Get(k)
			}
			wg.Done()
		}(v)
	}
	b.StartTimer()
	wg.Wait()
}

func BenchmarkCacheSetExpiring(b *testing.B) {
	benchmarkCacheSet(b)
}

func BenchmarkCacheSetNotExpiring(b *testing.B) {
	benchmarkCacheSet(b)
}

func benchmarkCacheSet(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](bcomparator.StringComparator())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("foo", "bar", DefaultExpire)
	}
}

func BenchmarkRWMutexMapSet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
	}
}

func BenchmarkCacheSetDelete(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](bcomparator.StringComparator(), SetCapture[string, string](nil))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("foo", "bar", DefaultExpire)
		tc.Delete("foo")
	}
}

func BenchmarkRWMutexMapSetDelete(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
		mu.Lock()
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkCacheSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](bcomparator.StringComparator())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Lock()
		tc.set("foo", "bar", DefaultExpire)
		tc.delete("foo")
		tc.Unlock()
	}
}

func BenchmarkRWMutexMapSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkDeleteExpiredLoop(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](bcomparator.StringComparator(), SetDefaultExpire[string, string](5*time.Minute), SetCapture[string, string](nil))
	for i := 0; i < 100000; i++ {
		tc.set(strconv.Itoa(i), "bar", DefaultExpire)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.deleteExpire()
	}
}

func TestGetWithExpiration(t *testing.T) {
	tc := New[string, any](bcomparator.StringComparator())

	a, expiration, found := tc.GetWithExpire("a")
	if found || a != nil || !expiration.IsZero() {
		t.Error("Getting A found value that shouldn't exist:", a)
	}

	b, expiration, found := tc.GetWithExpire("b")
	if found || b != nil || !expiration.IsZero() {
		t.Error("Getting B found value that shouldn't exist:", b)
	}

	c, expiration, found := tc.GetWithExpire("c")
	if found || c != nil || !expiration.IsZero() {
		t.Error("Getting C found value that shouldn't exist:", c)
	}

	tc.Set("a", 1, DefaultExpire)
	tc.Set("b", "b", DefaultExpire)
	tc.Set("c", 3.5, DefaultExpire)
	tc.Set("d", 1, NoExpire)
	tc.Set("e", 1, 50*time.Millisecond)

	x, expiration, found := tc.GetWithExpire("a")
	if !found {
		t.Error("a was not found while getting a2")
	}
	if x == nil {
		t.Error("x for a is nil")
	} else if a2 := x.(int); a2+2 != 3 {
		t.Error("a2 (which should be 1) plus 2 does not equal 3; value:", a2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for a is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpire("b")
	if !found {
		t.Error("b was not found while getting b2")
	}
	if x == nil {
		t.Error("x for b is nil")
	} else if b2 := x.(string); b2+"B" != "bB" {
		t.Error("b2 (which should be b) plus B does not equal bB; value:", b2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for b is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpire("c")
	if !found {
		t.Error("c was not found while getting c2")
	}
	if x == nil {
		t.Error("x for c is nil")
	} else if c2 := x.(float64); c2+1.2 != 4.7 {
		t.Error("c2 (which should be 3.5) plus 1.2 does not equal 4.7; value:", c2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for c is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpire("d")
	if !found {
		t.Error("d was not found while getting d2")
	}
	if x == nil {
		t.Error("x for d is nil")
	} else if d2 := x.(int); d2+2 != 3 {
		t.Error("d (which should be 1) plus 2 does not equal 3; value:", d2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for d is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpire("e")
	if !found {
		t.Error("e was not found while getting e2")
	}
	if x == nil {
		t.Error("x for e is nil")
	} else if e2 := x.(int); e2+2 != 3 {
		t.Error("e (which should be 1) plus 2 does not equal 3; value:", e2)
	}
	v, _ := tc.member.Get("e")
	if expiration.UnixNano() != v.Expire {
		t.Error("expiration for e is not the correct time")
	}
	if expiration.UnixNano() < time.Now().UnixNano() {
		t.Error("expiration for e is in the past")
	}
}
