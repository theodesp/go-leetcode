package trees

/*
Binary Tree Implementation
 */

type BinaryTreeNode struct {
	value interface{}
	Left, Right *BinaryTreeNode
}

func NewBinaryTreeNode(value interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		value: value,
		Left:  nil,
		Right: nil,
	}
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}
// Traverse recursively
func (bt *BinaryTree)TraversePreorder(cb func(value interface{})) {
	Preorder(bt.Root, cb)
}
// Traverse iterate
func (bt *BinaryTree)TraversePreorderIter(cb func(value interface{})) {
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

func Preorder(node *BinaryTreeNode, cb func(value interface{})) {
	if node == nil {
		return
	}
	cb(node.value)
	Preorder(node.Left, cb)
	Preorder(node.Right, cb)
}

// Traverse recursively
func (bt *BinaryTree)TraverseInorder(cb func(value interface{})) {
	Inorder(bt.Root, cb)
}

func Inorder(node *BinaryTreeNode, cb func(value interface{})) {
	if node == nil {
		return
	}
	Preorder(node.Left, cb)
	cb(node.value)
	Preorder(node.Right, cb)
}

// Traverse iterate
func (bt *BinaryTree)TraverseInorderIter(cb func(value interface{})) {
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
func (bt *BinaryTree)TraversePostorder(cb func(value interface{})) {
	Postorder(bt.Root, cb)
}

func Postorder(node *BinaryTreeNode, cb func(value interface{})) {
	if node == nil {
		return
	}
	Preorder(node.Left, cb)
	Preorder(node.Right, cb)
	cb(node.value)
}

// Traverse iterate
func (bt *BinaryTree)TraversePostorderIter(cb func(value interface{})) {
	// Create two stacks
	s1 := []*BinaryTreeNode{}
	s2 := []*BinaryTreeNode{}

	// Push root to first stack
	s1 = append(s1, bt.Root)

	// Run while first stack is not empty
	for len(s1) > 0 {
		// Pop an item from s1 and append it to s2
		node := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, node)

		// Push left and right children of removed item to s1
		if node.Left != nil {
			s1 = append(s1, node.Left)
		}
		if node.Right != nil {
			s1 = append(s1, node.Right)
		}
	}

	// Print all elements of second stack
	for len(s2) > 0 {
		node := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		cb(node.value)
	}
}

type NodeLevel struct {
	Node *BinaryTreeNode
	Level int
}

// Traverse recursively
func (bt *BinaryTree)TraverseLevelorder() [][]int {
	result := [][]int{}
	level := 0
	// Use a queue as FIFO access
	queue := []NodeLevel{}

	if bt.Root == nil {
		return result
	}

	queue = append(queue, NodeLevel{Node:bt.Root, Level:level})
	result = append(result, []int{})
	for len(queue) > 0 {
		// Grab first element
		curr, rest := queue[0], queue[1:]
		queue = rest
		if len(result)-1 < curr.Level {
			result = append(result, []int{})
		}
		result[curr.Level] = append(result[curr.Level], curr.Node.value.(int))
		// Push left and right children into the queue
		if curr.Node.Left != nil {
			queue = append(queue, NodeLevel{Node:curr.Node.Left, Level:curr.Level+1})
		}
		if curr.Node.Right != nil {
			queue = append(queue, NodeLevel{Node:curr.Node.Right, Level:curr.Level+1})
		}
	}

	return result
}
