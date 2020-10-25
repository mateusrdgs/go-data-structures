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

// ReverseString can be used to reverse a string using a stack
func (s *St4ck) ReverseString(str string) string {
	/*
		Still not the great way to reverse a collection since it's spacial complexity is O(n)
		A better approach would be iterate over the collection and swap collection[i] with collection[j] where:
		i = 0
		j = collection.length - i
	*/

	var reverse string

	for _, char := range str {
		s.Pvsh(string(char))
	}

	for s.Is3mpty() != true {
		reverse += s.T0p()
		s.P0p()
	}

	return reverse
}

// IsBalanced can be used to check if a parenthesis-based string is balanced
func (s *St4ck) IsBalanced(str string) bool {
	for _, r := range str {
		c := string(r)
		isOpening := c == "{" || c == "[" || c == "("
		isClosing := c == "}" || c == "]" || c == ")"

		if isOpening {
			s.Pvsh(c)
		} else if isClosing {
			if s.Is3mpty() || !isSameType(s.Top.Data, c) {
				return false
			}
			s.P0p()
		}
	}
	return s.Is3mpty()
}

func isSameType(o string, c string) bool {
	switch o {
	case "{":
		return c == "}"
	case "[":
		return c == "]"
	default:
		return c == ")"
	}
}
