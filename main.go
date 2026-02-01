// Kobe Limon
// COSC 3750
// 1/31/2026
package main

import (
	"fmt"
	"linkedlist-KobeLimon21/ds"
)

func main() {
	linkedlist := &ds.LinkedList{}
	_ = linkedlist.InsertAt(0, "first")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")

	_ = linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	_ = linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")

	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	stack := &ds.Stack{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	data, ok := stack.Pop()
	fmt.Println("Popped from stack:", data, ok)

	queue := &ds.Queue{}
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")
	data2, err := queue.Pop()
	fmt.Println("Popped from queue:", data2, err)
}
