package list

func (l *List) Middle() interface{}  {
	var slow, fast *Node
	slow = l.head
	fast = l.head.next
	for fast != nil {
		if fast.next == nil {
			fast = fast.next
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
	return slow.data
}

//0 -> 1 -> 2 -> nil
