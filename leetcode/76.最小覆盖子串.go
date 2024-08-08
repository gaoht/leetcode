package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
 
func minWindow(s string, t string) string {
	slow, fast := 0, 0
	max := math.MaxInt
	minString := ""
	tHash := buildCharHash(t)
	for fast < len(s) {
		if _, ok := tHash[s[fast]]; ok {
			tHash[s[fast]]--
			if isT(tHash){
				length := fast - slow + 1
				if length < max {
					max  = length
					minString = s[slow : fast+1]
				}
				// 移动窗口到下一个在窗口内的字符位置
				i := slow + 1
				if _, ok := tHash[s[slow]]; ok {
					tHash[s[slow]]++
				}
				for i <= fast {
					if _, ok := tHash[s[i]]; ok {
						break
					}
					i++
				}
				slow = i
				// 下次循环还用fast，还原fast的值
				tHash[s[fast]]++
			} else {
				fast++
			}
		} else {
			fast++
		}
	}
	return minString
}
// @lc code=end

func buildCharHash(t string) map[byte]int{
	m := make(map[byte]int)
	for _, c := range []byte(t) {
		m[c]++
	}
	return m
}


// 是否整个包含t
func isT(h map[byte]int) bool{
	for _, v := range h {
		if  v  > 0 {
			return false
		}
	}
	return true
}