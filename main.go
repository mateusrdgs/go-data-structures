package main

import (
	"fmt"

	"github.com/mateusrdgs/go-data-structures/stack"
)

func main() {
	stk := &stack.St4ck{}

	fmt.Println(stk.IsBalanced("([])"))
}
