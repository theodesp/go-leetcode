package list

type ListNode struct {
	Val int
	Next *ListNode
}

/*
 Solution: Iterate over the list using a map of seen values.
 If we have seen the current element before then just delete that element.
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	seen := make(map[int]interface{})
	curr := head
	for curr != nil {
		if _, ok := seen[curr.Val]; ok {
			prev.Next = curr.Next
		} else {
			seen[curr.Val] = struct {}{}
			prev = curr
		}
		curr = prev.Next
	}
	return head
}

