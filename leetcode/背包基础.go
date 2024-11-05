package leetcode


func bagSolve2D(capacity, items int, sizes []int, values []int) int {
	dp := make([][]int, items)
	for i := 0; i < items; i++ {
		dp[i] = make([]int, capacity + 1)
	}
	// 初始化第一个材料
	for j := 0; j <= capacity; j++ {
		if sizes[0] <= j {
			dp[0][j] = values[0]
		}
	}
	for i := 1; i < items; i++ {
		for j := capacity; j >=0; j-- {
			if sizes[i] > j {
				dp[i][j] = dp[i - 1][j]
				continue
			}
			dp[i][j] = maxV(values[i], values[i] + dp[i-1][j - sizes[i]])
		}
	}
	return dp[items - 1][capacity]
}
func bagSolve(capacity, items int, sizes []int, values []int) int {
	dp := make([]int, capacity + 1)
	for i := 0; i < items; i++ {
		for j := capacity; j >= sizes[i]; j-- {
			dp[j] = maxV(dp[j], values[i] + dp[j - sizes[i]])
		}
	}
	return dp[capacity]
}


func maxV(a, b int) int {
	if a > b {
		return a
	}
	return b
}
