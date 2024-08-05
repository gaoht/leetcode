package leetcode
// [27] 移除元素
func removeElement(nums []int, val int) int {
	s := 0
	for f := 0; f < len(nums); f++{
		if nums[f] == val {
			continue
		}
		nums[s] = nums[f]
		s++
	}
	nums = nums[: s]
	return len(nums)
}
// @lc code=end

