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

// Function main currently runs the algorithm on six different graphs.
// As of now all tests return correct solutions.  Very exciting.
func main() {

	start := time.Now()
	flow, solution := graph.SolveGraph1()
	runtime := time.Now().Sub(start)
	fmt.Println("Graph 1\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph2()
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 2\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph3()
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 3\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph4()
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 4\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph5()
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 5\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime, "\n")

	start = time.Now()
	flow, solution = graph.SolveGraph6()
	runtime = time.Now().Sub(start)
	fmt.Println("Graph 6\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime, "\n")

	start = time.Now()
	flow, solution, sink, source := graph.SolveRandomGraph()
	runtime = time.Now().Sub(start)
	list := solution.GetNodeList()
	fmt.Println("Random Graph\nNode List:", list, "\nSource Node:", source, "\nSink Node:", sink, "\nMaximum Flow:", flow, "\nFinal Graph:", solution, "\nTotal Runtime:", runtime)
}
