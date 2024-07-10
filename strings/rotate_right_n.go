package strings
// RotateRightN 字符串的右旋转操作是把字符串尾部的若干个字符转移到字符串的前面。给定一个字符串 s 和一个正整数 k，请编写一个函数，将字符串中的后面 k 个字符移到字符串的前面，实现字符串的右旋转操作。 
// 例如，对于输入字符串 "abcdefg" 和整数 2，函数应该将其转换为 "fgabcde"。
func RotateRightN(s string, n int) string{
	b := []byte(s)
	reverseBytes(b, 0, len(b) - 1)
	reverseBytes(b, 0, n - 1)
	reverseBytes(b, n, len(b) - 1)
	return string(b)
}