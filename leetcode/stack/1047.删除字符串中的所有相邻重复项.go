package leetcode
// [1047] 删除字符串中的所有相邻重复项

// @lc code=start
func removeDuplicates(s string) string {
	stack := make([]rune, 0)
	for _, r := range s {
		if len(stack) > 0 && stack[len(stack) - 1] == r {
			stack = stack[: len(stack) - 1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
// @lc code=end

