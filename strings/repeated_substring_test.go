package strings

import "testing"

// RepeatedSubstringPattern 459

// s = "abab"
// s = "aba"
// s = "abcabcabcabc"

func TestRepeatedSubstringPattern(t *testing.T)  {
	if !RepeatedSubstringPattern("abab") {
		t.Errorf("abab should be true");
	}
	if RepeatedSubstringPattern("aba") {
		t.Errorf("aba should be false");
	}
	if !RepeatedSubstringPattern("abcabcabcabc") {
		t.Errorf("abcabcabcabc should be true");
	}
}