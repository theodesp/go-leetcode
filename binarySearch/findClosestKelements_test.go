package binarySearch

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T)  {
	i := FindClosestElements([]int{1,2,3,4,5}, 2, 3)
	if !reflect.DeepEqual(i, []int{2,3}) {
		t.Fatalf("expected %v: got: %v", []int{2,3}, i)
	}

	i = FindClosestElements([]int{0,0,1,2,3,3,4,7,7,8}, 3, 5)
	if !reflect.DeepEqual(i, []int{3,3,4}) {
		t.Fatalf("expected %v: got: %v", []int{3,3,4}, i)
	}

	i = FindClosestElements([]int{0,0,0,1,3,5,6,7,8,8}, 2, 2)
	if !reflect.DeepEqual(i, []int{1,3}) {
		t.Fatalf("expected %v: got: %v", []int{1,3}, i)
	}

	i = FindClosestElements([]int{0,1,1,1,2,3,6,7,8,9}, 9, 4)
	if !reflect.DeepEqual(i, []int{0,1,1,1,2,3,6,7,8}) {
		t.Fatalf("expected %v: got: %v", []int{0,1,1,1,2,3,6,7,8}, i)
	}
}
