package main

//import "fmt"

type WEdge struct {
	src    int
	dest   int
	weight int
}

func NewWEdge(src, dest, weight int) WEdge {
	return WEdge{src, dest, weight}
}

type WGraph struct {
	graph []([]WEdge)
}

func NewWGraph(vertices int) *WGraph {
	graph := make([]([]WEdge), vertices)

	for i := range graph {
		graph[i] = make([]WEdge, 0)
	}
	return &WGraph{graph}
}

func (g *WGraph) AddWEdge(src, dest, weight int) {
	edge := NewWEdge(src, dest, weight)
	g.graph[src] = append(g.graph[src], edge)
}

func createGraph(vertices int) WGraph {
	graph := NewWGraph(vertices)

	graph.AddWEdge(0, 2, 2)

	graph.AddWEdge(1, 2, 10)
	graph.AddWEdge(1, 3, 0)

	graph.AddWEdge(2, 0, 2)
	graph.AddWEdge(2, 1, 10)
	graph.AddWEdge(2, 3, -1)

	graph.AddWEdge(3, 1, 0)
	graph.AddWEdge(3, 2, -1)

	return *graph
}

// func main() {
// 	V := 4

// 	graph := createGraph(V)

// 	//print 2's neighbours

// 	for i := 0; i < len(graph.graph[2]); i++ {
// 		edge := graph.graph[2][i]
// 		fmt.Printf("%d %d\n", edge.dest, edge.weight)
// 	}

// }
