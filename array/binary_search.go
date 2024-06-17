package array
// 704
func BinarySearch(a []int, target int) int {
	left := 0 
	right := len(a) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if a[mid] > target {
			right = mid - 1
		} else if a[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
