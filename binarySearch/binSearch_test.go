package binarySearch

import (
	"testing"
)

func TestBinSearch(t *testing.T)  {
	i := Search([]int{-1,0,3,5,9,12}, 9)
	if i != 4 {
		t.Fail()
	}

	i = Search([]int{2,5}, 5)
	if i != 1 {
		t.Fail()
	}
}
