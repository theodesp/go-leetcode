package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/caching"
)

func main()  {
	l := caching.NewLFUCache(5)
	l.Put(1,1)
	l.Put(2,2)
	l.Put(3,3)
	l.Put(4,4)
	l.Put(5,5)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(1))
	l.Put(6,6)
	fmt.Println(l.Get(6))
}
