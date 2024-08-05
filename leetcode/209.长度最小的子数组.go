package leetcode

import "math"

// [209] 长度最小的子数组

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// target = 7, nums = [2,3,1,2,4,3]  2
	// target = 4, nums = [1,4,4] 1
	// target = 11, nums = [1,1,1,1,1,1,1,1]
	if len(nums) == 0 {
		return 0
	}
	length, i, j, sum := math.MaxInt, 0, 0, 0
	for j < len(nums) {
		if sum + nums[j] >= target {
			if j - i + 1 < length {
				length = j - i + 1
			}
			sum -= nums[i]
			i++
		} else {
			sum += nums[j]
			j++
		}
	}
	if length ==  math.MaxInt{
		return 0
	}
	return length
}
// @lc code=end 

