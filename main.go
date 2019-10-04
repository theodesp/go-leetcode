package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/hashTable"
)

func main()  {
	ht := hashTable.NewHashTable(10)
	ht.PutDoubleHash(7, "h1")
	ht.PutDoubleHash(20, "hello")
	ht.PutDoubleHash(33, "sunny")
	ht.PutDoubleHash(46, "weather")
	ht.PutDoubleHash(59, "wow")
	ht.PutDoubleHash(72, "forty")

	fmt.Println(ht)
}
