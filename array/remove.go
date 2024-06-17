package array
// 27
func Remove(a []int, target int) []int{
	j := 0
	for i := 0; i < len(a); i++ {
		if a[i] == target {
			continue
		}
		a[j] = a[i]
		j++
	}
	return a
}