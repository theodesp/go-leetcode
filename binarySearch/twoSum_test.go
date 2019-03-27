package binarySearch

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T)  {
	i := TwoSum([]int{2,7,11,15}, 9)
	if !reflect.DeepEqual(i, []int{1,2}) {
		t.Fatalf("TwoSum: got %v", i)
	}

	i = TwoSum([]int{3,2,4}, 6)
	if !reflect.DeepEqual(i, []int{1,2}) {
		t.Fatalf("TwoSum: got %v", i)
	}
}
