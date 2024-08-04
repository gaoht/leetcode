/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package main
// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := getLeft(nums, target)
	r := getLeft(nums, target + 1) - 1
	if l >= len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	return []int{l, r}
}

func getLeft(nums []int, target int) int {
	left, right := -1, len(nums)
	for left + 1 != right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}


// @lc code=end

