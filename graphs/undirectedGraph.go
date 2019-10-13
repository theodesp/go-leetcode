package graphs

type UndirectedGraph struct {
	edges map[int]map[int]int
}

func NewUndirectedGraph() *UndirectedGraph{
	return &UndirectedGraph{edges: make(map[int]map[int]int)}
}
func (g *UndirectedGraph)RemoveVertex(v int)  {
	for key, _ := range g.edges[v] {
		g.RemoveEdge(key, v)
	}
	delete(g.edges, v)
}

func (g *UndirectedGraph)AddVertex(v int)  {
	g.edges[v] = make(map[int]int)
}

func (g *UndirectedGraph)AddEdge(v1, v2, w int)  {
	g.edges[v1][v2] =  w
	g.edges[v2][v1] =  w
}

func (g *UndirectedGraph)RemoveEdge(v1, v2 int)  {
	if _, okV1 := g.edges[v1];okV1 {
		if _, okV2 := g.edges[v1][v2];okV2 {
			delete(g.edges[v1], v2)
		}
	}
	if _, okV2 := g.edges[v2];okV2 {
		if _, okV1 := g.edges[v2][v1];okV1 {
			delete(g.edges[v2], v1)
		}
	}
}

