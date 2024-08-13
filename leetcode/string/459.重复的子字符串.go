package leetcode
// [459] 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	j := 1
	sLen := len(s)
	for  j <= sLen / 2 {
		// 不能整除
		if sLen / j * j != sLen {
			j++
			continue
		}
		for i := j; i + j <= sLen; i += j {
			if s[ : j] != s[i : i+j] {
				break
			}
			if i + j == sLen {
				return true
			}
		}
		j++
	}
	return false
}

