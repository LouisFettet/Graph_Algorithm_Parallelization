// parallel-algo.go
// Parallelized implementation of the Edmonds-Karp Algorithm for computing 
// the maximum flow in a graph's flow network.

package graph

import "time"

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
	length := len(nodelist)
	path := make(map[Node]Node, length)
	// Initialize the channels with function MakeChannels.  Four channels
	// should be able to run on most modern systems without any issues. 
	// Pass in the length so that a correct buffer size can be allocated. 
	// channel1, channel2, channel3, channel4 := MakeChannels(length)
	// Initialize a node which can be used to tell if a node has been 
	// discovered yet or not, and give every node that key to begin.
	notvisited := Node{-1, -1}
	go SetNotVisited(nodelist, path, notvisited, 1, length/4)
	//<-channel1
	go SetNotVisited(nodelist, path, notvisited, length/4, length/4*2)
	//<-channel2
	go SetNotVisited(nodelist, path, notvisited, length/4*2, length/4*3)
	//<-channel3
	go SetNotVisited(nodelist, path, notvisited, length/4*3, length)
	//<-channel4
	time.Sleep(time.Nanosecond)
	//fmt.Println(path)
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
	/*
		go FindPath(g, q, source, sink, path, capmap)
		fmt.Println(path[source])
	*/
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

/*
func MakeChannels(length int) (chan int, chan int, chan int, chan int) {
	channel1 := make(chan int, length*2)
	channel2 := make(chan int, length*2)
	channel3 := make(chan int, length*2)
	channel4 := make(chan int, length*2)
	return channel1, channel2, channel3, channel4
}
*/

func SetNotVisited(nodelist []Node, path map[Node]Node, notvisited Node, head int, tail int) {
	for _, node := range nodelist[head:tail] {
		path[node] = notvisited
		//		channel <- 0
	}
}

/*
func FindPath(g *Graph, q *Queue, source Node, sink Node, path map[Node]Node, capmap map[Node]int) {
	path[source] = Node{0, 0}
	fmt.Println("kk")
}
*/
