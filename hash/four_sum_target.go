package hash

import "sort"

// FourSum 18
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
    a := make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i]  > 0 && nums[i] > target {
            break
        }
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        for j := i+1; j < len(nums); j++ {
            if nums[i] + nums[j] > 0 && nums[i] + nums[j] > target {
                break
            }
            // 去重
            if j > i + 1 && nums[j] == nums[j - 1] {
                continue
            }
            left := j + 1
            right := len(nums) - 1 
            for left < right {
                if nums[i] + nums[j] + nums[left] + nums[right] > target {
                    right --
                } else if nums[i] + nums[j] + nums[left] + nums[right] < target {
                    left ++
                } else {
                    a = append(a, []int{nums[i], nums[j], nums[left], nums[right]})
                    for left < right && nums[left + 1] == nums[left] {
                        left++
                    } 
                    for left < right && nums[right - 1] == nums[right] {
                        right--
                    }
                    left++
                    right--
                }
            }
        }
    }
	return a
}