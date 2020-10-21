package linkedlist

import "fmt"

// Node provides typing to a linked list node
type Node struct {
	Data int
	Next *Node
}

// List provides typing to a linked list
type List struct {
	Len  int
	Head *Node
}

// InsertAtBeginning should insert a new node to the beginning of the list
func (l *List) InsertAtBeginning(n int) {
	node := &Node{Data: n}
	node.Next = l.Head
	l.Head = node

	l.increaseLength()
}

// InsertAtEnd should insert a new node to the end of the list
func (l *List) InsertAtEnd(n int) {
	node := &Node{Data: n}

	if l.Head == nil {
		l.Head = node
	} else {
		temp := l.Head

		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = node
	}

	l.increaseLength()
}

// InsertAtPosition should insert a new node into the given position
func (l *List) InsertAtPosition(n int, pos int) {
	if pos == 0 {
		l.InsertAtBeginning(n)
	} else if pos <= l.Len {
		node := &Node{Data: n}
		temp := l.Head

		for i := 0; i <= pos-2; i++ {
			temp = temp.Next
		}
		node.Next = temp.Next
		temp.Next = node

		l.increaseLength()
	} else {
		l.InsertAtEnd(n)
	}
}

// Print should log every item in the list
func (l *List) Print() {
	temp := l.Head

	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}

func (l *List) increaseLength() {
	l.Len = l.Len + 1
}
