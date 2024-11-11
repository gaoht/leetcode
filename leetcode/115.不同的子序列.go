package leetcode

import "fmt"
func numDistinct(s string, t string) int {
    dp := make([][]int, len(t) + 1)
    for i := range dp {
        dp[i] = make([]int, len(s) + 1)
    }
    for j := range dp[0] {
        dp[0][j] = 1
    }
    for i := 1; i <= len(s); i++ {
        for j := 1; j <= len(t); j++ {
            if s[i - 1] == t[j - 1] {
                dp[j][i] = dp[j-1][i-1] + dp[j][i-1]
            } else {
                dp[j][i] = dp[j][i-1]
            }
			
        }
		fmt.Println()
		// fmt.Println(i, dp[i][j])
    }
    return dp[len(t)][len(s)]


}