package strings
// ReverseStr 541
func ReverseStr(s string, k int) string {
	b := []byte(s)
    for i := 0; i < len(b); i += 2 * k {
        if i + k <= len(b) {
            reverseString(b, i, i + k - 1)
        } else {
            reverseString(b, i, len(b) - 1)
        } 
    }
    return string(b)
}
func reverseString(s []byte, left int, right int) {
    for left < right {
        t := s[left]
        s[left] = s[right]
        s[right] = t
        left++
        right--
    }
}