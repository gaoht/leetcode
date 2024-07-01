package strings

import "testing"

func TestReplaceNumber(t *testing.T){
	a := "a1"
	b := ReplaceNumber(a)
	if b != "anumber" {
		t.Errorf("expect anumber, but %s", b)
	}
	a = "a1b2c3"
	b = ReplaceNumber(a)
	if b != "anumberbnumbercnumber" {
		t.Errorf("expect anumberbnumbercnumber, but %s", b)
	}
}