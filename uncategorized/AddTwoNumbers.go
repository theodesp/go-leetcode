package uncategorized

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0, Next: nil}
	carry := 0
	left, right := l1, l2
	curr := result

	for left != nil || right != nil {
		var x, y int
		if left != nil {
			x = left.Val
			left = left.Next
		} else {
			x = 0
		}
		if right != nil {
			y = right.Val
			right = right.Next
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10, Next: nil}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry, Next: nil}
	}

	return result.Next
}
