package binarySearch

import (
	"testing"
)

func TestSmallestDistancePairs(t *testing.T)  {
	i := SmallestDistancePair([]int{1,3,1}, 3)
	if i != 2 {
		t.Fail()
	}
}