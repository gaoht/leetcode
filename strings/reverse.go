package strings
// ReverseString 344
func ReverseString(s []byte)  {
	left := 0
    right := len(s) - 1
    for left < right {
        t := s[left]
        s[left] = s[right]
        s[right] = t
        left++
        right--
    }
}