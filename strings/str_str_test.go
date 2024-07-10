package strings

import "testing"
func TestStrStr(t *testing.T){
	haystack, needle := "sadbutsad", "sad"
	if Strstr(haystack, needle) != 0 {
		t.Errorf("%s, %s should be %d", haystack, needle, 0)
	}
	haystack, needle = "leetcode", "leeto"
	if Strstr(haystack, needle) != -1 {
		t.Errorf("%s, %s should be %d", haystack, needle, -1)
	}
	haystack, needle = "a", "a"
	if Strstr(haystack, needle) != 0 {
		t.Errorf("%s, %s should be %d", haystack, needle, 0)
	}
}