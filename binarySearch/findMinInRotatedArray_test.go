package binarySearch

import (
	"testing"
)

func TestFindMinInRotatedArray(t *testing.T)  {
	i := FindMinInRotatedArray([]int{4,5,6,7,0,1,2})
	if i != 0 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 0)
	}

	i = FindMinInRotatedArray([]int{3,4,5,1,2})
	if i != 1 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 1)
	}

	i = FindMinInRotatedArray([]int{3, 4, 6, 1, 2})
	if i != 1 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 1)
	}
}
