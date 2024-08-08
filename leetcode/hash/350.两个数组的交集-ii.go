package leetcode
// [350] 两个数组的交集 II

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	nums := make([]int, 0)
	for _, v := range nums2 {
		hash[v] ++
	}
	for _, v := range nums1 {
		if _, ok := hash[v]; ok {
			if hash[v] > 0 {
				hash[v]--
				nums = append(nums, v)
			}
		}
	}
	return nums
}
