package binarySearch

import (
	"testing"
)

func TestFindMin(t *testing.T)  {
	i := FindMin([]int{4,4,4,5,6,1,2})
	if i != 1 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 1)
	}

	i = FindMin([]int{1,1,1,0,1,1})
	if i != 0 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 0)
	}

	i = FindMin([]int{2,2,2,0,1})
	if i != 0 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 0)
	}
}
