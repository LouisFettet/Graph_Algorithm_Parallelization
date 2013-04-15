//algorithm.go
//Implementation of the Edmonds-Karp Algorithm for Computing the Maximum Flow in a Flow Network

package graph

import (
	"fmt"
)

func BreadthFirstSearch(g *Graph) {
	q := NewQueue(3)
	q.Push(Node{0, 0})
	q.Push(Node{1, 1})
	q.Push(Node{2, 2})
	fmt.Println(q.Pop(), q.Pop(), q.Pop())
}

func EdmondsKarp(g *Graph, source Node, sink Node) (int, *Graph) {
	maxflow := 0
	return maxflow, g
}
