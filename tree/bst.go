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
/*
	Is a binary tree in which for each node,
	the value of all nodes in left subtree
	is lesser or equal the value of the root node,
	and the value of all nodes in right subtree is greater
*/
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
	} else if value <= n.Data {
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

// PreOrder should traverse the whole tree using a pre-order traversal strategy (<root> <left> <right>)
func (t *BST) PreOrder() {
	if t.Root == nil {
		return
	}

	fmt.Println(t.Root.Data)

	if t.Root.left != nil {
		t.Root.left.preOrder()
	}

	if t.Root.right != nil {
		t.Root.right.preOrder()
	}
}

func (n *Node) preOrder() {
	if n == nil {
		return
	}

	n.left.preOrder()

	fmt.Println(n.Data)

	n.right.preOrder()
}

// InOrder should traverse the whole tree using an in-order traversal strategy (<left> <root> <right>)
func (t *BST) InOrder() {
	if t.Root == nil {
		return
	}

	if t.Root.left != nil {
		t.Root.left.inOrder()
	}

	fmt.Println(t.Root.Data)

	if t.Root.right != nil {
		t.Root.right.inOrder()
	}
}

func (n *Node) inOrder() {
	if n == nil {
		return
	}

	n.left.inOrder()

	fmt.Println(n.Data)

	n.right.inOrder()
}

// PostOrder should traverse the whole tree using a post-order traversal strategy (<left> <right> <root>)
func (t *BST) PostOrder() {
	if t.Root == nil {
		return
	}

	if t.Root.left != nil {
		t.Root.left.postOrder()
	}

	if t.Root.right != nil {
		t.Root.right.postOrder()
	}

	fmt.Println(t.Root.Data)
}

func (n *Node) postOrder() {
	if n == nil {
		return
	}

	n.left.postOrder()

	n.right.postOrder()

	fmt.Println(n.Data)
}

// IsBinarySearchTree should traverse the whole tree and check if the current tree is a BST
func (t *BST) IsBinarySearchTree() bool {
	if t.Root == nil {
		return true
	}
	return t.Root.isBinarySearchTree(math.MinInt16, math.MaxInt16)
}

func (n *Node) isBinarySearchTree(minValue, maxValue int) bool {
	if n == nil {
		return true
	}

	return n.Data >= minValue && n.Data < maxValue &&
		n.left.isBinarySearchTree(minValue, n.left.Data) &&
		n.right.isBinarySearchTree(n.left.Data, maxValue)
}

// Delete should traverse the tree in order to delete the node with the specified value
func (t *BST) Delete(v int) {
	if t.Root == nil {
		return
	}
	t.Root = t.Root.delete(v)
}

func (n *Node) delete(v int) *Node {
	if n == nil {
		return nil
	} else if v < n.Data {
		n.left = n.left.delete(v)
	} else if v > n.Data {
		n.right = n.right.delete(v)
	} else {
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.left == nil {
			n = n.right
		} else if n.right == nil {
			n = n.left
		} else {
			min := n.right.searchMinValue()
			n.Data = min
			n.right = n.right.delete(min)
		}
	}
	return n
}

// GetSucessor should traverse the tree in order to get the sucessor of the given value
func (t *BST) GetSucessor(v int) *Node {

	// Search the node - O(h)
	node := t.Root.searchByValue(v)

	if node == nil {
		return nil
	}

	// Node has right subtree
	if node.right != nil {
		temp := node.right
		for temp.left != nil {
			temp = temp.left
		}
		return temp
	}

	// No right subtree
	var ancestor *Node = t.Root
	var sucessor *Node = nil

	for ancestor != node {
		if node.Data < ancestor.Data {
			sucessor = ancestor
			ancestor = ancestor.left
		} else {
			ancestor = ancestor.right
		}
	}

	return sucessor
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
