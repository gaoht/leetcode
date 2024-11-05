package leetcode

import (
	"fmt"
)
func permuteUnique(nums []int) [][]int {
    ans := [][]int{}
    var dfs func(path []int,  traced map[int]bool)
    l := len(nums)
    dfs = func(path []int,  traced map[int]bool) {
        if len(path) == l {
            tmp := make([]int, l)
            copy(tmp, path)
            ans = append(ans, tmp)
            return
        }
        used := make(map[int]int)
		
        for i := 0; i < len(nums); i++ {
			fmt.Printf("%v \n", used)
            if used[nums[i]] > 0{
                continue
            } else {
                used[nums[i]] = 1  
            }
            if traced[i] {
                continue
            } else {
                traced[i] = true
            }
          
            path = append(path, nums[i])
            dfs(path, traced)
            path = path[0 : len(path) - 1]
            traced[i] = false
        }
    }
    dfs([]int{}, make(map[int]bool))
    return ans
}