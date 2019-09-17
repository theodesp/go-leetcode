package uncategorized
//
///*
//Given a non-empty binary tree, return the average value of the nodes on
//each level in the form of an array.
//Example 1:
//Input:
//    3
//   / \
//  9  20
//    /  \
//   15   7
//Output: [3, 14.5, 11]
//Explanation:
//The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on
//level 2 is 11. Hence return [3, 14.5, 11].
//Note:
//The range of node's value is in the range of 32-bit signed integer.
//*/
//
///*
// Solution 1: Do a BFS on the tree but store it like [[3],[9,20],[15,7]]
// Then go and compute the average on each array.
//*/
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func AverageOfLevels(root *TreeNode) []float64 {
//	arr := [][]int{}
//	result := []float64{}
//	curr := root
//
//	for curr.Left != nil && curr.Right != nil {
//		arr = append(arr, curr.Val)
//	}
//
//	for _, el := range arr {
//		result = append(result, average(el))
//	}
//	return result
//}
//
//
//func average(arr[]int) float64  {
//	sum := int64(0)
//	for _, el := range arr {
//		sum += int64(el)
//	}
//	return float64(sum / int64(len(arr)))
//}
