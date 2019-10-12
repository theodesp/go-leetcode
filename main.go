package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/binarySearch"
)

func main()  {
	arr := [][]int{
		[]int{10, 8, 10, 10},
		[]int{14, 13, 12, 11},
		[]int{15, 9, 11, 21},
		[]int{16, 17, 19, 20},
	}

	fmt.Println(binarySearch.FindPeakElement2d(arr))
}
