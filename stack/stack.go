package stack

// Stack provides typing to stack data structure
type Stack struct {
	Data []int
}

var top int = -1

// Push should add a new value to the stack
func (s *Stack) Push(n int) {
	if top == (cap(s.Data) - 1) {
		println("Stack overflow")
		return
	}
	top++
	s.Data[top] = n
}

// Pop should remove a value to the stack
func (s *Stack) Pop(n int) {
	if top == -1 {
		println("No element to pop")
		return
	}
	top--
}

// Top should return the most "upper" value of the stack
func (s *Stack) Top() int {
	if top > 0 {
		return s.Data[top]
	}
	println("No elements on stack")
	return -1
}
