package leetcode

import (
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
    s := strconv.Itoa(n) 
    ans := make([]byte, 0, len(s))
    jie := false
    for i := len(s) - 1; i >= 0; i-- {
        si := s[i]
        if jie {
            si = si - 1
            jie = false
        }
        if i == 0 {
            if  si > '0' {
                ans = append(ans, si)
            }
            continue
        }
        if si >= s[i - 1] {
            ans = append(ans, si)
            continue
        } else {
            ans = append(ans, '9')
            for j := len(ans) - 2; j >= 0; j-- {
                if ans[j] < '9' {
                    ans[j] = '9'
                } else {
                    break
                }
            }
            jie = true
        }
    }
    return reverse(ans)
}
func reverse(b []byte) int{
    left := 0
    right := len(b) - 1
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    } 
    i, _ := strconv.Atoi(string(b))
    return i
}