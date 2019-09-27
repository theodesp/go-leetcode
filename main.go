package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/stack"
)

func main()  {
	s := stack.NewSortableStack()
	s.Push(4)
	s.Push(3)
	s.Push(6)
	s.Push(9)
	s.Push(1)
	fmt.Println(s.Sort())
}
