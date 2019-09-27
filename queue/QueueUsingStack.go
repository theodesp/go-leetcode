package queue

/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.
 */
type S struct {
	items []int
}

func (s *S)Push(val int)  {
	s.items = append(s.items, val)
}

func (s *S)Pop()int  {
	if len(s.items) == 0 {
		return -1
	}
	last, butLast := s.items[len(s.items)-1], s.items[0:len(s.items)-1]
	s.items = butLast
	return last
}

func (s S)Peek() int  {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func (s S)IsEmpty()bool  {
	return len(s.items) == 0
}

type MyQueue struct {
	stack *S
	first int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: &S{},
		first: -1,
	}
}


/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int)  {
	if q.stack.IsEmpty() {
		q.first = x
	}
	q.stack.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.stack.IsEmpty() {
		return -1
	}
	// // pop an item from the stack
	item := q.stack.Pop()

	// if stack becomes empty, return
	// the popped item
	if q.stack.IsEmpty() {
		return item
	}
	// recursive call
	x := q.Pop()

	// push popped item back to the stack
	q.Push(item)

	// return the result of deQueue() call
	return x
}


/** Get the front element. */
func (q *MyQueue) Peek() int {
	return q.first
}


/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.stack.IsEmpty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */