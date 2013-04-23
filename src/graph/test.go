// test.go
// Generates graphs based on examples that can be found in the sample-problems
// directory and runs the Edmonds-Karp algorithm on them.

package graph

import "math/rand"

// Function SolveGraph1 runs the algorithm on the graph 
// sample-problems/problem1;solution-5.png
func SolveGraph1() (int, *Graph) {
	gr := GenBlankGraph()
	a := Node{0, 0}
	b := Node{0, 2}
	c := Node{1, 1}
	d := Node{2, 0}
	e := Node{2, 2}
	f := Node{3, 0}
	g := Node{3, 2}
	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)
	gr.AddNode(f)
	gr.AddNode(g)
	gr.AddNeighbour(a, b, 3)
	gr.AddNeighbour(a, d, 3)
	gr.AddNeighbour(b, c, 4)
	gr.AddNeighbour(c, a, 3)
	gr.AddNeighbour(c, d, 1)
	gr.AddNeighbour(c, e, 2)
	gr.AddNeighbour(d, e, 2)
	gr.AddNeighbour(d, f, 6)
	gr.AddNeighbour(f, g, 9)
	gr.AddNeighbour(e, b, 1)
	gr.AddNeighbour(e, g, 1)
	flow, solution := EdmondsKarp(gr, a, g)
	return flow, solution
}

// Function SolveGraph2 runs the algorithm on the graph 
// sample-problems/problem2;solution-5.svg
func SolveGraph2() (int, *Graph) {
	gr := GenBlankGraph()
	s := Node{0, 1}
	o := Node{1, 0}
	p := Node{1, 2}
	q := Node{2, 0}
	r := Node{2, 2}
	t := Node{3, 1}
	gr.AddNode(s)
	gr.AddNode(o)
	gr.AddNode(p)
	gr.AddNode(q)
	gr.AddNode(r)
	gr.AddNode(t)
	gr.AddNeighbour(s, o, 3)
	gr.AddNeighbour(s, p, 3)
	gr.AddNeighbour(o, p, 2)
	gr.AddNeighbour(o, q, 3)
	gr.AddNeighbour(p, r, 2)
	gr.AddNeighbour(q, r, 4)
	gr.AddNeighbour(q, t, 2)
	gr.AddNeighbour(r, t, 3)
	flow, solution := EdmondsKarp(gr, s, t)
	return flow, solution
}

// Function SolveGraph3 runs the algorithm on the graph 
// sample-problems/problem3;solution-15.jpg
func SolveGraph3() (int, *Graph) {
	gr := GenBlankGraph()
	s := Node{0, 1}
	u := Node{1, 0}
	v := Node{1, 2}
	t := Node{2, 1}
	gr.AddNode(s)
	gr.AddNode(u)
	gr.AddNode(v)
	gr.AddNode(t)
	gr.AddNeighbour(s, u, 10)
	gr.AddNeighbour(s, v, 5)
	gr.AddNeighbour(u, v, 15)
	gr.AddNeighbour(u, t, 5)
	gr.AddNeighbour(v, t, 10)
	flow, solution := EdmondsKarp(gr, s, t)
	return flow, solution
}

// Function SolveGraph4 runs the algorithm on the graph 
// sample-problems/problem4;solution-3.gif
func SolveGraph4() (int, *Graph) {
	gr := GenBlankGraph()
	x := Node{0, 1}
	a := Node{1, 0}
	b := Node{1, 1}
	c := Node{2, 1}
	d := Node{1, 2}
	e := Node{3, 2}
	y := Node{3, 1}
	gr.AddNode(x)
	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)
	gr.AddNode(y)
	gr.AddNeighbour(x, a, 3)
	gr.AddNeighbour(x, b, 1)
	gr.AddNeighbour(a, c, 3)
	gr.AddNeighbour(b, c, 5)
	gr.AddNeighbour(b, d, 4)
	gr.AddNeighbour(c, y, 2)
	gr.AddNeighbour(d, e, 2)
	gr.AddNeighbour(e, y, 3)
	flow, solution := EdmondsKarp(gr, x, y)
	return flow, solution
}

// Function SolveGraph5 runs the algorithm on the graph 
// sample-problems/problem5;solution-4.gif
func SolveGraph5() (int, *Graph) {
	gr := GenBlankGraph()
	a := Node{0, 2}
	b := Node{1, 2}
	c := Node{2, 3}
	d := Node{3, 1}
	e := Node{5, 3}
	f := Node{6, 1}
	g := Node{4, 1}
	h := Node{7, 0}
	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)
	gr.AddNode(f)
	gr.AddNode(g)
	gr.AddNode(h)
	gr.AddNeighbour(a, b, 5)
	gr.AddNeighbour(b, c, 5)
	gr.AddNeighbour(c, d, 5)
	gr.AddNeighbour(c, e, 2)
	gr.AddNeighbour(d, b, 5)
	gr.AddNeighbour(d, g, 2)
	gr.AddNeighbour(e, f, 5)
	gr.AddNeighbour(f, h, 4)
	gr.AddNeighbour(f, g, 5)
	gr.AddNeighbour(g, e, 5)
	gr.AddNeighbour(g, h, 1)
	flow, solution := EdmondsKarp(gr, a, h)
	return flow, solution
}

// Function SolveGraph6 runs the algorithm on the graph 
// sample-problems/problem6;solution-21.jpg
// Note:  This graph follows different rules of flow between nodes, but this
// implementation still provides the correct answer.
func SolveGraph6() (int, *Graph) {
	gr := GenBlankGraph()
	s := Node{0, 2}
	a := Node{1, 4}
	b := Node{3, 4}
	c := Node{2, 3}
	d := Node{1, 2}
	e := Node{3, 2}
	f := Node{2, 1}
	g := Node{1, 0}
	h := Node{3, 0}
	t := Node{4, 2}
	gr.AddNode(s)
	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)
	gr.AddNode(f)
	gr.AddNode(g)
	gr.AddNode(h)
	gr.AddNode(t)
	gr.AddNeighbour(s, a, 8)
	gr.AddNeighbour(s, d, 6)
	gr.AddNeighbour(s, g, 7)
	gr.AddNeighbour(a, b, 10)
	gr.AddNeighbour(a, c, 6)
	gr.AddNeighbour(a, d, 7)
	gr.AddNeighbour(b, c, 5)
	gr.AddNeighbour(b, t, 6)
	gr.AddNeighbour(c, e, 12)
	gr.AddNeighbour(d, c, 9)
	gr.AddNeighbour(d, f, 4)
	gr.AddNeighbour(e, t, 8)
	gr.AddNeighbour(f, e, 3)
	gr.AddNeighbour(f, h, 9)
	gr.AddNeighbour(g, f, 6)
	gr.AddNeighbour(g, h, 5)
	gr.AddNeighbour(h, e, 3)
	gr.AddNeighbour(h, t, 8)
	flow, solution := EdmondsKarp(gr, s, t)
	return flow, solution
}

// Function SolveRandomGraph() runs the algorithm on a random graph.
func SolveRandomGraph() (int, *Graph, Node, Node) {
	gr := GenRandomGraph(25, 75, 50)
	nodelist := gr.GetNodeList()
	source := nodelist[rand.Intn(25)]
	sink := nodelist[rand.Intn(25)]
	flow, solution := EdmondsKarp(gr, source, sink)
	return flow, solution, source, sink
}
