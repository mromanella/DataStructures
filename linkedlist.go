package datastructures

import "fmt"

// IndexError will be thrown when trying to index linkedlist and the index does not exist
type IndexError struct {
	message string
	index   int
}

func (ie *IndexError) Error() string {
	return ie.message
}

// Index The index that was given
func (ie *IndexError) Index() int {
	return ie.index
}

func outOfBounds(length, index int) bool {
	return index >= length || index < 0
}

// Node A Node in a linked list
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList Singly linked list
type LinkedList struct {
	head   *Node
	length int
}

// Get Gets the item at index. Throws an error if the index
// doesn't exist
func (ll *LinkedList) Get(index int) (*Node, error) {
	var n *Node
	if outOfBounds(ll.length, index) {
		return n, &IndexError{message: fmt.Sprintf("Not a valid index: %d", index), index: index}
	}
	n = ll.head
	for i := 0; i < ll.length; i++ {
		if i == index {
			break
		}
		n = n.Next
	}
	return n, nil
}

// Add Adds a Node to the end of the linked list
func (ll *LinkedList) Add(value interface{}) {
	// Make new Node with value given
	newNode := &Node{Value: value, Next: nil}
	// If we don't have a head then add one
	if ll.head == nil {
		ll.head = newNode
	} else {
		// Otherwise loop through till we find the tail and append to it
		n := ll.head
		for n.Next != nil {
			n = n.Next
		}
		n.Next = newNode
	}
	ll.length++
}

// Insert Adds the value at the specified index
// Returns an error if the index does not exist
func (ll *LinkedList) Insert(value interface{}, index int) error {
	if outOfBounds(ll.length, index) {
		return &IndexError{message: fmt.Sprintf("Not a valid index: %d", index), index: index}
	}

	prevIndex := index - 1
	if prevIndex >= 0 {
		// We have to insert it between 2
		prev, _ := ll.Get(prevIndex)
		newNode := &Node{Value: value, Next: prev.Next}
		prev.Next = newNode
	} else {
		// New head
		h := ll.head
		newNode := &Node{Value: value, Next: h}
		ll.head = newNode
	}
	ll.length++
	return nil
}

// Remove Removes the Node at the index and returns it or throws an IndexError if the index is out of
// bounds
func (ll *LinkedList) Remove(index int) (*Node, error) {
	if outOfBounds(ll.length, index) {
		return nil, &IndexError{message: fmt.Sprintf("Not a valid index: %d", index), index: index}
	}

	var toRemove *Node
	prevIndex := index - 1
	if prevIndex >= 0 {
		// Disconnect at the previous index and point the previous index to what the node at index
		// pointed to
		prev, _ := ll.Get(prevIndex)
		toRemove = prev.Next
		prev.Next = toRemove.Next
	} else {
		// Removing head
		toRemove = ll.head
		ll.head = ll.head.Next
	}
	ll.length--
	return toRemove, nil
}

// Contains Checks that the list contains the value
func (ll *LinkedList) Contains(value interface{}) bool {
	n := ll.head
	for n.Next != nil {
		if n.Value == value {
			return true
		}
		n = n.Next
	}
	return false
}

// ToString creates a string representation of the linked list
func (ll *LinkedList) ToString() string {
	n := ll.head
	str := fmt.Sprintf("%v", n.Value)
	for n.Next != nil {
		n = n.Next
		str = fmt.Sprintf("%v -> %v", str, n.Value)
	}
	return str
}

// Length Returns the amount of items in the linked list
func (ll *LinkedList) Length() int {
	return ll.length
}
