package main

import (
	doublylinkedlist "github.com/mateusrdgs/go-data-structures/doubly-linked-list"
)

func main() {
	list := &doublylinkedlist.List{Len: 0}

	list.InsertAtTail(10)
	list.InsertAtTail(8)
	list.InsertAtTail(6)
	list.InsertAtTail(4)
	list.InsertAtTail(2)
	list.RemoveAtHead()
}
