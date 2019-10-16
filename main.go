package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/graphs"
)

func main() {
	g := graphs.NewDirectedGraph()
	g.AddVertex('A')
	g.AddVertex('B')
	g.AddVertex('C')
	g.AddVertex('D')
	g.AddVertex('E')
	g.AddVertex('F')
	g.AddEdge('B','A',1)
	g.AddEdge('D','C',1)
	g.AddEdge('D','B',1)
	g.AddEdge('B','A',1)
	g.AddEdge('A','F',1)
	g.AddEdge('E','C',1)
	s := g.TopologicalSort()
	fmt.Println(s)
}
