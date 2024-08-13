package leetcode
// [239] 滑动窗口最大值

func maxSlidingWindow(nums []int, k int) []int {
	window := make([]int, 0)
	ans := make([]int, 0, len(nums) - k + 1)
	for i := 0; i < k; i++ {
		window = push(window, nums[i])
	}
	ans = append(ans, window[0])
	for i := k; i < len(nums); i++ {
		if window[0] == nums[i - k] {
			window = window[1:]
		}
		window = push(window, nums[i])
		ans = append(ans, window[0])
	}
	return ans
}

func push(nums []int, num int) []int {
	for len(nums) > 0 && nums[len(nums) - 1] < num {
		nums = nums[ : len(nums) - 1]
	}
	nums = append(nums, num)
	return nums
}