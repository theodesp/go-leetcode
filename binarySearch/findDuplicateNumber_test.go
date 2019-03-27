package binarySearch

import (
	"testing"
)

func TestFindDuplicateNumber(t *testing.T)  {
	i := FindDuplicate([]int{3,1,3,4,2})
	if i != 3 {
		t.Fail()
	}

	i = FindDuplicate([]int{1,3,4,2,2})
	if i != 2 {
		t.Fail()
	}

	i = FindDuplicate([]int{1,3,4,2,1})
	if i != 1 {
		t.Fail()
	}
}
