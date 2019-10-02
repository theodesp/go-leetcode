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
	seen := make(map[int]interface{})
	prev := head
	for prev.Next != nil {
		if _, ok := seen[prev.Val]; ok {
			prev.Next = prev.Next.Next
		} else {
			seen[prev.Val] = struct {}{}
			prev = prev.Next
		}
	}
	return head
}

