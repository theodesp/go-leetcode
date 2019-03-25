package binarySearch

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T)  {
	i := Intersect([]int{1,2,2,1}, []int{2,2})
	if !reflect.DeepEqual(i, []int{2, 2}) {
		t.Fatalf("Intersection: got %v", i)
	}

	i = Intersect([]int{4,9,5}, []int{9,4,9,8,4})
	if !reflect.DeepEqual(i, []int{9,4}) {
		t.Fatalf("Intersection: got %v", i)
	}

	i = Intersect([]int{1,2,2,3}, []int{3,2,2,2})
	if !reflect.DeepEqual(i, []int{2, 2, 3}) {
		t.Fatalf("Intersection: got %v", i)
	}

	i = Intersect([]int{3,1,2}, []int{1,3})
	if !reflect.DeepEqual(i, []int{3, 1}) {
		t.Fatalf("Intersection: got %v", i)
	}

}

