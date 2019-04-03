package list

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
}

func NewNode(data interface{}) *Node  {
	return &Node{
		data: data,
	}
}

func NewList() *List  {
	return &List{}
}

func (l *List) Push(data interface{})  {
	n := NewNode(data)
	n.next = l.head
	l.head = n
}

func (l *List) Reverse()  {
	var prev, next *Node
	curr := l.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *List) ForEach(callBack func(data interface{}))  {
	t := l.head
	for t != nil {
		callBack(t.data)
		t = t.next
	}
}

