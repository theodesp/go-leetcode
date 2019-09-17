package uncategorized

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{9, nil, nil}}}
	if !reflect.DeepEqual(AverageOfLevels(tree), []float64{3, 14.5, 11}) {
		t.Fail()
	}
}
