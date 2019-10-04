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
	curr := head
	for curr.Next != nil {
		if _, ok := seen[curr.Val]; ok {
			curr.Next = curr.Next.Next
		} else {
			seen[curr.Val] = struct {}{}
			curr = curr.Next
		}
	}
	return head
}

