package main

import (
	"fmt"

	"./linkedlist"
)

func main() {
	var ll linkedlist.LinkedList = &linkedlist.SinglyLinkedList{}

	ll.AddNode(2)
	ll.AddNode(5)
	ll.AddNode(6)

	fmt.Println(ll.ToString())
	// fmt.Println(ll.Length())

}
