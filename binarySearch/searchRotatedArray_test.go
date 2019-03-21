package binarySearch

import (
	"testing"
)

func TestRotatedSearch(t *testing.T)  {
	i := SearchRotated([]int{4,5,6,7,0,1,2}, 0)
	if i != 4 {
		t.Fail()
	}

	i = SearchRotated([]int{4,5,6,7,0,1,2}, 3)
	if i != -1 {
		t.Fail()
	}

	i = SearchRotated([]int{1, 3}, 3)
	if i != 1 {
		t.Fail()
	}
}
