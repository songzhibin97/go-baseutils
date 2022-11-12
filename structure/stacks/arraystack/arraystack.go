// Package arraystack implements a stack backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
package arraystack

import (
	"fmt"
	"strings"

	"github.com/songzhibin97/go-baseutils/structure/lists/arraylist"
	"github.com/songzhibin97/go-baseutils/structure/stacks"
)

// Assert Stack implementation
var _ stacks.Stack[any] = (*Stack[any])(nil)

// Stack holds elements in an array-list
type Stack[E any] struct {
	list *arraylist.List[E]
}

// New instantiates a new empty stack
func New[E any]() *Stack[E] {
	return &Stack[E]{list: arraylist.New[E]()}
}

// Push adds a value onto the top of the stack
func (stack *Stack[E]) Push(value E) {
	stack.list.Add(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack[E]) Pop() (value E, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack[E]) Peek() (value E, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

// Empty returns true if stack does not contain any elements.
func (stack *Stack[E]) Empty() bool {
	return stack.list.Empty()
}

// Size returns number of elements within the stack.
func (stack *Stack[E]) Size() int {
	return stack.list.Size()
}

// Clear removes all elements from the stack.
func (stack *Stack[E]) Clear() {
	stack.list.Clear()
}

// Values returns all elements in the stack (LIFO order).
func (stack *Stack[E]) Values() []E {
	size := stack.list.Size()
	elements := make([]E, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

// String returns a string representation of container
func (stack *Stack[E]) String() string {
	b := strings.Builder{}
	b.WriteString("ArrayStack\n")
	for index, value := range stack.list.Values() {
		b.WriteString(fmt.Sprintf("(index:%d value:%v) ", index, value))
	}
	return b.String()
}

// Check that the index is within bounds of the list
func (stack *Stack[E]) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}

// UnmarshalJSON @implements json.Unmarshaler
func (stack *Stack[E]) UnmarshalJSON(bytes []byte) error {
	return stack.list.UnmarshalJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (stack *Stack[E]) MarshalJSON() ([]byte, error) {
	return stack.list.MarshalJSON()
}
