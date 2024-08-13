package leetcode


// [20] 有效的括号
func isValid(s string) bool {
	left := map[rune]struct{} {
		'(': {},
		'[': {},
		'{': {},
	}
	right := map[rune]rune {
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := make([]rune, 0)
	for _, b := range s {
		if _, ok := left[b]; ok {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack) - 1 : ][0] != right[b] {
				return false
			} else {
				stack = stack[0 : len(stack) - 1]
			}
		}
	}
	return len(stack) == 0
}

