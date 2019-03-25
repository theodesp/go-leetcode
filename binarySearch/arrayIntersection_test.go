package binarySearch

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T)  {
	i := Intersection([]int{1,2,2,1}, []int{2,2})
	if !reflect.DeepEqual(i, []int{2}) {
		t.Fatalf("Intersection: got %v", i)
	}

	i = Intersection([]int{4,9,5}, []int{9,4,9,8,4})
	if !reflect.DeepEqual(i, []int{9,4}) {
		t.Fatalf("Intersection: got %v", i)
	}

	i = Intersection([]int{2,1}, []int{1,1})
	if !reflect.DeepEqual(i, []int{1}) {
		t.Fatalf("Intersection: got %v", i)
	}

	i = Intersection([]int{3,1,2}, []int{1,3})
	if !reflect.DeepEqual(i, []int{3, 1}) {
		t.Fatalf("Intersection: got %v", i)
	}

}

