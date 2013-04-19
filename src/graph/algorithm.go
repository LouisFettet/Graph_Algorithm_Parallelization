// algorithm.go
// Implementation of the Edmonds-Karp Algorithm for computing the maximum
// flow in a graph's flow network.

// Package graph provides both primitives for initializing graph structures
// and functions to solve the maximum flow of a graph.  Testing methods and 
// graphic interfaces are also available. 
package graph

// Function EdmondsKarp accepts a graph with a valid source and sink node,
// and returns the maximum flow as an integer along with a graph with a valid
// flow network. Utilizes BFS in order to find valid paths through the graph.
func EdmondsKarp(g *Graph, source Node, sink Node) (int, *Graph) {
	// Initialize the maximum flow.
	maxflow := 0
	// Loop until the BFS cannot return a valid path.
	for {
		pathcap, path := BreadthFirstSearch(g, source, sink)
		if pathcap == 0 {
			break
		}
		// Add the flow of the path found by the BFS to the maxflow.
		maxflow = maxflow + pathcap
		// Backtrack search through the graph and save the new flow.
		node := sink
		for node != source {
			// Grab the parent of the node.
			parent := path[node]
			// Grab the current weight of the connection.
			currweight := g.GetWeight(parent, node)
			// And add the new flow to the current weight.
			// Note: UpdateWeight also updates the residual path.
			g.UpdateWeight(parent, node, (currweight + pathcap))
			// Set the node equal to the parent so the backtrack
			// search can continue.
			node = parent
		}
	}
	// The BFS couldn't return a valid path, so we return.
	return maxflow, g
}

// Function BreadthFirstSearch accepts a graph with a valid source and sink 
// node, and returns a valid path through the graph's flow network along with
// the flow of the found path.
func BreadthFirstSearch(g *Graph, source Node, sink Node) (int, map[Node]Node) {
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
	//Set to infinity; math.Inf() is a float64, so we just make it huge.
	capmap[source] = 100000
	q := GenQueue(0)
	q.Enqueue(source)
	for q.GetSize() > 0 {
		u := q.Dequeue()
		for _, v := range g.GetNeighbours(u) {
			if v.capacity-v.weight > 0 && path[v.neighbour_node] == notvisited {
				// Path can proceed from u to v by pushing 
				// flow forward.
				path[v.neighbour_node] = u
				capmap[v.neighbour_node] = Min(capmap[u], v.capacity-v.weight)
				if v.neighbour_node != sink {
					// We have not reached the sink. We 
					// enqueue v.neighbour_node and 
					// continue.
					q.Enqueue(v.neighbour_node)
				} else {
					// We have reached the sink and we 
					// return.
					return capmap[sink], path
				}
			} else if v.capacity < 0 && v.weight < 0 && path[v.neighbour_node] == notvisited {
				// Path can proceed from u to v by pushing 
				// flow backward.
				path[v.neighbour_node] = u
				capmap[v.neighbour_node] = Min(capmap[u], v.weight-v.capacity)
				if v.neighbour_node != sink {
					// We have not reached the sink.  
					// We enqueue v.neighbour_node 
					// and continue.
					q.Enqueue(v.neighbour_node)
				} else {
					// We have reached the sink, so we
					// return.
					return capmap[sink], path
				}
			}
		}
	}
	return 0, path
}
