package binarySearch

import (
	"testing"
)

func TestFindPeakElement(t *testing.T)  {
	i := FindPeakElement([]int{1,2,3,1})
	if i != 2 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 2)
	}

	i = FindPeakElement([]int{1,2,1,3,5,6,4})
	if i != 1 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 1)
	}

	i = FindPeakElement([]int{2, 4, 6, 5, 6, 7, 8, 9})
	if i != 2 {
		t.Fatalf("FindPeakElement: got %v, wanted %v", i, 7)
	}
}
