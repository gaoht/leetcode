package leetcode

import "sort"

// [18] 四数之和

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	if len(nums) < 4 {
		return ans
	}
	for i := 0; i <= len(nums) - 4; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
        if nums[i] >=0 && nums[i] > target {
            break
        }
		for j :=  i + 1; j <= len(nums) -3; j++ {
			if j - i != 1 && nums[j] == nums[j-1] {
				continue
			}
            if nums[i] + nums[j] >= 0 && nums[i] + nums[j] > target {
                break
            }
			left := j + 1;
			right := len(nums) - 1
			for left < right {
                if left - 1 != j && nums[left] == nums[left - 1] {
					left ++
                    continue
                }
				if nums[i] + nums[j] + nums[left] + nums[right] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left ++
					right --
				} else if nums[i] + nums[j] + nums[left] + nums[right] < target {
					left ++
				} else {
					right --
				}
			}
		}
	}
	return ans
}
