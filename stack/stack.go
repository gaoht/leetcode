package stack

type IntStack struct {
	stack []int
	size int
}

func NewStack() *IntStack{
	return &IntStack{
		stack: make([]int, 0),
		size: 0,
	}
}
func (s *IntStack) Push(a int) {
	s.stack = append(s.stack, a)
	s.size++
}

func (s *IntStack) Pop() int {
	e := s.stack[len(s.stack) - 1]
	s.stack = s.stack[:len(s.stack) - 1]
	s.size --
	return e
}

func (s *IntStack) Peek() int {
	return s.stack[len(s.stack) - 1]
}
func (s *IntStack) Top() int{
	e := s.stack[len(s.stack) - 1]
	return e
}

func (s *IntStack) IsEmpty() bool{
	return s.size == 0
}
