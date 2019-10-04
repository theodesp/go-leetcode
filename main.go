package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/trees"
)

func main()  {
	bt := trees.NewBinaryTree()
	bt.Root = trees.NewBinaryTreeNode(42)
	bt.Root.Left = trees.NewBinaryTreeNode(41)
	bt.Root.Right = trees.NewBinaryTreeNode(50)
	bt.Root.Left.Left = trees.NewBinaryTreeNode(10)
	bt.Root.Left.Right = trees.NewBinaryTreeNode(40)
	bt.Root.Right.Left = trees.NewBinaryTreeNode(45)
	bt.Root.Right.Right = trees.NewBinaryTreeNode(75)
	bt.TraversePreorder(func(value interface{}) {
		fmt.Println(value)
	})
}
