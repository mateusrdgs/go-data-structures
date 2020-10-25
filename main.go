package main

import (
	"github.com/mateusrdgs/go-data-structures/stack"
)

func main() {
	stk := &stack.St4ck{}

	stk.Pvsh("2")
	stk.Pvsh("4")
	stk.Pvsh("6")

	stk.Pr1nt()

}
