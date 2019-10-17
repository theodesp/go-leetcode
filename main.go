package main

import (
	"fmt"
	"github.com/theodesp/go-leetcode/dynamicProgramming"
)

func main() {
	fmt.Println(dynamicProgramming.KnapSackDP(50, 3, []int{10,20,30}, []int{60, 100, 120}))
}
