package stack

import "testing"

func TestStask(t *testing.T) {
	s := NewStack()
	s.Push(1)
	if s.size != 1 {
		t.Errorf("size should 1")
	}
	e := s.Pop()
	if e != 1 {
		t.Errorf("pop should 1")
	}
	if !s.IsEmpty() {
		t.Errorf("stack should be empty")
	}
}