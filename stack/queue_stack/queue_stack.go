package stack

// MyStack 225
type MyStack struct {
	m *Queue
	n *Queue
}

func Constructor() MyStack {
	return MyStack{
		m: NewQueue(),
		n: NewQueue(),
	}
}


func (s *MyStack) Push(x int)  {
	if !s.m.IsEmpty() {
		s.m.Push(x)
	} else {
		s.n.Push(x)
	}
}


func (s *MyStack) Pop() int {
	m, n := s.m, s.n
	if s.m.IsEmpty() {
		m, n = s.n, s.m
	}
	for m.Size() > 1 {
		n.Push(m.Pop())
	}
	return m.Pop()
}


func (s *MyStack) Top() int {
	m, n := s.m, s.n
	if s.m.IsEmpty() {
		m, n = s.n, s.m
	}
	for m.Size() > 1 {
		n.Push(m.Pop())
	}
	e := m.Peek()
	n.Push(m.Pop())
	return e
}


func (s *MyStack) Empty() bool {
	return s.m.IsEmpty() && s.n.IsEmpty()
}


