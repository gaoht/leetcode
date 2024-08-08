package leetcode

// [1] 两数之和


func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, v := range nums {
		hash[v] = i
	}
	for i, v := range nums {
		if _, ok := hash[target - v]; ok  && hash[target - v] != i{
			return []int{i, hash[target - v]}
		}
	}
	return []int{}
}

