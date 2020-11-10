package linkedlist

import "fmt"

// LinkedList interface
type LinkedList interface {
	Length() int
	AddNode(value interface{})
	ToString() string
}

// Node A node in a linked list
type Node struct {
	Value interface{}
	Next  *Node
}

// SinglyLinkedList Singly linked list
type SinglyLinkedList struct {
	Head   *Node
	length int
}

// AddNode Adds a Node to the linked list. Returns a reference to the Node
func (ll SinglyLinkedList) AddNode(value interface{}) {
	// Make new node with value given
	newNode := &Node{Value: value, Next: nil}
	// If we don't have a head then add one
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		// Otherwise loop through till we find the tail and append to it
		n := ll.Head
		for n.Next != nil {
			n = n.Next
		}
		n.Next = newNode
	}
	ll.length += 1
}

// ToString creates a string of the linked list
func (ll SinglyLinkedList) ToString() string {
	n := ll.Head
	str := fmt.Sprintf("%v", n.Value)
	for n.Next != nil {
		n = n.Next
		str = fmt.Sprintf("%v -> %v", str, n.Value)
	}
	return str
}

// Length Returns the length of the linked list
func (ll SinglyLinkedList) Length() int {
	return ll.length
}
