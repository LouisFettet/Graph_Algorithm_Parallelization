package main

import (
	"fmt"
	"graph"
)

func main() {
	g := graph.GenBlankGraph()
	a := graph.Node{0, 0}
	b := graph.Node{0, 1}
	c := graph.Node{1, 2}
	d := graph.Node{2, 3}
	e := graph.Node{4, 4}
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddNode(d)
	g.AddNode(e)
	g.AddConnection(a, b, 0, 10)
	g.AddConnection(b, c, 0, 9)
	g.AddConnection(c, d, 0, 8)
	g.AddConnection(d, e, 0, 7)
	g.AddConnection(b, d, 0, 6)
	fmt.Println("The graph's nodemap contains:\n", g.Nodemap)
}
