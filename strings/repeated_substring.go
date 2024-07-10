package strings

// RepeatedSubstringPattern 459

// s = "abab"
// s = "aba"
// s = "abcabcabcabc"

func RepeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[0] && 
			len(s) / i > 1 && 
			len(s) % i == 0 {
			sub := s[0 : i]
			b := true
			for j := i; j < len(s); j = j + i {
				if sub != s[j: j+i] {
					b = false
					break;
				}
			}
			if b {
				return b
			}
		}
	}
	return false
}