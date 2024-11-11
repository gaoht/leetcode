package leetcode
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := make([]int, 0, len(nums2))
    hash := map[int]int{}
    for _, n := range nums1{
        hash[n] = -1
    }
    stack = append(stack, 0)
    for i := 1; i < len(nums2); i++ {
        for len(stack) > 0 && nums2[stack[len(stack) - 1]] < nums2[i] {
            top := stack[len(stack) - 1]
            hash[nums2[top]] = nums2[i]
            stack = stack[0 : len(stack) - 1]
        }
        stack = append(stack, i)
    }
    for i, n := range nums1 {
        nums1[i] = hash[n]
    }
    return nums1
}