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
	 	   {{neighbour1} weight1 capacity1}, 
	           {{neighbour2} weight2 capacity2} 
	      ]
	 ]
	*/
	Nodemap map[Node][]Neighbour
}

func GenBlankGraph() *Graph {
	//Mandatory initialization of graph's nodemap structure
	return &Graph{Nodemap: make(map[Node][]Neighbour)}
}

func (g *Graph) GetNeighbours(node Node) []Neighbour {
	/*
	 Graph-associated method to return the neighbours of a node.
	*/
	return g.Nodemap[node]
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

func (g *Graph) AddConnection(node1 Node, node2 Node, cost int, total int) *Graph {
	/*
	 Graph-associated method to add or modify a connection between two
	 nodes and record them into the map.  Returns the graph exactly as is
	 if the nodes do not exist.
	*/
	//Check to make sure the nodes are different
	if node1 == node2 {
		return g
	}
	//Check for existence of nodes in graph
	if _, found := g.Nodemap[node1]; !found {
		return g
	} else if _, found := g.Nodemap[node2]; !found {
		return g
	}
	//Check to see if connection already exists between nodes
	//We need only check one node to confirm a connection, but we will
	//make copies of both lists.
	neighbourlistnode1 := g.GetNeighbours(node1)
	neighbourlistnode2 := g.GetNeighbours(node2)
	for neighbour := range neighbourlistnode1 {
		if neighbourlistnode1[neighbour].neighbour_node == node2 {
			//The connection exists, so we have to check for an
			//update to the weight.  First we verify the capacity.
			if neighbourlistnode1[neighbour].capacity == total {
				//We now know the function was called to 
				//update the weight of the connection.
				//So we iterate through both node keys,
				//update the correct location, and return g.
				for i := 0; i < len(neighbourlistnode1); i++ {
					if neighbourlistnode1[i].neighbour_node == node2 {
						updatedconnect := Neighbour{node2, cost, total}
						g.Nodemap[node1][i] = updatedconnect
					}
				}
				for i := 0; i < len(neighbourlistnode2); i++ {
					if neighbourlistnode2[i].neighbour_node == node1 {
						updatedconnect := Neighbour{node1, cost, total}
						g.Nodemap[node2][i] = updatedconnect
					}
				}
				return g
			} else {
				//The nodes are connected, but the function was called using
				//a different capacity.  Capacity is a constant that once
				//set cannot be changed.  Only the weight can change.
				return g
			}
		}
	}
	//If we get here it means that both nodes exist in the graph, but
	//there is no connection between them.  At this point we know to
	//just add them into the nodemap.
	neighbour1 := Neighbour{node2, cost, total}
	neighbour2 := Neighbour{node1, cost, total}
	g.Nodemap[node1] = append(g.Nodemap[node1], neighbour1)
	g.Nodemap[node2] = append(g.Nodemap[node2], neighbour2)
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
	nodelist := []Node{}
	for node := range g.Nodemap {
		nodelist = append(nodelist, node)
	}
	//Loop to add connections
	for i := 0; i < connectnum; i++ {
		node1 := nodelist[rand.Intn(nodenum)]
		node2 := nodelist[rand.Intn(nodenum)]
		//Can't connect a node to itself
		for node1 == node2 {
			node2 = nodelist[rand.Intn(nodenum)]
		}
		g.AddConnection(node1, node2, 0, rand.Intn(100))
	}
	return g
}