package hash


// IsHappy 202
func IsHappy2(n int) bool {
    m := make(map[int]bool)
    for !m[n] && n != 1{
        m[n] = true
        n = happySum(n)
    }
    return n == 1
}
func happySum(n int) int{
    sum := 0
    for n > 0{
        sum += (n % 10) * (n % 10)
        n = n / 10
    }
    return sum
}

func IsHappy(n int) bool{
	fast := happySum(n)
	fast = happySum(fast)
	slow := happySum(n)
	for fast != slow {
		slow = happySum(slow)
		fast = happySum(fast)
		fast = happySum(fast)
	}
	return fast == 1
}