package main

import (
	"fmt"
)

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

type WEdge struct {
	src    int
	dest   int
	weight int
}

func NewWEdge(src, dest, weight int) WEdge {
	return WEdge{src, dest, weight}
}

type WGrpah struct {
	graph []([]WEdge)
}

func NewWGraph(vertices int) *WGrpah {
	graph := make([]([]WEdge), vertices)

	for i := range graph {
		graph[i] = make([]WEdge, 0)
	}
	return &WGrpah{graph}
}

func (g *WGrpah) AddWEdge(src, dest, weight int) {
	edge := NewWEdge(src, dest, weight)
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

	// find neighbours 1
	fmt.Println()
	for _, edge := range graph.graph[1] {
		fmt.Printf(" %d ", edge.dest)
	}
	fmt.Println()

	//weighted adjacency list
	vertices = 4
	wgraph := NewWGraph(vertices)
	wgraph.AddWEdge(0, 2, 2)
	wgraph.AddWEdge(1, 2, 10)
	wgraph.AddWEdge(1, 3, 0)
	wgraph.AddWEdge(2, 0, 2)
	wgraph.AddWEdge(2, 1, 10)
	wgraph.AddWEdge(2, 3, -1)
	wgraph.AddWEdge(3, 1, 0)
	wgraph.AddWEdge(3, 2, -1)

	fmt.Println()
	for _, edge := range wgraph.graph[2] {
		fmt.Printf(" %d ", edge.dest)
	}
	fmt.Println()
}
