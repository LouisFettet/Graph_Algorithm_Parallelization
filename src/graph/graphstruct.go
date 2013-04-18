//graphstruct.go
//Implementation of a Graph Data Structure

package graph

import (
	//random integers
	"math/rand"
)

type Node struct {
	/*
	 Node structure containing an integer coordinate set. 
	      {x, y}
	*/
	X, Y int
}

type Neighbour struct {
	/*
	 Neighbour structure containing a "connected" node, along with
	 integers for the weight and capacity of the connection.
	      {{neighbour_node.x, neighbour_node.y}, weight, capacity}
	*/
	neighbour_node   Node
	weight, capacity int
}

type Graph struct {
	/*
	 Graph structure containing a map in which the nodes are values and a 
	 list of neighbours connected to that node is the key.
	 map[
	      {node1}:[
	           {{neighbour1} weight1 capacity1}, 
	           {{neighbour2} weight2 capacity2} 
	      ]
	      {node2}:[
	           {{neighbour3} weight3 capacity3}, 
	           {{neighbour4} weight4 capacity4} 
	      ]
	 ]
	*/
	Nodemap map[Node][]Neighbour
}

func GenBlankGraph() *Graph {
	//Mandatory initialization of graph's nodemap structure
	return &Graph{
		Nodemap: make(map[Node][]Neighbour),
	}
}

func (g *Graph) GetNeighbours(node Node) []Neighbour {
	/*
	 Graph-associated method to return the neighbours of a node.
	*/
	return g.Nodemap[node]
}

func (g *Graph) GetWeight(source Node, destination Node) int {
	for _, possdest := range g.GetNeighbours(source) {
		if possdest.neighbour_node == destination {
			return possdest.weight
		}
	}
	return 0
}

func (g *Graph) GetNodeList() []Node {
	/*
	 Graph-associated method to return a list of all initialized nodes.
	*/
	nodelist := []Node{}
	for node := range g.Nodemap {
		nodelist = append(nodelist, node)
	}
	return nodelist
}

func (g *Graph) AddNode(node Node) *Graph {
	/*
	 Graph-associated method to initialize a node value, with no
	 neighbours, into the map.  Returns the graph exactly as is if the 
	 node is already initialized.
	*/
	if _, found := g.Nodemap[node]; !found {
		null := []Neighbour{}
		g.Nodemap[node] = null
	}
	return g
}

func (g *Graph) AddNeighbour(source Node, destination Node, capacity int) *Graph {
	/*
	 Graph-associated method to initialize a connection with no weight 
	 between a source node and destination node and record the connection 
	 and its residual connection into the map.
	 Returns the graph exactly as is if the nodes are the same, if the 
	 nodes do not exist, or if the connection already exists.
	*/
	//Check to make sure the nodes are different.
	if source == destination {
		return g
	}
	//Check for existence of nodes in graph.
	if _, found := g.Nodemap[source]; !found {
		return g
	} else if _, found := g.Nodemap[destination]; !found {
		return g
	}
	//Check to see if the connection already exists.
	//We need only check one node to confirm a connection.
	neighbourlist := g.GetNeighbours(source)
	for neighbour := range neighbourlist {
		if neighbourlist[neighbour].neighbour_node == destination {
			return g
		}
	}
	//If we get here it means that both nodes exist in the graph, but
	//there is no connection between them.  At this point we know to
	//just add them into the nodemap.
	neighbour := Neighbour{destination, 0, capacity}
	g.Nodemap[source] = append(g.Nodemap[source], neighbour)
	neighbour = Neighbour{source, 0, -(capacity)}
	g.Nodemap[destination] = append(g.Nodemap[destination], neighbour)
	return g
}

func (g *Graph) UpdateWeight(source Node, destination Node, weight int) *Graph {
	/*
	 Graph-associated method to update the weight of a connection.
	 Returns the graph exactly as is if the new weight is greater than 
	 the connection's initial capacity.
	 Note:  Unlike AddConnection(), the order in which the nodes are 
	 input into this function DOES matter. 
	      node1's connection with node2 is updated with the weight given
	      node2's connection with node1 is updated with the residual
	*/
	neighbourlistsource := g.GetNeighbours(source)
	for i := 0; i < len(neighbourlistsource); i++ {
		//Iterate until we find node2...
		if neighbourlistsource[i].neighbour_node == destination {
			//Check to make sure the weight is not greater than
			//the capacity.
			if Abs(weight) <= Abs(neighbourlistsource[i].capacity) {
				//Update the connection with the new weight.
				updatedconnect := Neighbour{destination, weight, neighbourlistsource[i].capacity}
				g.Nodemap[source][i] = updatedconnect
			}
		}
	}
	//Repeat the process for node2, but take the reciprocal.
	neighbourlistdest := g.GetNeighbours(destination)
	for i := 0; i < len(neighbourlistdest); i++ {
		if neighbourlistdest[i].neighbour_node == source {
			if Abs(weight) <= Abs(neighbourlistdest[i].capacity) {
				updatedconnect := Neighbour{source, -weight, neighbourlistdest[i].capacity}
				g.Nodemap[destination][i] = updatedconnect
			}
		}
	}
	return g
}

func GenRandomGraph(nodenum, connectnum int) *Graph {
	//Initialize a new graph
	g := GenBlankGraph()
	for i := 0; i < nodenum; i++ {
		//Randomly choose two integers to make a node.
		xcoord := rand.Intn(nodenum * 2)
		ycoord := rand.Intn(nodenum * 2)
		newnode := Node{xcoord, ycoord}
		//Check to make sure it doesn't exist.
		if _, found := g.Nodemap[newnode]; !found {
			g.AddNode(newnode)
		} else {
			//Step backward.
			i = i - 1
		}
	}
	//Put all the nodes into a list so they can be randomly chosen.
	nodelist := g.GetNodeList()
	//Loop to add connections
	for i := 0; i < connectnum; i++ {
		node1 := nodelist[rand.Intn(nodenum)]
		node2 := nodelist[rand.Intn(nodenum)]
		//Can't connect a node to itself, so we loop until the nodes 
		//are different.
		for node1 == node2 {
			node2 = nodelist[rand.Intn(nodenum)]
		}
		g.AddNeighbour(node1, node2, rand.Intn(100))
	}
	return g
}
