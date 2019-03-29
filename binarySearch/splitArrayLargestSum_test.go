package binarySearch

import (
	"testing"
)

func TestSplitArrayLargestSum(t *testing.T)  {
	i := SplitArray([]int{7,2,5,10,8}, 2)
	if i != 18 {
		t.Fail()
	}
}