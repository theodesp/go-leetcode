package graphs

type DirectedGraph struct {
	edges map[int]map[int]int
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{edges: make(map[int]map[int]int)}
}

func (g *DirectedGraph) AddVertex(v int)  {
	g.edges[v] = make(map[int]int)
}

func (g *DirectedGraph)AddEdge(v1, v2, w int)  {
	g.edges[v1][v2] =  w
}

func (g *DirectedGraph)RemoveEdge(v1, v2 int)  {
	if _, okV1 := g.edges[v1];okV1 {
		if _, okV2 := g.edges[v1][v2];okV2 {
			delete(g.edges[v1], v2)
		}
	}
}

func (g *DirectedGraph)RemoveVertex(v int)  {
	for key, _ := range g.edges[v] {
		// Remove anything that points to this vertex
		g.RemoveEdge(key, v)
	}
	// Finally remove vertex itself
	delete(g.edges, v)
}
