package list

type LinkedListNode struct {
	next *LinkedListNode
	data int
}

type LinkedList struct {
	Head *LinkedListNode
}

func NewLinkedList() *LinkedList  {
	return &LinkedList{Head: nil}
}

func (l *LinkedList)AddNode(data int) *LinkedListNode {
	node := &LinkedListNode{data: data}
	if l.Head == nil {
		l.Head = node
	} else {
		curr := l.Head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}

	return node
}

type MyLinkedList struct {
	Head *LinkedListNode
	Size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{Size:0}
}


/** Add a node of value val before the first element of the linked list.
After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int)  {
	// Allocate new node
	node := &LinkedListNode{data: val}
	// Point to current head
	node.next = l.Head
	// Set current head as new node
	l.Head = node
	l.Size += 1
}


/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int)  {
	// Allocate new node
	node := &LinkedListNode{data: val}
	// If the Linked List is empty, then make the new node as head
	if l.Head == nil {
		l.Head = node
	} else {
		// Else traverse till the last node
		curr := l.Head
		for curr.next != nil {
			curr = curr.next
		}
		// Change the next of last node
		curr.next = node
	}
	l.Size += 1
}


/** Add a node of value val before the index-th node in the linked list.
If index equals to the length of linked list,
the node will be appended to the end of linked list.
If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > l.Size {
		return
	}
	if index < 0 {
		l.AddAtHead(val)
		return
	}
	// Allocate new node
	node := &LinkedListNode{data: val}
	// If the Linked List is empty, then make the new node as head
	if l.Head == nil {
		l.Head = node
		l.Size += 1
	}
	if index == 0 {
		next := l.Head.next
		l.Head = node
		node.next = next
		l.Size += 1
	} else {
		// Else traverse till the given index node
		curr := l.Head
		var prev *LinkedListNode
		for i:=0;i<index && curr != nil;i+=1{
			prev = curr
			curr = curr.next
		}
		node.next = prev.next
		prev.next = node
	}
	l.Size += 1
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int)  {
	if index > l.Size || l.Size == 0 {
		return
	}
	if index == 0 {
		l.Head = l.Head.next
		l.Size-=1
	}
	curr := l.Head
	// Traverse until previous node. So curr.next in the node to be deleted
	for i:=0;i<index-1 && curr != nil;i+=1{
		curr = curr.next
	}

	// If position is more than number of nodes
	if curr == nil || curr.next == nil {
		return
	}

	// Node curr.next is the node to be deleted
	next := curr.next.next
	curr.next = next
	l.Size-=1
}

/** Get the value of the index-th node in the linked list.
If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index > l.Size || l.Head == nil || index < 0 {
		return -1
	}
	if index == 0 {
		return l.Head.data
	}
	curr := l.Head
	for i:=0;i<index && curr != nil;i+=1 {
		curr = curr.next
	}
	if curr == nil {
		return -1
	}
	return curr.data
}

/** Find out whether a value exists in a singly linked list */
func (l *MyLinkedList) Exists(value int) bool {
	if l.Head == nil {
		return false
	}
	curr := l.Head
	for curr.next != nil {
		if curr.data == value {
			return true
		}
		curr = curr.next
	}
	return false
}

