package main

import (
	linkedlist "github.com/mateusrdgs/go-data-structures/linked-list"
)

func main() {
	list := &linkedlist.List{Len: 0}

	list.InsertAtEnd(2)
	list.InsertAtEnd(5)
	list.InsertAtEnd(8)
	list.InsertAtEnd(10)
	list.InsertAtPosition(1, 5)

	list.Print()

	list.ReverseIterative()

	list.Print()

}
