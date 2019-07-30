package uncategorized

import "testing"

func TestAddTwoNumbers(t * testing.T)  {
	if addTwoNumbers(nil, nil) != nil {
		t.Fail()
	}

	if addTwoNumbers(&ListNode{Val: 2}, nil).Val != 2 {
		t.Fail()
	}

	if addTwoNumbers(&ListNode{Val: 5}, &ListNode{Val: 2}).Val != 7 {
		t.Fail()
	}

	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	r := addTwoNumbers(l1, l2)
	if r.Val != 7 && r.Next.Val != 0 && r.Next.Next.Val != 8 {
		t.Fail()
	}
}