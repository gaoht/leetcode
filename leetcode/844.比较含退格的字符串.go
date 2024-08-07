package leetcode
// [844] 比较含退格的字符串

func backspaceCompare(s string, t string) bool {
	return string(backspace([]byte(s))) ==  string(backspace([]byte(t)))
}

func backspace(b []byte) []byte{
	slow := 0
	for fast := 0; fast < len(b); fast++{
		if b[fast] == '#' {
			slow --
			if slow < 0 {
				slow = 0
			}
		} else {
			b[slow] = b[fast]
			slow++	
		}
	}
	return b[:slow]
}
// @lc code=end

