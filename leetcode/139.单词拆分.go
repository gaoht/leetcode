package leetcode


func wordBreak(s string, wordDict []string) bool {
    sLen := len(s)
    dp := make([]int, sLen + 1)
    dp[0] = 1
    dict := map[string]int{}
    for _, word := range wordDict {
        dict[word] = len(word)
    }
    for j := 0; j <= sLen; j++ {
		for i := 0; i < len(wordDict); i++ {
			if j >= len(wordDict[i]) && 
				s[j-len(wordDict[i]) : j] == wordDict[i] && 
				dp[j-len(wordDict[i])] > 0 {
					dp[j] = dp[j] + dp[j - len(wordDict[i])]
			}
        }
    }
    return dp[sLen] > 0
}

func wordBreakError(s string, wordDict []string) bool {
    sLen := len(s)
    dp := make([]int, sLen + 1)
    dp[0] = 1
    dict := map[string]int{}
    for _, word := range wordDict {
        dict[word] = len(word)
    }
    for i := 0; i < len(wordDict); i++ {
        for j := len(wordDict[i]); j <= sLen; j++ {
            if s[j - len(wordDict[i]) : j] == wordDict[i] && dp[j - len(wordDict[i])] > 0 {
				dp[j] = dp[j] + dp[j - len(wordDict[i])]
			}
        }
    }
    return dp[sLen] > 0
}