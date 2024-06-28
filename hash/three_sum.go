package hash

import "sort"

// ThreeSum 15
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	a := make([][]int, 0)
	for i := 0; i < len(nums); i++{
		if nums[i] > 0 {
			break
		}
		// 正确去重a方法
		if i > 0 && nums[i] == nums[i - 1] {
			continue;
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i] + nums[left] + nums[right] > 0 {
				right --
			} else if nums[i] + nums[left] + nums[right] < 0 {
				left ++
			} else {
				a = append(a, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right - 1] {
					right --
				}
				for left < right && nums[left] == nums[left + 1] {
					left ++
				}
				right --
				left ++
			}
		}
	}
	return a
}