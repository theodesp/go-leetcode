package main

import (
	"fmt"
	queue2 "github.com/theodesp/go-leetcode/queue"
)

func main()  {
	queue := queue2.Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())   // returns 1
	queue.Empty()
}
