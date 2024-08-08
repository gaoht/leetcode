package leetcode

import "sort"

// [15] 三数之和


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums) - 2; i++{
		// 重复数字排除
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left ++
			} else {
				// 重复数字排除
				ans = append(ans, []int{nums[i], nums[left], nums[right]})	
				left ++
				for left <= right && nums[left] == nums[left - 1] {
					left++
				}
			}
		}
	}
	return ans 
}

