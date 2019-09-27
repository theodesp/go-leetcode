package list

type DList struct {
	head *DListNode
}

type DListNode struct {
	prev *DListNode
	next *DListNode
	data int
}

func NewDList() *DList  {
	return &DList{
		head: nil,
	}
}

func (d *DList)PushHead(val int) {
	//  allocate node
	node := &DListNode{
		prev:nil,
		next:nil,
		data:val,
	}

	// Make next of new node as head and previous as NULL
	node.next = d.head
	node.prev = nil

	// change prev of head node to new node
	if d.head != nil {
		d.head.prev = node
	}
	// move the head to point to the new node
	d.head = node
}

func (d *DList)PushTail(val int) {
	//  allocate node
	node := &DListNode{
		prev:nil,
		next:nil,
		data:val,
	}

	/*
	If the Linked List is empty, then make the new
	          node as head
	 */
	if d.head == nil {
		d.head = node
		return
	}

	// traverse till the last node
	curr := d.head
	for curr.next != nil {
		curr = curr.next
	}
	// Change the current of last node
	curr.next = node

	// Make last node as previous of new node
	node.prev = curr

	return
}