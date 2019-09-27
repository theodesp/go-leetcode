package stack

import "github.com/theodesp/go-leetcode/queue"

type SortableStack struct {
	unsorted *queue.S
}

func NewSortableStack() *SortableStack  {
	return &SortableStack{
		&queue.S{},
	}
}

func (s *SortableStack)Push(val int)  {
	s.unsorted.Push(val)
}

func (s *SortableStack)Pop()int  {
	return s.unsorted.Pop()
}

func (s SortableStack)Peek() int  {
	return s.unsorted.Peek()
}

func (s SortableStack)IsEmpty()bool  {
	return s.unsorted.IsEmpty()
}

func (s *SortableStack)Sort() *queue.S {
	sorted := &queue.S{}
	for !s.unsorted.IsEmpty() {
		// Pop an element from unsorted stack call it temp
		temp := s.unsorted.Pop()
		/*
		While sorted stack is NOT empty and top of temporary stack is greater than temp,
		pop from temporary stack and push it to the input stack
		 */
		for !sorted.IsEmpty() && sorted.Peek() > temp {
			s.unsorted.Push(sorted.Pop())
		}
		sorted.Push(temp)
	}

	return sorted

}


