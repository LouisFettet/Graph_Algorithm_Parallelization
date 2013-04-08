package main

import (
	//printing
	"fmt"
	//random integers
//	"math/rand"
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
	 Neighbour structure for the graph
	*/
	neighbour        Node
	weight, capacity int
}

type Graph struct {
	/*
	 Graph structure containing a map data structure in which the nodes
	 are values and a list of arcs connected to that node is the key.
	      {map[
	        {node1.x, node1.y}:
	           [{{node1.x, node1.y},{node2.x, node2.y}, weight, capacity}]
	        {node2.x, node2.y}:
	           [{{node1.x, node1.y},{node2.x, node2.y}, weight, capacity}]
	      ]}
	*/
	nodemap map[Node][]Neighbour
	arclist []Arc
}

func genBlankGraph() *Graph {
	//Initialization of graph's nodemap structure
	return &Graph{nodemap: make(map[Node][]Neighbour)}
}

/*
func genRandomGraph(nodenum, arcnum int) *Graph {
	g := genBlankGraph()
	for i := 0; i < nodenum; i++ {
		xcoord := rand.Intn(nodenum / 2)
		ycoord := rand.Intn(nodenum / 2)
		nodetoadd := Node{xcoord, ycoord}
		if _, found := g.nodemap[nodetoadd]; !found {
			g.addNode(nodetoadd)
		} else {
			i = i - 1
		}
	}
	nodelist := []Node{}
	for node := range g.nodemap {
pofd		nodelist = append(nodelist, node)
		fmt.Println(nodelist)
	}
	length := len(nodelist)
	for i := 0; i < arcnum; i++ {
		arctoadd := Arc{nodelist[rand.Intn(length)], nodelist[rand.Intn(length)], 0, rand.Intn(100)}
		g.addArc(arctoadd)
	}
	return g
}
*/

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
		fmt.Println("\tNode ", node, " already exists in graph.")
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

func main() {

	n := Node{1, 3}
	m := Node{8, 9}
	p := Node{4, 2}
	r := Node{3, 4}
	s := Node{5, 10}

	a := Arc{n, m, 0, 10}
	b := Arc{m, n, 1, 10}
	c := Arc{n, m, 1, 10}
	d := Arc{m, p, 0, 3}
	e := Arc{r, s, 0, 5}
	//b := Arc{m, p, 0, 3}
	//c := Arc{m, p, 1, 3}
	//d := Arc{m, q, 0, 5}

	g := genBlankGraph()

	g.addNode(n)
	g.addNode(m)
	g.addNode(p)

	g.addArc(a)
	g.addArc(a)
	g.addArc(b)
	g.addArc(c)
	g.addArc(d)
	g.addArc(e)

	fmt.Println("\nThe graph contains: \n\t", g, "\n")

	/*
		g.addArc(a)

		fmt.Println("\nThe graph contains: \n\t", g, "\n")

		g.addArc(b)

		fmt.Println("\nThe graph contains: \n\t", g, "\n")

		g.addArc(c)

		fmt.Println("\nThe graph contains: \n\t", g, "\n")

		g.addArc(d)

		fmt.Println(g)*/
	//k := genRandomGraph(10, 5)
	//fmt.Println(k)
}

/*
func (g *Graph) addArc(arc Arc) *Graph {
	if _, found := g.nodemap[arc.node1]; !found {
		fmt.Println("Error, could not find node.")
	} else {
		arclist1 := g.nodemap[arc.node1]
		arclist2 := g.nodemap[arc.node2]
		if len(arclist1) == 0 {
			g.nodemap[arc.node1] = append(arclist1, arc)
			g.nodemap[arc.node2] = append(arclist2, arc)
		} else {
		for i:=0; i<len(arclist1); i++ {
			fmt.Println(arclist1[i])
			if arclist1[i] == arc {
				fmt.Println("Error, arc already exists.")
			}
			if arclist2[i] == arc {
				fmt.Println("Error, arc already exists.")
			}else {
				g.nodemap[arc.node1] = append(arclist1, arc)
				g.nodemap[arc.node2] = append(arclist2, arc)
				fmt.Println("arc has been added")
			} 
		}
	}
	}
	fmt.Println(g)
	return g
}
*/

/*
func main(){
	n := Node{3,5}
	m := Node{2,4}
	a := Arc{n,m,5,0}
	g := genGraph()
	arclist := []Arc{}
	arclist = append(arclist,a)
	//_, ok := g.nodemap[n]
	g.nodemap[n] = arclist
	fmt.Println(g.nodemap)
}
*/

/*
func (g *Graph) addNode(node Node) *Graph {
	g.nodes	= append(g.nodes, node)
	return g
}


func (g *Graph) addArc(arc Arc) *Graph {
	g.arcs = append(g.arcs, arc)
	return g
}

func genGraph(n, a int) Graph {
	g := Graph{}
	for i:=0; i<n; i++ {
		j := (i+1)%n
		fmt.Println(i,j)
		newnode :=
		newarc := Arc{i,j,1000,0}
		g.addArc(newarc)
	}
	return g
}

func main(){
	g := genGraph(50,100)

	g := Graph{}
	n := Node{1,3}
	m := Node{4,5}
	a := Arc{n,m,6,0}
	g.addNode(n)
	g.addNode(m)
	g.addArc(a)

	fmt.Println(g)
}

*/

/*
type Graph struct {
	Node struct {
		x, y int
	}
	Arc struct {
		node1, node2 Node
		capacity, weight int
	}
}

//type Build interface {}

func main(){
	g := Graph{}
	n := g.Node
	a := g.Arc
	fmt.Println(g, n, a)
}
*/

/*
type Node struct {
	x, y int
}

type Arc struct {
	node1, node2 Node
	capacity, weight int
}

type Graph struct {
	nodes []Node
}


func (g *Graph) addNode (node Node) []Node {
	graph.nodes = append(graph.nodes, node)
	return graph.nodes
}

*/