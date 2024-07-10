package strings

// Strstr 28
func Strstr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i :=0;  i + len(needle) <= len(haystack); i++ {
		if haystack[i : i+len(needle)] == needle {
			return i
		}
	}
	return -1
}