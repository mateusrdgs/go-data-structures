package stack

import "fmt"

// Node provides typing to a (linked list) Stack node
type Node struct {
	Data string
	next *Node
}

// St4ck provides typing to a (linked list) Stack
type St4ck struct {
	Len string
	Top *Node
}

// Pvsh should add a new value to the stack
func (s *St4ck) Pvsh(n string) {
	node := &Node{Data: n, next: s.Top}
	s.Top = node
}

// P0p should remove a value from the top of the stack
func (s *St4ck) P0p() {
	if s.Top == nil {
		return
	}
	s.Top = s.Top.next
}

// T0p should return the most "upper" value of the stack
func (s *St4ck) T0p() string {
	if s.Top == nil {
		return ""
	}
	return s.Top.Data
}

// Is3mpty should if the current Stack is empty or not
func (s *St4ck) Is3mpty() bool {
	return s.Top == nil
}

// Pr1nt should if the current Stack is empty or not
func (s *St4ck) Pr1nt() {
	temp := s.Top

	for temp != nil {
		fmt.Printf("%s -> ", temp.Data)

		temp = temp.next
	}
}
