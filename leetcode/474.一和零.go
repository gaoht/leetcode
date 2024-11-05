package leetcode
func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for i := 0; i < m + 1; i++ {
        dp[i] = make([]int, n+1)
    }
    str01s := count1_0(strs) 
    for i := 0; i < len(strs); i++ {
        // 0
        for j := m; j >= str01s[i][0]; j-- {
            for k := n; k >= str01s[i][1]; k-- {
                dp[j][k] = maxL(dp[j][k], dp[j-str01s[i][0]][k-str01s[i][1]] + 1)
            }
        }
    }
    return dp[m][n]
}

func count1_0(strs []string) [][2]int {
    ans := make([][2]int, len(strs))
    for i, str := range strs {
        ans[i] = [2]int{0, 0}
        for ii := range str {
            if str[ii] == '0' {
                ans[i][0]++
            } else {
                ans[i][1]++
            }
        }
    }
    return ans
}

func maxL(a, b int) int {
    if a > b {
        return a
    }
    return b
}