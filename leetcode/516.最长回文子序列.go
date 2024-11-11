package leetcode

import "fmt"
func longestPalindromeSubseq(s string) int {
    dp := make([][]int, len(s))
    for i := range dp {
        dp[i] = make([]int, len(s))
        dp[i][i] = 1
    }
    for j := 1; j < len(s); j++ {
        for i := 0; i < j; i++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = maxV(dp[i+1][j], dp[i][j-1])
            }
        }  
    }
	fmt.Println(dp)
    return dp[0][len(s)-1]
}