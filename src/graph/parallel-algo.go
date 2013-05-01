// parallel-algo.go
// Parallelized implementation of the Edmonds-Karp Algorithm for computing 
// the maximum flow in a graph's flow network.

package graph

// Initialize a number of CPU cores.
var NCPU = 4

// Set up a semaphore that will be used in the ParallelBFS...
type empty struct{}
type semaphore chan empty

// Function ParallelEdmondsKarp accepts a graph with a valid source and sink
// node, and returns the maximum flow as an integer along with a graph with a
// valid flow network. The function runs a parallelized BFS in order to find
// valid paths through the graph.
func ParallelEdmondsKarp(g *Graph, source Node, sink Node) (int, *Graph) {
	// Initialize the maximum flow.
	maxflow := 0
	// Loop until the BFS cannot return a valid path.
	for {
		pathcap, path := ParallelBFS(g, source, sink)
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

// Function ParallelBFS accepts a graph with a valid source and sink 
// node, and returns a valid path through the graph's flow network along with
// the flow of the found path using channels to speed up runtime.
func ParallelBFS(g *Graph, source Node, sink Node) (int, map[Node]Node) {

	// Create a map in which nodes have node keys corresponding to a
	// parent/source to child/destination relationship.  This will be the
	// path returned.
	nodelist := g.GetNodeList()
	nodecount := len(nodelist)
	path := make(map[Node]Node, nodecount)

	// Initialize channel with buffer size of cores used.
	c := make(chan int, NCPU)

	// Initialize a node which can be used to tell if a node has been 
	// discovered yet or not, and give every node that key to begin.
	notvisited := Node{-1, -1}
	for i := 0; i < NCPU; i++ {
		// We will set each node concurrently for slight speed-up...
		go func(i int) {
			head := i * nodecount / NCPU
			tail := (i + 1) * nodecount / NCPU
			for _, node := range nodelist[head:tail] {
				path[node] = notvisited
			}
			c <- 1
		}(i)
	}
	// Drain the channel...
	for i := 0; i < NCPU; i++ {
		<-c
	}

	// Give the source a different key to ensure it is not rediscovered.
	path[source] = Node{-2, -2}

	// Initialize another map that records the capacity of a found path 
	// to a node.
	capmap := make(map[Node]int, nodecount)

	// Set the source's flow to infinity; math.Inf() is a float64, so we 
	// just make it huge.
	capmap[source] = 1000000

	// Initialize a queue and enqueue the source node.
	q := GenQueue(0)
	q.Enqueue(source)

	// Loop until the queue is empty.
	for q.GetSize() > 0 {

		// Grab the first node in the queue...
		u := q.Dequeue()

		// And grab all of its neighbours.
		neighbours := g.GetNeighbours(u)

		// Initialize a slice for valid neighbours, where the path
		// can proceed from u to v.
		var valids []Neighbour

		// Validate the neighbours to see which ones to iterate over.
		for _, v := range neighbours {
			// If there is available capacity and the neighbour
			// has not been visited yet...
			if path[v.Neighbour_Node] == notvisited && (v.Capacity-v.Weight > 0 || (v.Weight < 0 && v.Capacity < 0)) {
				valids = append(valids, v)
			}
		}

		// Initialize the semaphore using the amount of valids.
		semlen := len(valids)
		sem := make(semaphore, semlen)

		// Concurrently iterate over each valid.
		for _, v := range valids {
			go func(u Node, v Neighbour) {
				// Set u to be the parent of v.
				path[v.Neighbour_Node] = u
				// Check to see whether the connection is 
				// residual and then take the minimum of the 
				// flow of u and and the available capacity 
				// of v.
				if v.Capacity > 0 {
					capmap[v.Neighbour_Node] = Min(capmap[u], v.Capacity-v.Weight)
				} else if v.Capacity < 0 {
					capmap[v.Neighbour_Node] = Min(capmap[u], v.Weight-v.Capacity)
				}
				switch {
				case v.Neighbour_Node != sink:
					// We have not reached the sink. We 
					// enqueue v.Neighbour_Node and 
					// continue.
					q.Enqueue(v.Neighbour_Node)
				case v.Neighbour_Node == sink:
					// We have reached the sink, so
					// we empty the queue, break, and 
					// return immediately.
					q = GenQueue(0)
					break
				}
				sem <- empty{}
			}(u, v)
		}
		// Drain the semaphore; wait for goroutines to finish.
		for i := 0; i < semlen; i++ {
			<-sem
		}
	}
	return capmap[sink], path
}
