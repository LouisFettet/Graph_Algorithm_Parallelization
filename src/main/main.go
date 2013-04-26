// main.go
// Executable File for Testing and Example Purposes

// Package main is a simple package for testing the Edmonds-Karp algorithm
// on graphs generated within the graph package.
package main

import (
	"fmt"
	"graph"
	"time"
)

var start time.Time
var flow int
var solution *graph.Graph
var runtime time.Duration
var list []graph.Node
var source graph.Node
var sink graph.Node

// Function main currently runs the algorithm on six different graphs.
// As of now all tests return correct solutions.  Very exciting.
func main() {

	start = time.Now()
	flow, solution = graph.SolveGraph1(false)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 1\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph1(true)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 1\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph2(false)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 2\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph2(true)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 2\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph3(false)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 3\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph3(true)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 3\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph4(false)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 4\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph4(true)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 4\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph5(false)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 5\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph5(true)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 5\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph6(false)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 6\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph6(true)
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 6\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")

	start = time.Now()
	flow, solution, sink, source = graph.SolveRandomGraph(false)
	runtime = time.Now().Sub(start)
	list = solution.GetNodeList()
	fmt.Println("Random Graph\nMaximum Flow:", flow, "\nTotal Runtime (Serial):", runtime, "\n")

	start = time.Now()
	flow, solution, sink, source = graph.SolveRandomGraph(true)
	runtime = time.Now().Sub(start)
	list = solution.GetNodeList()
	fmt.Println("Random Graph\nMaximum Flow:", flow, "\nTotal Runtime (Parallel):", runtime, "\n")
}
