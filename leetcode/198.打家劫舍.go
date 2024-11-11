package leetcode

import "fmt"
func rob(nums []int) int {
    dp := make([]int, len(nums))
    if len(nums) == 0 {
        return 0
    }
    dp[0] = nums[0]
    if len(nums) >= 2 {
        dp[1] = maxV(nums[0], nums[1])
    }
    for i := 2; i < len(nums); i++ {
		fmt.Println(i, dp, nums[i])
        dp[i] = maxV(dp[i - 1], dp[i - 2] + nums[i])
    }
    return dp[len(nums) - 1]
}

func rob2(nums []int) int {
    dp := make([][2]int, len(nums))
    if len(nums) == 0 {
        return nums[1]
    }
    dp[0] = [2]int{nums[0], 1}
    for i := 1; i < len(nums); i++ {
        dp1 := dp[i - 1]
        dp2 := [2]int{0, 0}
        if i >= 2 {
            dp2 = dp[i - 2]
        }
        if i != len(nums) - 1 {
            dp[i] = max2([2]int{dp2[0] + nums[i], dp2[1]}, dp1)
        } else {
            if dp2[1] > 0 {
                dp[i] = max2([2]int{nums[i], 0}, dp1)
            } else {
                dp[i] = max2([2]int{dp2[0] + nums[i], dp2[1]}, dp1)
            } 
        }
		fmt.Println(dp)
    }
    return dp[len(nums) - 1][0]
}

func max2(a, b [2]int) [2]int{
    if a[0] > b[0] {
        return a
    } else if a[0] == b[0]{
        if a[1] == 1 && b[1] == 1{
            return a
        } else {
            return [2]int{a[0], 0}
        }
    }
    return b
}