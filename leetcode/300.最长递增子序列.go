package leetcode

import (
	"fmt"
	"math"
)

func lengthOfLIS(nums []int) int {
    max := 0
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        if max < len(path) {
			fmt.Println(path)
            max = len(path)
        }
        for i := start; i < len(nums); i++ {
            end := math.MinInt
            if len(path) > 0 {
                end = path[len(path) - 1]
            }
            if end < nums[i] {
                path = append(path, nums[i])
                dfs(i + 1, path)
                path = path[0 : len(path) - 1]
            } 
        }
    }
    dfs(0, []int{})
    return max
}