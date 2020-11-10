package linkedlist

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

// LinkedList interface
type LinkedList interface {
	Length() int
	GetItem(index int) (*Node, error)
	AddItem(value interface{})
	InsertItem(value interface{}, index int) error
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

// GetItem Gets the item at index. Throws an error if the index
// doesn't exist
func (ll *SinglyLinkedList) GetItem(index int) (*Node, error) {
	var n *Node
	if index >= ll.length || index < 0 {
		return n, &IndexError{message: "Not a valid index", index: index}
	}
	n = ll.Head
	for i := 0; i < ll.length; i++ {
		if i == index {
			break
		}
		n = n.Next
	}
	return n, nil
}

// AddItem Adds a Node to the end of the linked list
func (ll *SinglyLinkedList) AddItem(value interface{}) {
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
	ll.length++
}

// InsertItem Adds the value at the specified index
// Returns an error if the index does not exist
func (ll *SinglyLinkedList) InsertItem(value interface{}, index int) error {
	if index >= ll.length || index < 0 {
		return &IndexError{message: "Not a valid index", index: index}
	}
	if index == 0 {
		newNode := &Node{Value: value, Next: ll.Head}
		ll.Head = newNode
	} else {
		i := 0
		n := ll.Head
		for n.Next != nil {
			if i == (index - 1) {
				newNode := &Node{Value: value, Next: n.Next}
				n.Next = newNode
				break
			}
			i++
		}
	}
	ll.length++
	return nil
}

// ToString creates a string of the linked list
func (ll *SinglyLinkedList) ToString() string {
	n := ll.Head
	str := fmt.Sprintf("%v", n.Value)
	for n.Next != nil {
		n = n.Next
		str = fmt.Sprintf("%v -> %v", str, n.Value)
	}
	str = fmt.Sprintf("%v -> %v", str, nil)
	return str
}

// Length Returns the amount of items in the linked list
func (ll *SinglyLinkedList) Length() int {
	return ll.length
}
