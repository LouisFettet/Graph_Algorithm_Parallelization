// parallel-algo.go
// Implementation of a parallelized Breadth First Search

package graph

// Function ParallelBFS accepts a graph with a valid source and sink 
// node, and returns a valid path through the graph's flow network along with
// the flow of the found path using channels to speed up runtime.
func ParallelBFS(g *Graph, source Node, sink Node) (int, map[Node]Node) {
	// Create a map in which nodes have node keys corresponding to a
	// parent/source to child/destination relationship.  This will be the
	// path returned.
	nodelist := g.GetNodeList()
	length := len(nodelist)
	path := make(map[Node]Node, length)
	// Initialize a node which can be used to tell if a node has been 
	// discovered yet or not, and give every node that key to begin.
	notvisited := Node{-1, -1}
	for _, node := range nodelist {
		path[node] = notvisited
	}
	// Give the source a different key to ensure it is not rediscovered.
	path[source] = Node{-2, -2}
	// Initialize another map that records the capacity of a found path 
	// to a node.
	capmap := make(map[Node]int, length)
	// Set the source's flow to infinity; math.Inf() is a float64, so we 
	// just make it huge.
	capmap[source] = 1000000
	// Initialize a queue and enqueue the source node.
	q := GenQueue(0)
	q.Enqueue(source)
	// Loop until the queue is empty.
	for q.GetSize() > 0 {
		// Grab the first node in the queue and check all neighbours
		// until one is found where flow can be pushed.
		u := q.Dequeue()
		for _, v := range g.GetNeighbours(u) {
			// If there is available capacity and the neighbour
			// has not been visited yet...
			if v.Capacity-v.Weight > 0 && path[v.Neighbour_Node] == notvisited {
				// Path can proceed from u to v by pushing 
				// flow forward. Set u to be the parent of v.
				path[v.Neighbour_Node] = u
				// Take the minimum of the flow of u and 
				// and the available capacity of v.
				capmap[v.Neighbour_Node] = Min(capmap[u], v.Capacity-v.Weight)
				if v.Neighbour_Node != sink {
					// We have not reached the sink. We 
					// enqueue v.Neighbour_Node and 
					// continue.
					q.Enqueue(v.Neighbour_Node)
				} else {
					// We have reached the sink and we 
					// return.
					return capmap[sink], path
				}
				// Else if capacity and weight are both 
				// negative (residual connection), and the 
				// neighbour has not been visited
			} else if v.Capacity < 0 && v.Weight < 0 && path[v.Neighbour_Node] == notvisited {
				// Path can proceed from u to v by pushing 
				// flow backward. Set u to be the parent of
				// v.
				path[v.Neighbour_Node] = u
				// Take the minimum of the flow of u and 
				// and the available capacity of v.
				capmap[v.Neighbour_Node] = Min(capmap[u], v.Weight-v.Capacity)
				if v.Neighbour_Node != sink {
					// We have not reached the sink.  
					// We enqueue v.Neighbour_Node 
					// and continue.
					q.Enqueue(v.Neighbour_Node)
				} else {
					// We have reached the sink, so we
					// return.
					return capmap[sink], path
				}
			}
		}
	}
	// No paths were found, so we return 0 and whatever path was built.
	return 0, path
}
