package main

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
