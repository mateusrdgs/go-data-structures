package tree

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
	} else if n.Data <= node.Data {
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
