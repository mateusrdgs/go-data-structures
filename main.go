package main

import (
	"fmt"

	"github.com/mateusrdgs/go-data-structures/stack"
)

func main() {
	stk := &stack.Stack{
		Data: make([]int, 5),
	}

	stk.Top()

	stk.Push(10)
	stk.Push(20)
	stk.Push(30)
	stk.Push(40)
	stk.Push(50)

	fmt.Println(stk.Top())

}
