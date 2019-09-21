package trees

/*
Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/* Solution 1
  Save left = Right Subtree
  Save right = Left Subtree

  Return root with left and right swapped.
  Continue recursively
 */

func InvertTree(root *TreeNode) *TreeNode  {
	if root == nil {
		return root
	}
	right := InvertTree(root.Right)
	left := InvertTree(root.Left)
	root.Left = right
	root.Right = left

	return root
}

/* Solution 2
Using BFS. Use a temp node to swap the left and right nodes. Continue until we have nodes.
*/

func InvertTreeBFS(root *TreeNode) *TreeNode  {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr, rest := queue[0], queue[1:]
		temp := curr.Left
		curr.Left = curr.Right
		curr.Right = temp
		queue = rest
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return root
}
