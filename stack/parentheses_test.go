package stack

import "testing"

// IsValid 20 括号匹配
func TestIsValid(t *testing.T)  {
	s := "()"
	if !IsValid(s) {
		t.Errorf("() should be true");
	}
	s = "()[]{}"
	if !IsValid(s) {
		t.Errorf("()[]{} should be true");
	}
	s = "(]"
	if IsValid(s) {
		t.Errorf("(] should be true");
	}

}