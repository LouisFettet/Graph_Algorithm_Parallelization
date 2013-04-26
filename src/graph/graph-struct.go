// graphstruct.go
// Implementation of a graph structure with functions to painlessly modify the
// contents and to extract relevant data.

// Package graph provides both primitives for initializing graph structures
// and functions to solve the maximum flow of a graph.  Testing methods and 
// graphic interfaces are also available. 
package graph

import (
	//random integers
	"math/rand"
)

// Node structure containing a simple (x, y) integer coordinate set. 
type Node struct {
	X, Y int
}

// Neighbour structure containing a "connected" node, along with integers for 
// the weight and capacity of the connection.
type Neighbour struct {
	Neighbour_Node   Node
	Weight, Capacity int
}

// Graph structure containing a map in which the nodes are values and a list 
// of neighbours connected to that node is the key.
type Graph struct {
	Nodemap map[Node][]Neighbour
}

// Function GenBlankGraph initializes a new empty graph.
func GenBlankGraph() *Graph {
	return &Graph{
		Nodemap: make(map[Node][]Neighbour),
	}
}

// Function GetNeighbours is a graph-associated method that returns the 
// neighbours of a given node.
func (g *Graph) GetNeighbours(node Node) []Neighbour {
	return g.Nodemap[node]
}

// Function GetWeight is a graph-associated method that returns the weight
// of a connection from a given source node to a given destination node.
func (g *Graph) GetWeight(source Node, destination Node) int {
	// Loop until we find the correct neighbour.  
	for _, possdest := range g.GetNeighbours(source) {
		if possdest.Neighbour_Node == destination {
			return possdest.Weight
		}
	}
	// Return zero if connection doesn't exist.
	return 0
}

// Function GetNodeList is a graph-associated method that returns a list
// of all initialized nodes.  Useful for iteration.
func (g *Graph) GetNodeList() []Node {
	// Initialize an empty list.
	nodelist := []Node{}
	// Grab every node in the graph and append them to the list.
	for node := range g.Nodemap {
		nodelist = append(nodelist, node)
	}
	return nodelist
}

// Function AddNode is a graph-associated method that initializes a node
// value, with no neighbours, into the map.  Returns the graph exactly as is
// if the node is already initialized.
func (g *Graph) AddNode(node Node) *Graph {
	if _, found := g.Nodemap[node]; !found {
		null := []Neighbour{}
		g.Nodemap[node] = null
	}
	return g
}

// Function AddNeighbour is a graph-associated method that initializes a 
// connection with no weight between a given source and destination node and 
// records the connection and its residual connection into the graph.  Returns
// the graph exactly as is if the nodes are the same, if the nodes do not 
// exist, or if the connection already exists.
func (g *Graph) AddNeighbour(source Node, destination Node, capacity int) *Graph {
	// Check to make sure the nodes are different.
	if source == destination {
		return g
	}
	// Check for existence of nodes in graph.
	if _, found := g.Nodemap[source]; !found {
		return g
	} else if _, found := g.Nodemap[destination]; !found {
		return g
	}
	// Check to see if the connection already exists.
	// We need only check one node to confirm a connection.
	neighbourlist := g.GetNeighbours(source)
	for neighbour := range neighbourlist {
		if neighbourlist[neighbour].Neighbour_Node == destination {
			return g
		}
	}
	// If we get here it means that both nodes exist in the graph, but
	// there is no connection between them.  At this point we know to
	// just add them into the nodemap.
	neighbour := Neighbour{destination, 0, capacity}
	g.Nodemap[source] = append(g.Nodemap[source], neighbour)
	neighbour = Neighbour{source, 0, -(capacity)}
	g.Nodemap[destination] = append(g.Nodemap[destination], neighbour)
	return g
}

// Function UpdateWeight is a graph-associated method that updates the weight
// of both the connection between the source and destination and the residual 
// connection between the destination and source. Returns the graph exactly 
// as is if the new weight is greater than the connection's capacity.
func (g *Graph) UpdateWeight(source Node, destination Node, weight int) *Graph {
	neighbourlistsource := g.GetNeighbours(source)
	for i := 0; i < len(neighbourlistsource); i++ {
		//Iterate until we find the destination...
		if neighbourlistsource[i].Neighbour_Node == destination {
			//Check to make sure the weight is not greater than
			//the capacity.
			if Abs(weight) <= Abs(neighbourlistsource[i].Capacity) {
				//Update the connection with the new weight.
				updatedconnect := Neighbour{destination, weight, neighbourlistsource[i].Capacity}
				g.Nodemap[source][i] = updatedconnect
			}
		}
	}
	//Repeat the same process for the destination, but take the residual.
	neighbourlistdest := g.GetNeighbours(destination)
	for i := 0; i < len(neighbourlistdest); i++ {
		if neighbourlistdest[i].Neighbour_Node == source {
			if Abs(weight) <= Abs(neighbourlistdest[i].Capacity) {
				updatedconnect := Neighbour{source, -weight, neighbourlistdest[i].Capacity}
				g.Nodemap[destination][i] = updatedconnect
			}
		}
	}
	// Both connections have been updated, and we return.
	return g
}

// Function GenRandomGraph initializes a new graph, fills it with a given 
// number of nodes and connections with random values, and returns the graph.
func GenRandomGraph(nodenum, connectnum, maxcap int) *Graph {
	// Initialize a new graph.
	g := GenBlankGraph()
	for i := 0; i < nodenum; i++ {
		//Randomly choose two integers to make a node.
		xcoord := rand.Intn(nodenum * 2)
		ycoord := rand.Intn(nodenum * 2)
		newnode := Node{xcoord, ycoord}
		// Check to make sure it doesn't exist.
		if _, found := g.Nodemap[newnode]; !found {
			g.AddNode(newnode)
		} else {
			// Step backward if it already exists.
			i = i - 1
		}
	}
	// Put all the nodes into a list so they can be randomly chosen.
	nodelist := g.GetNodeList()
	// Loop to add connections.
	for i := 0; i < connectnum; i++ {
		node1 := nodelist[rand.Intn(nodenum)]
		node2 := nodelist[rand.Intn(nodenum)]
		//Can't connect a node to itself, so we loop until the nodes 
		//are different.
		for node1 == node2 {
			node2 = nodelist[rand.Intn(nodenum)]
		}
		randcap := rand.Intn(maxcap)
		// A capacity of zero is useless, so loop until not zero.
		for randcap == 0 {
			randcap = rand.Intn(maxcap)
		}
		g.AddNeighbour(node1, node2, randcap)
	}
	return g
}
