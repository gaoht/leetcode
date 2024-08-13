package leetcode

// [151] 反转字符串中的单词

func reverseWords(s string) string {
	bytes := []byte(s)
	bytes = clearSpace(bytes)
	reverseRangedStr(bytes, 0, len(bytes) - 1)
	return string(bytes)
}

func clearSpace(bytes []byte) []byte{
	i := 0
	start := 0
	for j := 0; j < len(bytes); j += 1 {
		if bytes[j] == ' ' {
			if i == 0 {
				continue
			}
			if bytes[i - 1] == ' ' {
				continue
			}
			bytes[i] = ' '
			reverseRangedStr(bytes, start, i - 1)
			start = i + 1
		} else {
			bytes[i] = bytes[j]
		}
		i++
	} 
	if i <= len(bytes) && bytes[i - 1] == ' ' {
		return bytes[0 : i - 1]
	} else {
		reverseRangedStr(bytes, start, i - 1)
		return bytes[0 : i]
	}
}