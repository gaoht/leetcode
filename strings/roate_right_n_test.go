package strings

import "testing"

func TestRotateRightN(t *testing.T){
	s := "abcdefg"
	s = RotateRightN(s, 2)
	if s != "fgabcde" {
		t.Errorf("should be fgabcdef, but %s", s)
	}
}