package leetcode
// [541] 反转字符串 II


func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i, n := 0, 0; i < len(bytes); i, n = i + k, n + 1{
		if n % 2 == 0 {
			end := i + k - 1 
			if end >= len(bytes) {
				end = len(bytes) - 1
			}
			reverseRangedStr(bytes, i, end)
		}
	}
	return string(bytes)
}
func reverseRangedStr(bytes []byte, start, end int) {
	for left, right := start, end; left < right; left, right = left + 1, right - 1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}
}

 

