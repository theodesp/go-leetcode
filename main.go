package main

import (
	"github.com/theodesp/go-leetcode/trees"
)

func main()  {
	t := trees.NewBinarySearchTree()
	t.InsertRecur(1)
	t.InsertRecur(3)
	t.InsertRecur(2)
	t.InsertRecur(4)
	t.InsertRecur(5)
	t.InsertRecur(8)
	t.InsertRecur(9)
	trees.PrintNodesKDistanceRoot(t.Root, 2)
}
