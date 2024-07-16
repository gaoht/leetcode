package stack

import (
	"strconv"
)

// EvalRPN 150
func EvalRPN(tokens []string) int {
	s := make([]string, 0)
	for _, t := range tokens {
		left := 0
		right := 0
		v  := 0
		switch t {
		case "+",  "-",  "*", "/":
			l := s[len(s) - 1]
			s = s[: len(s) - 1]
			right, _ = strconv.Atoi(l)
			l = s[len(s) - 1]
			s = s[: len(s) - 1]
			left, _ = strconv.Atoi(l)
		}
		switch t {
		case "+":
			v = left + right
			s = append(s, strconv.Itoa(v))
		case "-":
			v = left - right
			s = append(s, strconv.Itoa(v))
		case "*":
			v = left *  right
			s = append(s, strconv.Itoa(v))
		case "/":
			v = left / right	
			s = append(s, strconv.Itoa(v))
		default:
			s = append(s, t)
		}
	}
	l, _ := strconv.Atoi(s[0])
	return l
}

