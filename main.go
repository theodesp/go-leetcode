package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/caching"
)

func main()  {
	l := caching.NewLRUCache(2)
	l.Set(1,1)
	l.Set(2,2)
	fmt.Println(l.Get(1))
	l.Set(3,3)
	fmt.Println(l.Get(2))
	l.Set(4,4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}
