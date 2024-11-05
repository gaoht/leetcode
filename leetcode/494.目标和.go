package leetcode

import "math"

// import "fmt"
func findTargetSumWaysDfs(nums []int, target int) int {
    sum := 0
	// nums2 := make([]int, 0, len(nums))
    for _, n := range nums {
        sum += n
    }   
    if target > sum {
        return 0
    } 
    if (sum + target) % 2 == 1 {
        return 0
    }
    ans := 0
    left := (sum + target) / 2
    var dfs func(path []int, start, sum int)
    dfs = func(path []int, start, sum int) {
        if sum == left {
            ans++
        }
        if sum > left {
            return
        }
        for i := start; i < len(nums); i++ {
            sum += nums[i]
			path = append(path, i)
            dfs(path, i + 1, sum)
            sum -= nums[i]
			path = path[0 : len(path) - 1]
        }
    }
    dfs([]int{}, 0, 0)
    return ans
}


func findTargetSumWays2D(nums []int, target int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    left := (sum + target) / 2
    if sum < int(math.Abs(float64(target))) {
        return 0
    }
    if (sum + target) % 2 == 1 {
        return 0
    }
    dp := make([][]int, len(nums))
    zeros := 0
    for i, n := range nums {
        dp[i] = make([]int, left + 1)
        // 初始化第0列
        if n == 0 {
            zeros++
        }
        dp[i][0] = int(math.Pow(float64(2), float64(zeros)))
    } 

    // 初始化第1行
    if nums[0] <= left {
        dp[0][nums[0]] = 1
    }
    for i := 1; i < len(nums); i++ {
        for j := 1; j <= left; j++ {
            if nums[i] > j {
                dp[i][j] = dp[i - 1][j]
            } else {
                dp[i][j] = dp[i - 1][j] + dp[i - 1][j - nums[i]]
            }
        }
    }
    return dp[len(nums) - 1][left]   
}




func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    left := (sum + target) / 2
    if sum < int(math.Abs(float64(target))) {
        return 0
    }
    if (sum + target) % 2 == 1 {
        return 0
    }
    dp := make([]int, left + 1)
    // 初始化第1行
    dp[0] = 1
    for i := 0; i < len(nums); i++ {
        for j := nums[i]; j <= left; j++ {
            dp[j] = dp[j] + dp[j - nums[i]]
        }
    }
    return dp[left]   
}

