package main

import (
	"fmt"

	"github.com/mateusrdgs/go-data-structures/stack"
)

func main() {
	stk := &stack.St4ck{}
	reverse := stk.ReverseString("mateus")

	fmt.Println(reverse)
}
