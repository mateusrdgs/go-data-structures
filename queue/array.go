package queue

// Queue provides typing to stack data structure
type Queue struct {
	Data []int
}

var head int = -1
var tail int = -1

// Enqueue should add a new value to the end of the queue
func (q *Queue) Enqueue(n int) {
	if q.isFull() {
		return
	} else if q.IsEmpty() {
		tail = 0
		head = tail
	} else {
		tail = (tail + 1) % len(q.Data)
	}
	q.Data[tail] = n
}

// Dequeue should remove a value from the front of the queue
func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		return
	} else if head == tail {
		tail = -1
		head = tail
	} else {
		head = (head + 1) % len(q.Data)
	}
}

// IsEmpty should return if the current queue is empty
func (q *Queue) IsEmpty() bool {
	return head == -1 && tail == -1
}

// Front should return the element at the front of the queue
func (q *Queue) Front() int {
	return q.Data[head]
}

func (q *Queue) isFull() bool {
	return ((tail + 1) % len(q.Data)) == head
}
