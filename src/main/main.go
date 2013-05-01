// main.go
// Executable File for Testing and Example Purposes

// Package main is a simple package for testing the Edmonds-Karp algorithm
// on graphs generated within the graph package.
package main

import (
	"fmt"
	"graph"
	"math/rand"
	"time"
)

var start time.Time
var flow int
var solution *graph.Graph
var runtime time.Duration
var list []graph.Node
var source graph.Node
var sink graph.Node
var total time.Duration
var n int

// Function main currently runs the algorithm on six different graphs and one 
// "random graph." 
func main() {

	n = 100

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph1(false)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Graph 1\nMax Flow(Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph1(true)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph2(false)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Graph 2\nMax Flow (Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph2(true)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")
	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph3(false)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Graph 3\nMax Flow (Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph3(true)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph4(false)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Graph 4\nMax Flow (Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph4(true)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph5(false)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Graph 5\nMax Flow (Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph5(true)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph6(false)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Graph 6\nMax Flow (Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution = graph.SolveGraph6(true)
		runtime = time.Now().Sub(start)
		total += runtime
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")

	total = 0

	g := graph.GenRandomGraph(1000, 3000, 100)
	nodelist := g.GetNodeList()
	s := nodelist[rand.Intn(1000)]
	t := nodelist[rand.Intn(1000)]
	for s == t {
		t = nodelist[rand.Intn(1000)]
	}

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution, sink, source = graph.SolveRandomGraph(g, s, t, false)
		runtime = time.Now().Sub(start)
		total += runtime
		// Reset random graph...
		for _, node := range nodelist {
			for _, neighbour := range g.GetNeighbours(node) {
				g.UpdateWeight(node, neighbour.Neighbour_Node, 0)
			}
		}
	}

	fmt.Println("Random Graph\nMax Flow(Serial):", flow, "\nAverage Runtime (Serial):", (int64(total) / int64(n) / 1000), "us")

	total = 0

	for i := 0; i < n; i++ {
		start = time.Now()
		flow, solution, sink, source = graph.SolveRandomGraph(g, s, t, true)
		runtime = time.Now().Sub(start)
		total += runtime
		// Reset random graph...
		for _, node := range nodelist {
			for _, neighbour := range g.GetNeighbours(node) {
				g.UpdateWeight(node, neighbour.Neighbour_Node, 0)
			}
		}
	}
	fmt.Println("Max Flow(Concurrent):", flow, "\nAverage Runtime (Concurrent):", (int64(total) / int64(n) / 1000), "us\n")

}
