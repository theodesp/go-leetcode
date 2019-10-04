package trees

/*
Binary Tree Implementation
 */

type BinaryTreeNode struct {
	value interface{}
	Left, Right *BinaryTreeNode
}

func NewBinaryTreeNode(value interface{}) *BinaryTreeNode  {
	return &BinaryTreeNode{
		value: value,
		Left:  nil,
		Right: nil,
	}
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree  {
	return &BinaryTree{Root: nil}
}
// Traverse recursively
func (bt *BinaryTree)TraversePreorder(cb func(value interface{}))  {
	Preorder(bt.Root, cb)
}
// Traverse iterate
func (bt *BinaryTree)TraversePreorderIter(cb func(value interface{}))  {
	stack := []*BinaryTreeNode{bt.Root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cb(node)
		// push right first and then left so left will be popped up first
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func Preorder(node *BinaryTreeNode, cb func(value interface{}))  {
	if node == nil {
		return
	}
	cb(node.value)
	Preorder(node.Left, cb)
	Preorder(node.Right, cb)
}

// Traverse recursively
func (bt *BinaryTree)TraverseInorder(cb func(value interface{}))  {
	Inorder(bt.Root, cb)
}

func Inorder(node *BinaryTreeNode, cb func(value interface{}))  {
	if node == nil {
		return
	}
	Preorder(node.Left, cb)
	cb(node.value)
	Preorder(node.Right, cb)
}

// Traverse iterate
func (bt *BinaryTree)TraverseInorderIter(cb func(value interface{}))  {
	// Start at root
	curr := bt.Root
	stack := []*BinaryTreeNode{}
	done := false
	for !done {
		if curr != nil {
			// push root into stack
			stack = append(stack, curr)
			// Traverse left until we have no left trees left
			curr = curr.Left
		} else {
			// As long as we have elements in the stack
			if len(stack) > 0 {
				// Pop the last stack element pushed.
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// Inorder callback
				cb(node.value)
				// Traverse Right now
				node = node.Right
			} else {
				done = true
			}
		}
	}
}

// Traverse recursively
func (bt *BinaryTree)TraversePostorder(cb func(value interface{}))  {
	Postorder(bt.Root, cb)
}

func Postorder(node *BinaryTreeNode, cb func(value interface{}))  {
	if node == nil {
		return
	}
	Preorder(node.Left, cb)
	Preorder(node.Right, cb)
	cb(node.value)
}
