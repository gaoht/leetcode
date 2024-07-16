package stack



type ByteStack struct {
	stack []byte
	size int
}

func NewByteStack() *ByteStack{
	return &ByteStack{
		stack: make([]byte, 0),
		size: 0,
	}
}
func (s *ByteStack) Push(a byte) {
	s.stack = append(s.stack, a)
	s.size++
}

func (s *ByteStack) Pop() byte {
	e := s.stack[len(s.stack) - 1]
	s.stack = s.stack[:len(s.stack) - 1]
	s.size --
	return e
}

func (s *ByteStack) Top() byte{
	e := s.stack[len(s.stack) - 1]
	return e
}

func (s *ByteStack) IsEmpty() bool{
	return s.size == 0
}

// RemoveDuplicates 1047
func RemoveDuplicates(s string) string {
	b := []byte(s)
	t := NewByteStack()
	for _, v := range b {
		if !t.IsEmpty() {
			if v == t.Top() {
				t.Pop()
				continue
			}
		}
		t.Push(v)
	}
	return string(t.stack)
}