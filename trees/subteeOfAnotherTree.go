package trees

/*
Given two non-empty binary trees s and t, check whether tree t
has exactly the same structure and node values with a subtree of s.
A subtree of s is a tree consists of a node in s and all of this node's descendants.
The tree s could also be considered as a subtree of itself.
 */

/*
To do this, traverse the binary tree in any way
and check whether the one that it is currently on is
the same as the subtree.
 */

func IsSubtree(s *BinarySearchTreeNode, t *BinarySearchTreeNode) bool {
	// Base cases
	if s == nil {
		return false
	}
	// t should always be equal or bigger than s
	if t == nil {
		return true
	}

	if areSameTrees(s, t) {
		return true
	}

	return IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}

func areSameTrees(s *BinarySearchTreeNode, t *BinarySearchTreeNode) bool  {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return s.value == t.value &&
		areSameTrees(s.Left,t.Left) &&
		areSameTrees(s.Right, t.Right)
}
