package graphs

type DirectedGraph struct {
	edges map[int][]int
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{edges: make(map[int][]int)}
}

func (g *DirectedGraph) AddVertex(v int)  {
	g.edges[v] = []int{}
}

func (g *DirectedGraph)AddEdge(v1, v2, w int)  {
	g.edges[v1] = append(g.edges[v1], v2)
}

func (g *DirectedGraph)RemoveEdge(v1, v2 int)  {
	if _, ok := g.edges[v1];ok {
		if _, ok := g.edges[v1];ok {
			i := indexOf(g.edges[v1], v2)
			if i != - 1 {
				g.edges[v1] = append(g.edges[v1][:i], g.edges[v1][i:]...)
			}
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

// Time Complexity: O(V + E)
func (g *DirectedGraph)TraverseBFS(v int, cb func(v int)) {
	queue := []int{v}
	seen := make(map[int]bool)

	// while the queue is not empty
	for len(queue) > 0 {
		vertex, rest := queue[0], queue[1:]
		queue = rest
		// if we haven't seen that vertex
		if _, ok := seen[vertex];!ok {
			seen[vertex] = true
			cb(vertex)
			for key, _ := range g.edges[v] {
				// Push the edges to the queue
				queue = append(queue, key)
			}
		}
	}
}

func (g *DirectedGraph)TopologicalSort() []string {
	seen := make(map[int]bool)
	stack := []string{}
	for key, _ := range g.edges {
		if _, ok := seen[key]; !ok {
			g.TopologicalSortUntil(key, seen, &stack)
		}
	}
	return stack
}

func (g *DirectedGraph)TopologicalSortUntil(v int, seen map[int]bool, stack *[]string) {
	seen[v] = true
	for _, val := range g.edges[v] {
		if _, ok := seen[val]; !ok {
			g.TopologicalSortUntil(val, seen, stack)
		}
	}
	*stack = append([]string{string(v)}, *stack...)
}