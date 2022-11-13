package linkedliststack

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := New[int]()
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if actualValue := stack.Values(); actualValue[0] != 3 || actualValue[1] != 2 || actualValue[2] != 1 {
		t.Errorf("Got %v expected %v", actualValue, "[3,2,1]")
	}
	if actualValue := stack.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := stack.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestStackPeek(t *testing.T) {
	stack := New[int]()
	if actualValue, ok := stack.Peek(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestStackPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	if actualValue, ok := stack.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := stack.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := stack.Pop(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, ok := stack.Pop(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := stack.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func TestStackIterator(t *testing.T) {
	stack := New[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	// Iterator
	it := stack.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	stack.Clear()
	it = stack.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty stack")
	}
}

func TestStackIteratorBegin(t *testing.T) {
	stack := New[string]()
	it := stack.Iterator()
	it.Begin()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "c")
	}
}

func TestStackIteratorFirst(t *testing.T) {
	stack := New[string]()
	it := stack.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "c")
	}
}

func TestStackIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		stack := New[string]()
		it := stack.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty stack")
		}
	}

	// NextTo (not found)
	{
		stack := New[string]()
		stack.Push("xx")
		stack.Push("yy")
		it := stack.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty stack")
		}
	}

	// NextTo (found)
	{
		stack := New[string]()
		stack.Push("aa")
		stack.Push("bb")
		stack.Push("cc")
		it := stack.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty stack")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 2 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "aa")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestStackSerialization(t *testing.T) {
	stack := New[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	var err error
	assert := func() {
		if actualValue, expectedValue := stack.Values(), []string{"c", "b", "a"}; !reflect.DeepEqual(actualValue, expectedValue) {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := stack.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := stack.MarshalJSON()
	assert()

	err = stack.UnmarshalJSON(bytes)
	assert()

	bytes, err = json.Marshal([]interface{}{"a", "b", "c", stack})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`["1","2","3"]`), &stack)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestStackString(t *testing.T) {
	c := New[int]()
	c.Push(1)
	if !strings.HasPrefix(c.String(), "LinkedListStack") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkPush[E int](b *testing.B, stack *Stack[E], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Push(E(n))
		}
	}
}

func benchmarkPop[E int](b *testing.B, stack *Stack[E], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Pop()
		}
	}
}

func BenchmarkLinkedListStackPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := New[int]()
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkLinkedListStackPush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkLinkedListStackPush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkLinkedListStackPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}
