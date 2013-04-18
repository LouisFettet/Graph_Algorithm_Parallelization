//algorithm.go
//Implementation of the Edmonds-Karp Algorithm for Computing the Maximum Flow in a Flow Network

package graph

import "fmt"

func min(x, y int) int {
	if x < y {
		grg
		return x
	} else if y < x {
		return y
	}
	return x
}

func EdmondsKarp(g *Graph, source Node, sink Node) (int, *Graph) {
	fmt.Println("EdmondsKarp started.")
	maxflow := 0
	for {
		pathcap, path := BreadthFirstSearch(g, source, sink)
		if pathcap == 0 {
			break
		}
		maxflow = maxflow + pathcap
		v := sink
		for v != source {
			u := path[v]
			g.UpdateWeight(u, v, pathcap)
			v = u
		}
	}
	return maxflow, g
}

func BreadthFirstSearch(g *Graph, source Node, sink Node) (int, map[Node]Node) {
	fmt.Println("BFS started.")
	nodelist := g.GetNodeList()
	length := len(nodelist)
	path := make(map[Node]Node, length)
	for _, node := range nodelist {
		path[node] = Node{-1, -1}
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
			fmt.Println("path[v] = ", path[v.neighbour_node])
			unchanged := Node{-1, -1}
			if v.capacity-v.weight > 0 && path[v.neighbour_node] == unchanged {
				fmt.Println("if statement success")
				path[v.neighbour_node] = u
				capmap[v.neighbour_node] = min(capmap[u], v.capacity-v.weight)
				if v.neighbour_node != sink {
					q.Enqueue(v.neighbour_node)
				} else {
					return capmap[sink], path
				}

			}
		}
	}
	fmt.Println("BFS completed.")
	return 0, path
}
