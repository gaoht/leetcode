package leetcode

import "fmt"
func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(word2) + 1)
        dp[i][0] = i
    }
    for i := 1; i <= len(word2); i++ {
        dp[0][i] = i
    }
    for i := 1; i <= len(word1); i++ {
        for j := 1; j <= len(word2); j++ {
            if word1[i - 1] == word2[j - 1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j] + 1, dp[i-1][j-1] + 2)
				dp[i][j] = min(dp[i][j], dp[i][j - 1] + 1)
            }
			fmt.Print(" ", dp[i][j])
        }
		fmt.Println()
    }
    return dp[len(word1)][len(word2)]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}