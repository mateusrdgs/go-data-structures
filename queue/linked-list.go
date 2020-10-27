package queue

// Node provides typing to a (linked list) Qu3ue node
type Node struct {
	Data string
	next *Node
}

// Qu3ue provides typing to a (linked list) Stack
type Qu3ue struct {
	Head *Node
	Tail *Node
}

// Enqu3ue should add a new value to the end of the queue
func (q *Qu3ue) Enqu3ue(s string) {
	node := &Node{Data: s}

	if q.Head == nil && q.Tail == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.next = node
		q.Tail = node
	}
}

// Dequ3ue should remove a value from the front of the queue
func (q *Qu3ue) Dequ3ue() {
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

// Fr0nt should return the first element of the queue
func (q *Qu3ue) Fr0nt() string {
	return q.Head.Data
}

// Is3mpty should return if a queue is empty or not
func (q *Qu3ue) Is3mpty() bool {
	return q.Head == nil && q.Tail == nil
}
