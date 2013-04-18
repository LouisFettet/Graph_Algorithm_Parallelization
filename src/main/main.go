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
	gograph.AddNeighbour(a, b, 3)
	gograph.AddNeighbour(a, d, 3)
	gograph.AddNeighbour(b, c, 4)
	gograph.AddNeighbour(c, a, 3)
	gograph.AddNeighbour(c, d, 1)
	gograph.AddNeighbour(c, e, 2)
	gograph.AddNeighbour(d, e, 2)
	gograph.AddNeighbour(d, f, 6)
	gograph.AddNeighbour(f, g, 9)
	gograph.AddNeighbour(e, b, 1)
	gograph.AddNeighbour(e, g, 1)
	fmt.Println(gograph)

	maxflow, solution := graph.EdmondsKarp(gograph, a, g)
	fmt.Println(maxflow, solution)
}
