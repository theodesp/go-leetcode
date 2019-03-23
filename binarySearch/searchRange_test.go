package binarySearch

import (
	"reflect"
	"testing"
)

func TestSearchTange(t *testing.T)  {
	i := SearchRange([]int{5,7,7,8,8,10}, 8)
	if !reflect.DeepEqual(i, []int{3,4}) {
		t.Fatalf("expected %v: got: %v", []int{3,4}, i)
	}

	i = SearchRange([]int{5,7,7,8,8,10}, 6)
	if !reflect.DeepEqual(i, []int{-1, -1}) {
		t.Fatalf("expected %v: got: %v", []int{-1, -1}, i)
	}

	i = SearchRange([]int{2, 2}, 2)
	if !reflect.DeepEqual(i, []int{0, 1}) {
		t.Fatalf("expected %v: got: %v", []int{0, 1}, i)
	}
}
