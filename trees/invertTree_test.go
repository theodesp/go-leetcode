package trees

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := &TreeNode{4,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil}},
			&TreeNode{7,
				&TreeNode{6,nil,nil},
				&TreeNode{9, nil,nil}}}

	inverted := InvertTreeBFS(root)
	if !reflect.DeepEqual(inverted, &TreeNode{4,nil,nil}) {
		t.Fail()
	}
}