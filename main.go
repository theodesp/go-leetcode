package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/array"
)

func main()  {
	arrs := [][]int{
		[]int{1,2,3},
		[]int{1,2,3,4},
		[]int{1,2},
	}
	fmt.Println(array.FindCommonElementsKSortedArrays(arrs))
}
