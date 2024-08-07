package leetcode

// [80] 删除有序数组中的重复项 II

func removeDuplicates1(nums []int) int {
	l := len(nums)
	cursor := 1
	index := 1
	letter := 0
	count := 1

	for cursor < len(nums) {
		if nums[letter] == nums[cursor] {
			count++
			if count <= 2 {
				nums[index] = nums[cursor]
				index++
			} else {
				l--
			}
		} else {
			count = 1
			nums[index] = nums[cursor]
			letter = index
			index++
		}
		cursor++
	}
	return l
}


func removeDuplicates(nums []int) int {
	slow, fast := 2, 2
	if len(nums) <= 2 {
		return len(nums)
	}
	for fast < len(nums) {
		if nums[fast] != nums[slow - 2] {
			nums[slow] = nums[fast];
			slow++
		} 
		fast ++
	}
	return slow
}