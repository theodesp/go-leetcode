package stack


type Stack struct {
	queue []int
}

func NewStack() *Stack  {
	return &Stack{queue: []int{}}
}

func (s *Stack) Push(val int)  {
	s.queue = append([]int{val}, s.queue...)
}

func (s *Stack) Pop() int  {
	if s.IsEmpty() {
		return -1
	}
	for i:=0;i< len(s.queue); i+=1 {
		first, rest := s.queue[0], s.queue[1:]
		s.queue = rest
		s.queue = append(s.queue, first)
	}
	first, rest := s.queue[0], s.queue[1:]
	s.queue = rest
	return first
}

func (s *Stack) IsEmpty() bool  {
	return len(s.queue) == 0
}

