package trees

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