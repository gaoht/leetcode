package leetcode

import "fmt"


// [55] 跳跃游戏
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	var dfs func(path []int) bool
	target := len(nums) - 1
	dfs = func(path []int) bool {
        if target == 0 && len(path) == 0 {
            return true
        }
		if len(path) > 0 {
			if path[len(path) - 1] == target {
				fmt.Println(path)
				return true
			}
			if path[len(path) - 1] > target {
				return false
			}
		}
		last := 0
		if len(path) > 0 {
			last = path[len(path) - 1]
		}
		for i := 1; i <= nums[last]; i++ {
			if i + last <= target {
				path = append(path, i + last)
				if dfs(path) {
					return true
				}
				path = path[0 : len(path) - 1]
			}
		}
		return false
	}
	return dfs(make([]int, 0, len(nums)))
}

