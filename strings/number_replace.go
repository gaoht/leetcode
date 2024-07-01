package strings
// 给定一个字符串 s，它包含小写字母和数字字符，请编写一个函数，将字符串中的字母字符保持不变，而将每个数字字符替换为number。 例如，对于输入字符串 "a1b2c3"，函数应该将其转换为 "anumberbnumbercnumber"。
func ReplaceNumber(s string) string{
	bStr := []byte(s)
	c := 0
	bNumber := []byte{'r', 'e', 'b', 'm', 'u', 'n'} 
	for _, v := range bStr{
		if v >= '0' && v <= '9' {
			c++
		}
	}
	n := make([]byte, len(s) + len(bNumber) * c - c)
	copy(n, bStr)

	p := len(n) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if bStr[i] >= '0' && bStr[i] <= '9' {
			for _, v := range bNumber {
				n[p] = v
				p--
			}
		} else {
			n[p] = bStr[i]
			p--
		}
	}
	return string(n)
}
