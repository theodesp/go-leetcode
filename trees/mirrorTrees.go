package trees

/*
Given two Binary Trees, write a function that returns true if
two trees are mirror of each other, else false.

For two trees ‘a’ and ‘b’ to be mirror images, the following three conditions must be true:

Their root node’s key must be same
Left subtree of root of ‘a’ and right subtree root of ‘b’ are mirror.
Right subtree of ‘a’ and left subtree of ‘b’ are mirror.
 */

func AreMirrorTrees(a, b *BinarySearchTreeNode) bool  {
	// Base cases
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	/* Both non-empty, compare them recursively
	   Note that in recursive calls, we pass left
	   of one tree and right of other tree */
	return a.value == b.value &&
		AreMirrorTrees(a.Left, b.Right) &&
		AreMirrorTrees(a.Right, b.Left)
}
