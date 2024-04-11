package main

import "fmt"

type Edge struct {
	src  int
	dest int
}

func NewEdge(src, dest int) Edge {
	return Edge{src, dest}
}

type Graph struct {
	graph []([]Edge)
}

func NewGraph(vertices int) *Graph {
	graph := make([]([]Edge), vertices)
	for i := range graph {
		graph[i] = make([]Edge, 0)
	}
	return &Graph{graph}
}

func (g *Graph) AddEdge(src, dest int) {
	edge := NewEdge(src, dest)
	g.graph[src] = append(g.graph[src], edge)
}

func main() {
	vertices := 5
	graph := NewGraph(vertices)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 1)
	graph.AddEdge(3, 2)
	graph.AddEdge(4, 3)

	for i := 0; i < vertices; i++ {
		fmt.Printf("Vertex %d: ", i)
		for _, edge := range graph.graph[i] {
			fmt.Printf("(%d -> %d)", edge.src, edge.dest)
		}
	}
}
