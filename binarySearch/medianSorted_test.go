package binarySearch

import (
	"testing"
)

func TestMedianSorted(t *testing.T)  {
	i := FindMedianSortedArrays([]int{1, 3}, []int{})
	if i != 2.0 {
		t.Fail()
	}

	i = FindMedianSortedArrays([]int{1, 3, 7, 9}, []int{})
	if i != 5.0 {
		t.Fail()
	}

	i = FindMedianSortedArrays([]int{}, []int{1,5,6})
	if i != 5 {
		t.Fail()
	}

	i = FindMedianSortedArrays([]int{1, 3}, []int{2})
	if i != 2.0 {
		t.Fail()
	}

	i = FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	if i != 2.5 {
		t.Fail()
	}

	i = FindMedianSortedArrays([]int{2}, []int{})
	if i != 2 {
		t.Fail()
	}

	i = FindMedianSortedArrays([]int{1,2,3,5,6}, []int{4})
	if i != 3.5 {
		t.Fail()
	}

}
