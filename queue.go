package ds

import (
	"errors"
)

// Errors -----------------------------------------------------------------------------------------

var (
	// ErrEmptyQueue is an error that occurs when an impossible operation is done on a
	// queue such as dequeue on an empty queue
	ErrEmptyQueue = errors.New("operation performed on empty queue")
)

// Node -------------------------------------------------------------------------------------------

// QueueNode holds an Item
type QueueNode struct {
	value Item
	next  *QueueNode
}

// Main Data Structure ----------------------------------------------------------------------------

// Queue is a LIFO data structure with O(1) insertion time. It holds two pointers (head
// and tail) and an integer storing size
type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

// Modifiers --------------------------------------------------------------------------------------

// Enqueue adds an item to the tail of the list
// Complexity: O(1)
func (q *Queue) Enqueue(x Item) {
	if q.head == nil {
		newNode := &QueueNode{value: x}
		q.head = newNode
		q.tail = q.head
		q.size++
		return
	}
	newNode := &QueueNode{value: x}
	q.tail.next = newNode
	q.tail = newNode
	q.size++
}

// Dequeue returns the node at the head (removing it from queue). Returns error if
// queue is empty
// Complexity: O(1)
func (q *Queue) Dequeue() (*QueueNode, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	head := q.head
	q.head = q.head.next
	q.size--
	if q.IsEmpty() {
		q.tail = q.head
	}
	return head, nil
}

// Utilities --------------------------------------------------------------------------------------

// IsEmpty returns true if the queue is empty, false otherwise
// Complexity: O(1)
func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

// PeekHead returns the item at Head, error if queue is empty.
// Complexity: O(1)
func (q *Queue) PeekHead() (Item, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	return q.head.value, nil
}

// PeekTail returns item at Tail, error if queue is empty.
// Complexity: O(1)
func (q *Queue) PeekTail() (Item, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	return q.tail.value, nil
}
