package leetcode

import "fmt"
func change2D(amount int, coins []int) int {
    dp := make([][]int, len(coins))
    for i := 0; i < len(coins); i++ {
        dp[i] = make([]int, amount + 1)
        dp[i][0] = 1
    } 
    
    for i := coins[0]; i <= amount; i++ {
        if i % coins[0] == 0 {
            dp[0][i] = 1
        }
    }
    for i := 1; i < len(coins); i++ {
        for j := 1; j <= amount; j++ {
            if j < coins[i] {
                dp[i][j] = dp[i - 1][j]
            } else {
                dp[i][j] = dp[i][j - coins[i]] + dp[i-1][j]
            }
        }
		fmt.Println(dp[i])
    }
    return dp[len(coins) - 1][amount]
}


func change(amount int, coins []int) int {
    dp := make([]int, amount + 1)
    dp[0] = 1
    for i := 0; i < len(coins); i++ {
        for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j - coins[i]];
        }
		fmt.Println(dp)
    }
    return dp[amount]
}