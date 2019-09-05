package uncategorized

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(TwoSum([]int{2, 7, 11, 15}, 9), []int{0,1}) {
		t.Fail()
	}
	if !reflect.DeepEqual(TwoSum([]int{3,2,4}, 6), []int{1,2}) {
		t.Fail()
	}
}