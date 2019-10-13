package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/graphs"
)

func main() {
	g := graphs.NewUndirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddEdge(1,2,1)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddEdge(2,3,8)
	g.AddEdge(3,4,10)
	g.AddEdge(4,5,100)
	g.AddEdge(1,5,88)
	g.RemoveVertex(5)
	g.RemoveVertex(1)
	g.RemoveEdge(2,3)
	fmt.Println(g)
}
