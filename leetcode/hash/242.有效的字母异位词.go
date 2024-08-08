package leetcode
//  [242] 有效的字母异位词


// @lc code=start
func isAnagram(s string, t string) bool {
	b := [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		b[s[i] - 'a'] ++
		b[t[i] - 'a'] --
	}
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

