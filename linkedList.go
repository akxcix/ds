package ds

import (
	"fmt"
)

// LinkedListNode is a singly linked list node holding an integer
type LinkedListNode struct {
	value Item
	next  *LinkedListNode
}

// LinkedList is a singly linked list. head points to the first node, tail to the last.
// Insertion: O(1)
// Search: O(n)
// Deletion: O(n)
type LinkedList struct {
	head *LinkedListNode
}

// Modifiers --------------------------------------------------------------------------------------

// Insert inserts a new element at HEAD
// Complexity: O(1)
func (ll *LinkedList) Insert(x Item) {
	newNode := &LinkedListNode{
		value: x,
		next:  ll.head,
	}
	ll.head = newNode
}

// Delete removes the first occurance of an element from a linked list
// Complexity: O(n)
func (ll *LinkedList) Delete(x Item) {
	current := ll.head
	// if the value is at head, set head to next element
	if current.value == x {
		ll.head = current.next
		return
	}
	// otherwise check if next element has same value. if it does, set current nodes pointer
	// to it's next value
	for current.next != nil {
		if current.next.value == x {
			current.next = current.next.next
			return
		}
		current = current.next
	}

}

// Lookup -----------------------------------------------------------------------------------------

// IsEmpty returns true if the linked list is empty i.e. Head is null
// Complexity: O(1)
func (ll *LinkedList) IsEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false

}

// Search looks for an element in a linked list and returns the pointer to the Node
// when found, nil otherwise
// Complexity: O(n)
func (ll *LinkedList) Search(x Item) *LinkedListNode {
	current := ll.head
	for (current != nil) && (current.value != x) {
		current = current.next
	}
	return current
}

// Traverse traverse through a linked list and adds all values in an integer slice
// Complexity: O(n)
func (ll *LinkedList) Traverse() []Item {
	values := make([]Item, 0)
	current := ll.head
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}

// Print prints the slice obtained by using traverse
// Complexity: O(n)
func (ll *LinkedList) Print() {
	slice := ll.Traverse()
	fmt.Println(slice)
}
