package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/graphs"
)

func main() {
	g := graphs.NewUndirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddEdge(1,2,1)
	g.AddEdge(2,3,1)
	g.AddEdge(1,5,1)
	g.AddEdge(3,4,1)
	g.AddEdge(4,5,1)
	g.TraverseBFS(1, func(v int) {
		fmt.Println(v)
	})
}
