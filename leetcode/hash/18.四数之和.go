package leetcode

import "sort"

// [18] 四数之和

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums) - 3; i++{
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		if nums[i] > 0 && nums[i] > target {
			break
		}
		for j := i + 1; j < len(nums) - 2; j++ {
			if nums[i] + nums[j] > 0 && nums[i] + nums[j] > target {
				break		
			}
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right --
				} else if sum < target {
					left ++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left ++
					for left <= right && nums[left] == nums[left - 1]{
						left ++
					}
				}
			}
		}
	}
	return ans
}

