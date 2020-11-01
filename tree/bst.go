package tree

import (
	"fmt"
	"math"
)

// Node provides typing to a Tree node
type Node struct {
	Data  int
	left  *Node
	right *Node
}

// BST provides typing to binary search tree data structure
type BST struct {
	Root *Node
}

type queue struct {
	Head *queueNode
	Tail *queueNode
}

type queueNode struct {
	Data *Node
	next *queueNode
}

// Insert should add a new value to the tree
func (t *BST) Insert(n int) {
	node := &Node{Data: n, left: nil, right: nil}

	if t.Root == nil {
		t.Root = node
	} else {
		t.Root.insertNode(node)
	}
}

func (n *Node) insertNode(node *Node) {
	if n == nil {
		n = node
	} else if node.Data <= n.Data {
		if n.left == nil {
			n.left = node
		} else {
			n.left.insertNode(node)
		}
	} else {
		if n.right == nil {
			n.right = node
		} else {
			n.right.insertNode(node)
		}
	}
}

// Search should find a node with the provided value
func (t *BST) Search(n int) *Node {
	if t.Root == nil {
		return nil
	}
	return t.Root.searchByValue(n)
}

func (n *Node) searchByValue(value int) *Node {
	if n.Data == value {
		return n
	} else if n.Data <= value {
		return n.left.searchByValue(value)
	}
	return n.right.searchByValue(value)
}

// SearchMin should find a node with the min value into the tree
func (t *BST) SearchMin() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.left.searchMinValue()
}

func (n *Node) searchMinValue() int {
	if n.left == nil {
		return n.Data
	}
	return n.left.searchMinValue()
}

// SearchMax should find a node with the max value into the tree
func (t *BST) SearchMax() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.right.searchMaxValue()
}

func (n *Node) searchMaxValue() int {
	if n.right == nil {
		return n.Data
	}
	return n.right.searchMinValue()
}

// FindHeight should find the height of the current tree
func (t *BST) FindHeight() int {
	if t.Root == nil {
		return 0
	}

	left := t.Root.left.findHeight()
	right := t.Root.right.findHeight()
	max := math.Max(float64(left), float64(right))

	return int(max) + 1
}

func (n *Node) findHeight() int {
	if n == nil {
		return -1
	}

	left := n.left.findHeight()
	right := n.right.findHeight()
	max := math.Max(float64(left), float64(right))

	return int(max) + 1
}

// LevelOrder should traverse the whole tree using level-order traversal strategy
func (t *BST) LevelOrder() {
	if t.Root == nil {
		return
	}

	q := &queue{}

	q.enqueue(t.Root)

	for !q.isEmpty() {
		node := q.front()
		fmt.Println(node.Data)

		if node.left != nil {
			q.enqueue(node.left)
		}

		if node.right != nil {
			q.enqueue(node.right)
		}

		q.dequeue()
	}
}

func (q *queue) enqueue(n *Node) {
	node := &queueNode{Data: n}

	if q.Head == nil && q.Tail == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.next = node
		q.Tail = node
	}
}

func (q *queue) dequeue() {
	node := q.Head

	if q.Head == nil {
		return
	} else if q.Head == q.Tail {
		q.Head = q.Head.next
		q.Tail = q.Tail.next
	} else {
		q.Head = q.Head.next
	}

	node.next = nil
}

func (q *queue) front() *Node {
	return q.Head.Data
}

func (q *queue) isEmpty() bool {
	return q.Head == nil && q.Tail == nil
}
