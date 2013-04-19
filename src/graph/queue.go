// queue.go
// Generates a FIFO queue structure with accompanying functions for use by 
// BreadthFirstSearch in algorithm.go

// Package graph provides both primitives for initializing graph structures
// and functions to solve the maximum flow of a graph.  Testing methods and 
// graphic interfaces are also available.  
package graph

// FIFO queue structure containing a list of nodes and a size.
type Queue struct {
	items []Node
	size  int
}

// Function GenQueue initializes a new queue with a given length.
// Note: Use a length of 0 to begin; GenQueue is also used in function Enqueue
// to copy the contents and update the size. 
func GenQueue(length int) *Queue {
	return &Queue{
		items: make([]Node, length),
		size:  length,
	}
}

// Function GetSize returns the current size of the queue.
func (q *Queue) GetSize() int {
	return q.size
}

// Function Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(node Node) {
	// Increase the queue size.
	q.size++
	// Create a new queue with the new size and copy the elements over.
	newq := GenQueue(q.size)
	copy(newq.items, q.items)
	q.items = newq.items
	// Insert the node into the last spot of the queue.
	q.items[q.size-1] = node
}

// Function Dequeue returns an item from the first spot in the queue, deletes 
// it from the queue, and reorganizes the remaining items without disrupting
// the queue's order.
func (q *Queue) Dequeue() Node {
	// Return a nil node if the queue is empty.
	if q.size == 0 {
		return Node{}
	}
	node := q.items[0]
	// Slice off the first spot in the queue.
	q.items = q.items[1:]
	// Decrement the size of the queue.
	q.size--
	return node
}
