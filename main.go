package main

import (
	"github.com/mateusrdgs/go-data-structures/tree"
)

func main() {
	t := tree.BST{}

	t.Insert(15)
	t.Insert(10)
	t.Insert(20)
	t.Insert(25)
	t.Insert(8)
	t.Insert(17)
	t.Insert(12)
	t.Insert(11)

	t.Delete(15)

	t.InOrder()

}
