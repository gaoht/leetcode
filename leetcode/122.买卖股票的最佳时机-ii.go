package leetcode

import "fmt"

func maxProfitGreedy(prices []int) int {
    profits := 0
    for i := 1; i <= len(prices) - 1; i++ {
        if prices[i] - prices[i - 1] > 0 {
            profits += prices[i] - prices[i - 1]
        }
    }
    return profits
}

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    dp := make([][2]int, len(prices))
    // 持有股票
    dp[0][0] = -prices[0]
    // 不持有股票
    dp[0][1] = 0

    for i := 1; i < len(prices); i++ {
        dp[i][0] = maxVal(dp[i - 1][0], dp[i - 1][1] - prices[i])
        dp[i][1] = maxVal(dp[i - 1][0] + prices[i], dp[i - 1][1])
		fmt.Println(dp[i])
    }
    return dp[len(prices) - 1][1]
}

func maxVal(a, b int) int{
    if a > b {
        return a
    }
    return b
}