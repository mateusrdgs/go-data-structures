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

// DeleteAtBeginning should remove the first node
func (l *List) DeleteAtBeginning() {
	if l.Head != nil {
		l.Head = l.Head.Next

		l.decreaseLength()
	}
}

// DeleteAtEnd should remove the last node
func (l *List) DeleteAtEnd() {
	if l.Head != nil {
		temp := l.Head

		for temp.Next.Next != nil {
			temp = temp.Next
		}
		temp.Next = nil

		l.decreaseLength()
	}
}

// DeleteAtPosition should delete a node into the given position
func (l *List) DeleteAtPosition(pos int) {
	if pos == 0 {
		l.DeleteAtBeginning()
	} else if pos <= l.Len {
		if l.Head != nil {
			temp := l.Head

			for i := 0; i < pos-2; i++ {
				temp = temp.Next
			}

			node := temp.Next
			temp.Next = node.Next

			l.decreaseLength()
		}
	} else {
		l.DeleteAtEnd()
	}
}

// ReverseIterative should reserve the current list using an iterative algorithm
func (l *List) ReverseIterative() {
	current := l.Head
	var next *Node = nil
	var prev *Node = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}

// Print should log every item in the list in a iterative way
func (l *List) Print() {
	temp := l.Head

	for temp != nil {
		fmt.Printf("%d -> ", temp.Data)
		temp = temp.Next
	}
	fmt.Println("")
}

// PrintRecursive should log every item in the list in a recursive way
func PrintRecursive(n *Node) {
	if n == nil {
		fmt.Println("")
		return
	}
	fmt.Printf("%d -> ", n.Data)
	PrintRecursive(n.Next)
}

// ReversePrintRecursive should log every item in the list in a recursive way
func ReversePrintRecursive(n *Node) {
	if n == nil {
		return
	}
	ReversePrintRecursive(n.Next)
	fmt.Printf("%d -> ", n.Data)
}

// PrintLength should log the list length
func (l *List) PrintLength() {
	fmt.Println(l.Len)
}

func (l *List) increaseLength() {
	l.Len = l.Len + 1
}

func (l *List) decreaseLength() {
	l.Len = l.Len - 1
}
