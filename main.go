package main

import (
	"fmt"

	"github.com/mateusrdgs/go-data-structures/stack"
)

func main() {
	stk := &stack.St4ck{}

	fmt.Println(stk.EvaluatePostfix("23*54*+9-"))
	fmt.Println(stk.EvaluatePrefix("-+*23*549"))
}
