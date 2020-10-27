package main

import (
	"github.com/mateusrdgs/go-data-structures/queue"
)

func main() {
	qe := &queue.Queue{
		Data: make([]int, 10, 10),
	}

	qe.Enqueue(1)
	qe.Enqueue(2)
	qe.Dequeue()
	qe.Enqueue(3)
	qe.Enqueue(4)
	qe.Dequeue()
	qe.Dequeue()
	qe.Enqueue(5)
	qe.Enqueue(6)
	qe.Enqueue(7)
	qe.Enqueue(8)
	qe.Enqueue(9)
	qe.Enqueue(10)
	qe.Enqueue(11)
	qe.Enqueue(12)
	qe.Enqueue(13)
}
