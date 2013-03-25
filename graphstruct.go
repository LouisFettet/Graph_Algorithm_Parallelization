package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	x, y int
}

type Arc struct {
	node1, node2 Node
	weight, capacity int
}

type Graph struct {
	nodemap map[Node][]Arc
}

func genBlankGraph() *Graph {
	return &Graph{nodemap: make(map[Node][]Arc)}
}

func genRandomGraph(nodenum, arcnum int) *Graph {
	g := genBlankGraph()
	for i:=0; i<nodenum; i++ {
		xcoord := rand.Intn(nodenum/2)
		ycoord := rand.Intn(nodenum/2)
		nodetoadd := Node{xcoord,ycoord}
		if _, found := g.nodemap[nodetoadd]; !found{
			g.addNode(nodetoadd)
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
	for i:=0; i<arcnum; i++ {
		arctoadd := Arc{nodelist[rand.Intn(length)], nodelist[rand.Intn(length)], 0, rand.Intn(100)}
		g.addArc(arctoadd)
		}
	return g
}

func (g *Graph) addNode(node Node) *Graph {
	if _, found := g.nodemap[node]; !found {
		null := []Arc{}
		g.nodemap[node] = null
	} 
	return g
}

func (g *Graph) addArc(arc Arc) *Graph {
	//fmt.Println("addArc() started")
	if _, found := g.nodemap[arc.node1]; !found {
		//fmt.Println("Error, could not find node1 ", arc.node1)
	} else if _, found := g.nodemap[arc.node2]; !found {
		//fmt.Println("Error, could not find node2 ", arc.node2)
	} else {
		arclist1 := g.nodemap[arc.node1]
		arclist2 := g.nodemap[arc.node2]
		length1 := len(arclist1)
		length2 := len(arclist2)
		//fmt.Println("arclist1 is ", arclist1, "and len is ", length1)
		//fmt.Println("arclist2 is ", arclist2, "and len is ", length2)
		if length1 == 0 || length2 == 0 {
			//fmt.Println("the if check for length worked!")
			g.nodemap[arc.node1] = append(arclist1, arc)
			g.nodemap[arc.node2] = append(arclist2, arc)
			return g
		} else {
			for i:=0; i<length1; i++ {
				if arclist1[i] == arc {
					//fmt.Println("Error, arc ", arclist1[i], " already exists.(1)")
					return g
				}
			
			}
			for i:=0; i<length2; i++ {
				if arclist2[i] == arc {
					//fmt.Println("Error, arc ", arclist2[i], " already exists.(2)")
					return g
				}
			}
			g.nodemap[arc.node1] = append(arclist1, arc)
			g.nodemap[arc.node2] = append(arclist2, arc)
		}
	}
	return g
}

func main(){
	n := Node{1,3}
	m := Node{6,9}
	p := Node{4,2}
	a := Arc{n,m,0,10}
	b := Arc{m,p,0,3}
	g := genBlankGraph()
	g.addNode(n)
	g.addNode(m)
	g.addNode(p)
	g.addArc(a)
	g.addArc(b)
	k := genRandomGraph(10,5)
	fmt.Println(k)
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