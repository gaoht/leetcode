package leetcode
//[35] 搜索插入位置
// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := -1, len(nums)
	for left + 1 != right {
		mid := left + (right - left) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
// @lc code=end