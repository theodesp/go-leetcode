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
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
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
}


/** Add a node of value val before the index-th node in the linked list.
If index equals to the length of linked list,
the node will be appended to the end of linked list.
If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int)  {
	// Allocate new node
	node := &LinkedListNode{data: val}
	// If the Linked List is empty, then make the new node as head
	if l.Head == nil && index == 0 {
		l.Head = node
	} else {
		count := 0
		// Else traverse till the last node
		curr := l.Head
		for curr.next != nil && count < index {
			curr = curr.next
			count += 1
		}
		if count == index {

		}
		// Change the next of last node
		curr.next = node
	}
}


///** Delete the index-th node in the linked list, if the index is valid. */
//func (l *MyLinkedList) DeleteAtIndex(index int)  {
//
//}
//
///** Get the value of the index-th node in the linked list.
//If the index is invalid, return -1. */
//func (l *MyLinkedList) Get(index int) int {
//
//}
