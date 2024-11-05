package leetcode

import (
	"fmt"
	"strconv"
)
func calculate(s string) int {
    stack := make([]string, 0, len(s))
    expression := readExpression(s)
	fmt.Println(expression)
    for _, ex := range expression {
        if ex == "(" {
            stack = append(stack, ex)
            continue
        } else if ex == "+" || ex == "-" {
            stack = append(stack, ex)
            continue
        } else if ex == ")" {
            // 出栈
			result := stack[len(stack) - 1]
			stack = stack[0 : len(stack) - 2]
			stack = append(stack, result)
        } else {
			stack = append(stack, ex)
		}
		for len(stack) >= 3 && (stack[len(stack) - 2] == "+"  || stack[len(stack) - 2] == "-"){
			operator := stack[len(stack) - 2]
			right := stack[len(stack) - 1]
			left :=  stack[len(stack) - 3]
			result := calc(left, right, operator)
			stack = stack[0 : len(stack) - 3]
			stack = append(stack, result)
		}
    }
    if len(stack) > 0 {
        ans, _ := strconv.Atoi(stack[len(stack) - 1])
        return ans
    }
    return 0
}






func readExpression(s string)[]string {
    ans := make([]string, 0, len(s))
    pre := 0
    for i, b := range s {
        switch b {
            case '(', ')' :
                ans = appendDigits(pre, i, s, ans)
                ans = append(ans, string(b))
                pre = i + 1
            case '+':
                ans = appendDigits(pre, i, s, ans)
                ans = append(ans, string(b))
                pre = i + 1
            case '-':
                ans = appendDigits(pre, i, s, ans)
                // start with only '-', add '0'
                if len(ans) == 0 || ans[len(ans) - 1] == "(" {
                    ans = append(ans, "0")
                }
                ans = append(ans, "-")
                pre = i + 1
            case ' ':
                ans = appendDigits(pre, i, s, ans)
                pre = i + 1
            default:
                // 只有数字
				if i + 1 == len(s) {
					ans = appendDigits(pre, i + 1, s, ans)
				}
        }
    }
    return ans
}

func appendDigits(start int, end int, s string, ans []string) []string { 
    if start == end {
        return ans
    }
    return append(ans, s[start : end])
}


func calc(left, right string, operator string) string {
    l, _ := strconv.Atoi(left)
    r, _ := strconv.Atoi(right)
    if operator == "+" {
        return strconv.Itoa(l + r)
    } else {
        return strconv.Itoa(l - r)
    }
}