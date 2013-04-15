//main.go
//Executable File for Testing and Example Purposes

package main

import (
	"fmt"
	"graph"
)

func main() {
	gograph := graph.GenBlankGraph()
	a := graph.Node{0, 0}
	b := graph.Node{0, 2}
	c := graph.Node{1, 1}
	d := graph.Node{2, 0}
	e := graph.Node{2, 2}
	f := graph.Node{3, 0}
	g := graph.Node{3, 2}
	gograph.AddNode(a)
	gograph.AddNode(b)
	gograph.AddNode(c)
	gograph.AddNode(d)
	gograph.AddNode(e)
	gograph.AddNode(f)
	gograph.AddNode(g)
	gograph.AddConnection(a, b, 0, 3)
	gograph.AddConnection(a, c, 0, 3)
	gograph.AddConnection(a, d, 0, 3)
	gograph.AddConnection(b, c, 0, 4)
	gograph.AddConnection(b, e, 0, 1)
	gograph.AddConnection(c, d, 0, 1)
	gograph.AddConnection(c, e, 0, 2)
	gograph.AddConnection(d, e, 0, 2)
	gograph.AddConnection(d, f, 0, 6)
	gograph.AddConnection(e, g, 0, 1)
	gograph.AddConnection(f, g, 0, 9)
	flow, gograph := graph.EdmondsKarp(gograph, a, g)
	fmt.Println(flow, "\n", gograph)
	graph.BreadthFirstSearch(gograph)
}
