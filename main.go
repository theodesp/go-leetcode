package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/list"
)

func main()  {
	d:=list.NewDList()
	d.PushHead(1)
	d.PushHead(2)
	d.PushTail(3)
	fmt.Println(d)
}
