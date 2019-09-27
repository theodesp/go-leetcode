package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/list"
)

func main()  {
	l := list.Constructor()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1,2)
	l.DeleteAtIndex(1)
	fmt.Println(l.Get(-3))
}
