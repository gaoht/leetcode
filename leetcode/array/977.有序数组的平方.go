package leetcode
// [977] 有序数组的平方
func sortedSquares(nums []int) []int {
	n := make([]int, len(nums))
	for i, l, r := len(nums) - 1, 0, len(nums) - 1; i >= 0; i--{
		if nums[l] * nums[l] >= nums[r] * nums[r] {
			n[i] = nums[l] * nums[l]
			l++
		} else {
			n[i] = nums[r] * nums[r]
			r--
		}
	}
	return n
}
// @lc code=end

