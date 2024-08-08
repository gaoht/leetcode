package leetcode
// [202] 快乐数
func isHappy(n int) bool {
	hash := make(map[int]struct{})
	for {
		sum := getHappy(n)
		if sum == 1 {
			return true
		} else {
			if _, ok := hash[sum]; ok {
				return false
			}
			hash[sum] = struct{}{}
			n = sum
		}
	}
}

func getHappy(n int) int {
	sum := (n % 10) * (n % 10)
	for n >= 10 {
		n = n / 10
		sum +=  (n % 10) * (n % 10) 
	}
	return sum
}
