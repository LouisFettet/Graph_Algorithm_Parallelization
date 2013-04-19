// main.go
// Executable File for Testing and Example Purposes

// Package main is a simple package for testing the Edmonds-Karp algorithm
// on graphs generated within the graph package.
package main

import (
	"fmt"
	"graph"
)

// Function main currently runs the algorithm on six different graphs.
// As of now all tests return correct solutions.  Very exciting.
func main() {
	flow, solution := graph.SolveGraph1()
	fmt.Println(flow, solution)
	flow, solution = graph.SolveGraph2()
	fmt.Println(flow, solution)
	flow, solution = graph.SolveGraph3()
	fmt.Println(flow, solution)
	flow, solution = graph.SolveGraph4()
	fmt.Println(flow, solution)
	flow, solution = graph.SolveGraph5()
	fmt.Println(flow, solution)
	flow, solution = graph.SolveGraph6()
	fmt.Println(flow, solution)
}
