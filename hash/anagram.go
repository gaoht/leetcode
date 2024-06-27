package hash

func IsAnagram(s string, t string) bool{
	dict := [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		dict[s[i] - 'a']++
		dict[t[i] - 'a'] --
	}
	for _, v :=  range dict{
		if v != 0 {
			return false
		}
	}
	return true
}