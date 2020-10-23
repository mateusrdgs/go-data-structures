package doublylinkedlist

import (
	"fmt"
)

// Node provides typing to a linked list node
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

// List provides typing to a linked list
type List struct {
	Len  int
	Head *Node
}

// InsertAtHead should insert a new node to the beginning of the list
func (l *List) InsertAtHead(n int) {
	node := &Node{Data: n}

	if l.Head == nil {
		l.Head = node
	} else {
		l.Head.Prev = node
		node.Next = l.Head
		l.Head = node
	}

	l.increaseLength()
}

// InsertAtTail should insert a new node to the end of the list
func (l *List) InsertAtTail(n int) {
	node := &Node{Data: n}

	if l.Head == nil {
		l.Head = node
	} else {
		temp := l.Head

		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = node
		node.Prev = temp
	}

	l.increaseLength()
}

// RemoveAtHead should remove a node from the beginning of the list
func (l *List) RemoveAtHead() {

	if l.Head == nil {
		return
	}

	node := l.Head

	if node.Next != nil {
		l.Head = node.Next
		node.Next.Prev = nil
	} else {
		l.Head = nil
	}

	node = nil

	l.decreaseLength()
}

// RemoveAtTail should remove a node from the beginning of the list
func (l *List) RemoveAtTail() {

	if l.Head == nil {
		return
	}

	temp := l.Head

	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp.Next = nil

	l.decreaseLength()
}

// Print should log every item in the list in a iterative way
func (l *List) Print() {
	temp := l.Head

	for temp != nil {
		fmt.Printf("%d -> ", temp.Data)
		temp = temp.Next
	}
}

// ReversePrint should log every item in the list in a iterative way
func (l *List) ReversePrint() {
	temp := l.Head

	if temp == nil {
		return
	}

	for temp.Next != nil {
		temp = temp.Next
	}

	for temp != nil {
		fmt.Printf("%d -> ", temp.Data)
		temp = temp.Prev
	}
}

func (l *List) increaseLength() {
	l.Len = l.Len + 1
}

func (l *List) decreaseLength() {
	l.Len = l.Len - 1
}
