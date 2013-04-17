//queue.go
//FIFO Queue Structure for BreadthFirstSearch()

package graph

type Queue struct {
	// FIFO Queue structure containing a list of nodes and a size.
	items []Node
	size  int
}

func GenQueue(length int) *Queue {
	//Initializes a new Queue.
	return &Queue{
		items: make([]Node, length),
		size:  length,
	}
}

func (q *Queue) GetSize() int {
	//Returns the size of the queue.
	return q.size
}

func (q *Queue) Enqueue(node Node) {
	//Adds an item to the end of the queue.
	q.size++
	newq := GenQueue(q.size)
	copy(newq.items, q.items)
	q.items = newq.items
	q.items[q.size-1] = node
}

func (q *Queue) Dequeue() Node {
	//Returns an item from the first spot in the queue, deletes it, and 
	//reorganizes the remaining items without disrupting the order.
	if q.size == 0 {
		return Node{}
	}
	node := q.items[0]
	q.items = q.items[1:]
	q.size--
	return node
}
