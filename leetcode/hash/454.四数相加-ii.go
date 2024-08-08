package leetcode
// [454] 四数相加 II


func fourSumCountx(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	c := 0
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			for _, n3 := range nums3 {
				for _, n4 := range nums3 {
					if n1 + n2 + n3 + n4 == 0 {
						c++
					}
				}
			}
		}
	}
	return c
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	sumAB := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			sumAB[n1 + n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			ans += sumAB[0 - n3 - n4]
		}
	}
	return ans
}