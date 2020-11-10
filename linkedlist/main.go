package main

import (
	"fmt"

	"./linkedlist"
)

func main() {
	var ll linkedlist.LinkedList = &linkedlist.SinglyLinkedList{}

	ll.AddItem(2)
	ll.AddItem("five")
	ll.AddItem(6)
	n, err := ll.GetItem(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
		fmt.Println(ll.ToString())
		fmt.Println(ll.Length())
	}

}
