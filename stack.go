package ds

import (
	"errors"
	"fmt"
)

// Errors -----------------------------------------------------------------------------------------

var (
	// ErrEmptyStack is an error that occurs when an impossible operation is done on an empty stack
	// such as peeking top value or popping on an empty stack
	ErrEmptyStack = errors.New("operation performed on an empty stack")
)

// Node -------------------------------------------------------------------------------------------

// StackNode holds an Item
type StackNode struct {
	value Item
	prev  *StackNode
}

// Main Data Structure ----------------------------------------------------------------------------

// Stack is a LIFO data structure with O(1) insertion time. It holds pointer to the top
// node and an int storing size
type Stack struct {
	top  *StackNode
	size int
}

// Modifiers --------------------------------------------------------------------------------------

// Push adds an item at top of stack
// Complexity: O(1)
func (s *Stack) Push(x Item) {
	if s.top == nil {
		newNode := &StackNode{value: x}
		s.top = newNode
		s.size++
		return
	}
	newNode := &StackNode{value: x, prev: s.top}
	s.top = newNode
	s.size++
}

// Pop returns the node at the top of stack. Returns error if stack is empty
// Complexity: O(1)
func (s *Stack) Pop() (*StackNode, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	top := s.top
	s.top = s.top.prev
	s.size--
	return top, nil
}

// Utilities --------------------------------------------------------------------------------------

// IsEmpty returns a boolean which is true when the stack is empty
// Complexity: O(1)
func (s *Stack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

// Peek returns the value at the top of stack, error if stack is empty
// Complexity: O(1)
func (s *Stack) Peek() (Item, error) {
	if s.size == 0 {
		return nil, ErrEmptyStack
	}
	return s.top.value, nil
}

// Size returns the size of the stack
// Complexity: O(1)
func (s *Stack) Size() int {
	return s.size
}

// Traverse returns a slice of the elements from top to bottom
// Complexity: O(n)
func (s *Stack) Traverse() []Item {
	stack := make([]Item, 0)
	current := s.top
	for current != nil {
		stack = append(stack, current.value)
		current = current.prev
	}
	return stack
}

// Print displays the slice obtained by Stack.Traverse()
// Complexity: O(n)
func (s *Stack) Print() {
	fmt.Println(s.Traverse())
}
