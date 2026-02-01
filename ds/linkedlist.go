package ds

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Insert(value string) { // inserts element 
	n := &Node{data: value}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Size = 1
		return // first element 
	}

	l.Tail.Next = n
	l.Tail = n
	l.Size++
}

func (l *LinkedList) InsertAt(position int, value string) error {
	if position < 0 || position > l.Size {
		return errors.New("position out of range")
	}

	n := &Node{data: value}

	if l.Size == 0 {
		l.Head = n
		l.Tail = n
		l.Size = 1
		return nil // size is 0, list is empty 
	}

	if position == 0 { // head case
		n.Next = l.Head
		l.Head = n
		l.Size++
		return nil 
	}

	if position == l.Size { // tail case
		l.Tail.Next = n
		l.Tail = n
		l.Size++
		return nil
	}

	prev := l.Head
	for i := 1; i < position; i++ { // middle case
		prev = prev.Next
	}
	n.Next = prev.Next
	prev.Next = n
	l.Size++
	return nil
}

func (l *LinkedList) Remove(value string) error {
	if l.Size == 0 {
		return errors.New("list is empty")
	}

	if l.Head.data == value { // element is head, remove it 
		l.Head = l.Head.Next
		l.Size--
		if l.Size == 0 {
			l.Tail = nil
		}
		return nil
	}

	prev := l.Head
	cur := l.Head.Next

	for cur != nil { 
		if cur.data == value {
			prev.Next = cur.Next
			if cur == l.Tail {
				l.Tail = prev
			}
			l.Size--
			return nil
		}
		prev = cur
		cur = cur.Next
	}

	return errors.New("value not found")
}

func (l *LinkedList) RemoveAll(value string) error {
	if l.Size == 0 {
		return errors.New("list is empty") // no more elements in list 
	}

	for l.Head != nil && l.Head.data == value { // removes all head matches
		l.Head = l.Head.Next
		l.Size-- // decrement size 
	}

	if l.Head == nil { // list empty after removing 
		l.Tail = nil
		return nil
	}

	prev := l.Head // starts from head
	cur := l.Head.Next
	found := false 

	for cur != nil {
		if cur.data == value {
			found = true
			prev.Next = cur.Next
			l.Size-- // decrements size after finding a match, loop continues
			cur = prev.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}

	l.Tail = prev // update the tail

	if !found && l.Head.data != value {
		return errors.New("value not found") // no matches found 
	}
	return nil
}

func (l *LinkedList) RemoveAt(pos int) error { // 
	if l.Size == 0 {
		return errors.New("list is empty")
	}
	if pos < 0 || pos >= l.Size {
		return errors.New("position out of range") // invalid 
	}

	if pos == 0 {
		l.Head = l.Head.Next // remove head
		l.Size--
		if l.Size == 0 {
			l.Tail = nil
		}
		return nil
	}

	prev := l.Head
	for i := 1; i < pos; i++ {
		prev = prev.Next // traverses to position
	}

	if prev.Next == l.Tail { // removes tail
		prev.Next = nil 
		l.Tail = prev // updates tail
	} else {
		prev.Next = prev.Next.Next // removes mid element
	}
	l.Size--
	return nil
}

func (l *LinkedList) IsEmpty() bool { return l.Size == 0 } // checks if list is empty
func (l *LinkedList) GetSize() int  { return l.Size } // gets size of list

func (l *LinkedList) Reverse() {
	var prev *Node
	cur := l.Head
	l.Tail = l.Head // reverses tail and head order

	for cur != nil {
		next := cur.Next // stores next node 
		cur.Next = prev
		prev = cur
		cur = next 
	}

	l.Head = prev // updates head again
}

func (l *LinkedList) PrintList() {
	for n := l.Head; n != nil; n = n.Next {
		if n != l.Head { 
			fmt.Print(" -> ")
		}
		fmt.Print(n.data)
	}
	fmt.Println()
}
