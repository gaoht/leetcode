package leetcode

import "fmt"
func findMin(nums []int) int {
    left := -1
    right := len(nums) * 2
    for left + 1 != right {
        mid := left + (right - left) / 2
        m, l, r := mid % len(nums), (mid - 1 + len(nums)) % len(nums), (mid + 1) % len(nums)
        if nums[m] < nums[r] && nums[m] < nums[l] {
            return nums[mid % len(nums)]
        }
        if nums[m] <= nums[r] {
            left = mid
        } else {
            right = mid
        }
    }
	fmt.Printf("left=%d, right=%d \n", left, right)
    return 0
}