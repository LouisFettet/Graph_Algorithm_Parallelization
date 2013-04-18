//algorithm.go
//Implementation of the Edmonds-Karp Algorithm for Computing the Maximum Flow in a Flow Network

package graph

import "fmt"

func EdmondsKarp(g *Graph, source Node, sink Node) (int, *Graph) {
	fmt.Println("EdmondsKarp started.")
	maxflow := 0
	for {
		pathcap, path := BreadthFirstSearch(g, source, sink)
		fmt.Println("BFS returned a pathcap of  ", pathcap, "and a path of:\n", path)
		if pathcap == 0 {
			break
		}
		maxflow = maxflow + pathcap
		fmt.Println("The new maxflow is: ", maxflow)
		v := sink
		fmt.Println("Backtrack search begin.")
		for v != source {
			u := path[v]
			fmt.Println("v is ", v, "and u is ", u)
			currweight := g.GetWeight(u, v)
			fmt.Println("currweight for u, v is ", currweight)
			g.UpdateWeight(u, v, (currweight + pathcap))
			fmt.Println("The new weight is ", g.GetWeight(u, v))
			v = u
		}
	}
	fmt.Println("EdmondsKarp complete.")
	return maxflow, g
}

func BreadthFirstSearch(g *Graph, source Node, sink Node) (int, map[Node]Node) {
	fmt.Println("\nBFS started.")
	nodelist := g.GetNodeList()
	length := len(nodelist)
	path := make(map[Node]Node, length)
	notvisited := Node{-1, -1}
	for _, node := range nodelist {
		path[node] = notvisited
	}
	//Make sure source is not rediscovered.
	path[source] = Node{-2, -2}
	//Capacity of found path to node.
	capmap := make(map[Node]int, length)
	//Set to infinity; math.Inf() is a float64, so we just make it huge.
	capmap[source] = 100000
	q := GenQueue(0)
	q.Enqueue(source)
	fmt.Println("Everything initialized... while loop started.")
	for q.GetSize() > 0 {
		fmt.Println("New Iteration, Queue is: ", q, "and has size ", q.GetSize())
		u := q.Dequeue()
		fmt.Println("u = ", u)
		for _, v := range g.GetNeighbours(u) {
			fmt.Println("v.neighbour_node = ", v.neighbour_node)
			fmt.Println("v.weight = ", v.weight)
			fmt.Println("v.capacity = ", v.capacity)
			if v.capacity-v.weight > 0 && path[v.neighbour_node] == notvisited {
				fmt.Println("Path can proceed from u to v by pushing flow forward.")
				path[v.neighbour_node] = u
				capmap[v.neighbour_node] = Min(capmap[u], v.capacity-v.weight)
				fmt.Println("The capacity of the path is now: ", capmap[v.neighbour_node])
				if v.neighbour_node != sink {
					fmt.Println("We have not reached the sink. We enqueue v.neighbour_node and continue.")
					q.Enqueue(v.neighbour_node)
					fmt.Println("The queue is now: ", q)
				} else {
					fmt.Println("We have reached the sink and we return.")
					return capmap[sink], path
				}
			} else if v.capacity < 0 && v.weight < 0 && path[v.neighbour_node] == notvisited {
				fmt.Println("Path can proceed from u to v by pushing flow backward.")
				path[v.neighbour_node] = u
				capmap[v.neighbour_node] = Min(capmap[u], v.weight-v.capacity)
				fmt.Println("The capacity of the path is now: ", capmap[v.neighbour_node])
				if v.neighbour_node != sink {
					fmt.Println("We have not reached the sink.  We enqueue v.neighbour_node and continue.")
					q.Enqueue(v.neighbour_node)
					fmt.Println("The queue is now: ", q)
				} else {
					fmt.Println("We have reached the sink and we return.")
					return capmap[sink], path
				}
			}
		}
	}
	fmt.Println("BFS completed.")
	return 0, path
}
