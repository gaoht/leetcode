/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
package leetcode
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