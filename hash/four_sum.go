package hash

// FourSumCount 454
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	a := make(map[int]int) 
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2  {
			a[0 - v1 - v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += a[v3 + v4]
		}
	}
	return count
}