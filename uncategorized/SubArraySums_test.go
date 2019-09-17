package uncategorized

import "testing"

func TestSubarraySum(t *testing.T) {
	if SubarraySum([]int{1,1,1}, 2) != 2 {
		t.Fail()
	}
}