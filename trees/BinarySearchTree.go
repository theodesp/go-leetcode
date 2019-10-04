package trees

import "math"

/*
Binary Search Tree Implementation
*/

type BinarySearchTreeNode struct {
	value int
	Left, Right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		value: value,
		Left:  nil,
		Right: nil,
	}
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{Root: nil}
}

func (b *BinarySearchTree) Insert(value int)  {
	if b.Root == nil {
		b.Root = NewBinarySearchTreeNode(value)
	} else {
		//loop traverse until
		curr := b.Root
		for  {
			if value > curr.value {
				if curr.Right != nil {
					curr = curr.Right
				} else {
					curr.Left = NewBinarySearchTreeNode(value)
					break
				}
			}
			if value < curr.value {
				if curr.Left != nil {
					curr = curr.Left
				} else {
					curr.Right = NewBinarySearchTreeNode(value)
					break
				}
			} else {
				//case that both are the same
				break
			}
		}
	}
}

/*
Given the root node of a binary search tree (BST) and a value.
You need to find the node in the BST that the node's value equals the given value.
Return the subtree rooted with that node.
If such node doesn't exist, you should return NULL.

For example,

Given the tree:
        4
       / \
      2   7
     / \
    1   3

And the value to search: 2
You should return this subtree:

      2
     / \
    1   3
In the example above, if we want to search the value 5, since there is no node with
value 5, we should return NULL.

Note that an empty tree is represented by NULL,
therefore you would see the expected output (serialized tree format) as [], not null.
 */

func (b BinarySearchTree) Search(value int) *BinarySearchTreeNode  {
	return b.search(b.Root, value)
}


func (b BinarySearchTree) search(node *BinarySearchTreeNode, value int) *BinarySearchTreeNode  {
	if node == nil {
		return nil
	}
	if node.value > value {
		// traverse left
		return b.search(node.Left, value)
	} else if node.value < value {
		// traverse right
		return b.search(node.Right, value)
	} else {
		// found
		return node
	}
}

func (b *BinarySearchTree) LowestCommonAncestor(x,y int) *BinarySearchTreeNode  {
	return b.lowestCommonAncestor(b.Root, x,y)
}

func (b *BinarySearchTree) lowestCommonAncestor(node *BinarySearchTreeNode, x, y int) *BinarySearchTreeNode  {
	if node == nil {
		return nil
	}
	if math.Max(float64(x), float64(y)) < float64(node.value) {
		// If max of x and y is less than node value then LCA is located left of node
		return b.lowestCommonAncestor(node.Left, x, y)
	} else if math.Min(float64(x), float64(y)) > float64(node.value) {
		// If min of x and y is more than node value then LCA is located right of node
		return b.lowestCommonAncestor(node.Right, x, y)
	} else {
		// Otherwise this node is the LCA
		return node
	}
}


func (b *BinarySearchTree) Delete(value int) *BinarySearchTreeNode  {
	return b.delete(b.Root, value)
}

func (b *BinarySearchTree) delete(node *BinarySearchTreeNode, value int) *BinarySearchTreeNode  {
	if node == nil {
		return nil
	} else if value < node.value {
		// Traverse left
		node.Left = b.delete(node.Left, value)
	} else if value > node.value {
		// Traverse right
		node.Right = b.delete(node.Right, value)
	} else {
		// Found!
		// case no child
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			// Bubble up right child
			node = node.Right
			return node
		} else if node.Right == nil {
			// Bubble up left child
			node = node.Left
			return node
		} else {
			/*
			If the node has two children, either find the maximum of the left
			subtree or find the minimum of the right subtree to replace that
			node.
			 */
			temp := findMin(node.Right)
			node.value = temp.value
			// find and remove that node. Then assign it to Right.
			node.Right = b.delete(node.Right, temp.value)
			return node
		}
	}
	return node
}

func findMin(root *BinarySearchTreeNode) *BinarySearchTreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
