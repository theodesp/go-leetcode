package graphs

type UndirectedGraph struct {
	edges map[int][]int
}

func NewUndirectedGraph() *UndirectedGraph{
	return &UndirectedGraph{edges: make(map[int][]int)}
}
func (g *UndirectedGraph)RemoveVertex(v int)  {
	for key, _ := range g.edges[v] {
		g.RemoveEdge(key, v)
	}
	delete(g.edges, v)
}

func (g *UndirectedGraph)AddVertex(v int)  {
	g.edges[v] = []int{}
}

func (g *UndirectedGraph)AddEdge(v1, v2, w int)  {
	g.edges[v1] = append(g.edges[v1], v2)
	g.edges[v2] = append(g.edges[v2], v1)
}

func indexOf(data []int, element int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func (g *UndirectedGraph)RemoveEdge(v1, v2 int)  {
	if _, ok := g.edges[v1];ok {
		i := indexOf(g.edges[v1], v2)
		if i != - 1 {
			g.edges[v1] = append(g.edges[v1][:i], g.edges[v1][i:]...)
		}
	}
	if _, ok := g.edges[v2];ok {
		i := indexOf(g.edges[v2], v1)
		if i != - 1 {
			g.edges[v2] = append(g.edges[v2][:i], g.edges[v2][i:]...)
		}
	}
}

// Time Complexity: O(V + E)
func (g *UndirectedGraph)TraverseBFS(v int, cb func(v int)) {
	queue := []int{v}
	seen := make(map[int]bool)

	// while the queue is not empty
	for len(queue) > 0 {
		vertex, rest := queue[0], queue[1:]
		queue = rest
		// if we haven't seen that vertex
		for _, val := range g.edges[vertex] {
			if _, ok := seen[val]; !ok {
				seen[val] = true
				cb(val)
				// Push the edges to the queue
				queue = append(queue, val)
			}
		}
	}
}

