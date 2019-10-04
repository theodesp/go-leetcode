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
