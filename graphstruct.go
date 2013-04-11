package main

import (
	//printing
	"fmt"
	//random integers
	"math/rand"
)

type Node struct {
	/*
	 Node structure containing an integer coordinate set. 
	      {x, y}
	*/
	x, y int
}

type Arc struct {
	/*
	 Arc structure containing two node structures and integers for weight 
	 and capacity.  
	      {{node1.x, node1.y},{node2.x, node2.y}, weight, capacity}
	*/
	node1, node2     Node
	weight, capacity int
}

type Neighbour struct {
	/*
	 Neighbour structure for the graph's map.
	      {{neighbour.x, neighbour.y}, weight, capacity}
	*/
	neighbour        Node
	weight, capacity int
}

type Graph struct {
	/*
	 Graph structure containing two distinct structures:
	 1. A map data structure in which the nodes are values and a list 
	    of neighbours connected to that node is the key.
	 2. A list of all arcs connecting nodes in the graph.
	*/
	nodemap map[Node][]Neighbour
	arclist []Arc
}

func genBlankGraph() *Graph {
	//Initialization of graph's nodemap structure
	return &Graph{nodemap: make(map[Node][]Neighbour)}
}

func genRandomGraph(nodenum, arcnum int) *Graph {
	fmt.Println("genRandomGraph() called for ", nodenum, " nodes and ", arcnum, "connecting arcs.")
	g := genBlankGraph()
	for i := 0; i < nodenum; i++ {
		xcoord := rand.Intn(nodenum * 2)
		ycoord := rand.Intn(nodenum * 2)
		newnode := Node{xcoord, ycoord}
		if _, found := g.nodemap[newnode]; !found {
			g.addNode(newnode)
		} else {
			i = i - 1
		}
	}
	nodelist := []Node{}
	for node := range g.nodemap {
		nodelist = append(nodelist, node)
		fmt.Println(nodelist)
	}
	length := len(nodelist)
	for i := 0; i < arcnum; i++ {
		node1 := nodelist[rand.Intn(length)]
		node2 := nodelist[rand.Intn(length)]
		for node1 == node2 {
			node2 = nodelist[rand.Intn(length)]
		}
		arctoadd := Arc{node1, node2, 0, rand.Intn(100)}
		g.addArc(arctoadd)
	}
	return g
}

func (g *Graph) addNode(node Node) *Graph {
	/*
	 Graph-associated method to initialize a node value, with no attached
	 arcs, into the map.  Returns the graph as is if the node is already
	 initialized.
	*/
	fmt.Println("\naddNode() called for node ", node)
	if _, found := g.nodemap[node]; !found {
		null := []Neighbour{}
		g.nodemap[node] = null
		fmt.Println("\tNode ", node, " added to graph.")
	} else {
		fmt.Println("\tNode ", node, " already exists in graph. Exit.")
	}
	return g
}

func (g *Graph) addArc(arc Arc) *Graph {
	fmt.Println("\naddArc() called for arc ", arc)
	fmt.Println("\tNode existence check start.")
	if _, found := g.nodemap[arc.node1]; !found {
		fmt.Println("\t\tError, could not find node", arc.node1)
		return g
	} else if _, found := g.nodemap[arc.node2]; !found {
		fmt.Println("\t\tError, could not find node", arc.node2)
		return g
	}
	fmt.Println("\t\tBoth nodes exist.")
	fmt.Println("\tArc repetition check start.")
	for i := 0; i < len(g.arclist); i++ {
		if ((g.arclist[i].node1 == arc.node1) || (g.arclist[i].node1 == arc.node2)) && ((g.arclist[i].node2 == arc.node1) || (g.arclist[i].node2 == arc.node2)) {
			fmt.Println("\t\tArc connecting nodes already exists, so we exit.")
			return g
		}
	}
	fmt.Println("\t\tArc does not exist in graph, so it will be added.")
	g.arclist = append(g.arclist, arc)
	fmt.Println("\t\tArc has been added into the list; ", g.arclist)
	fmt.Println("\t\tNow neighbours will be populated.")
	neighbour1 := Neighbour{arc.node2, arc.weight, arc.capacity}
	neighbour2 := Neighbour{arc.node1, arc.weight, arc.capacity}
	g.nodemap[arc.node1] = append(g.nodemap[arc.node1], neighbour1)
	g.nodemap[arc.node2] = append(g.nodemap[arc.node2], neighbour2)

	fmt.Println("\t\tNeighbours of ", arc.node1, " are now ", g.nodemap[arc.node1])
	fmt.Println("\t\tNeighbours of ", arc.node2, " are now ", g.nodemap[arc.node2])
	return g
}

func (g *Graph) getNeighbours(n Node) []Neighbour {
	return g.nodemap[n]
}

func main() {
	g := genRandomGraph(5, 7)
	fmt.Println("\nThe graph's nodemap contains:\n", g.nodemap)
	a := Node{6,0}
	neighbourlist := g.getNeighbours(a)
	fmt.Println("\n", neighbourlist)
}