package array

// 977
func Square(a []int) []int{
	// []{-4, 0, 1, 2, 3}
	left := 0
	right := len(a) - 1
	ret := make([]int, len(a))
	for i := len(a) - 1; i >=0; i-- {
		if a[left] * a[left] < a[right] * a[right] {
			ret[i] = a[right] * a[right]
			right--
		} else {
			ret[i] = a[left] * a[left]
			left++
		}
	}
	return ret
}21345