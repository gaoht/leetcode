package leetcode
// [704] 二分查找
// @lc code=start
func search(nums []int, target int) int {
	left, right := -1, len(nums)
	for left + 1 != right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return -1
}
// @lc code=end
// http://www.programmercarl.com
