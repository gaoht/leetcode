package leetcode

import "math"

// [654] 最大二叉树


func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	max, index := findMax(nums)
	root := &TreeNode{
		Val: max,
	}
	if index > 0 {
		root.Left = constructMaximumBinaryTree(nums[ : index])
	}
	if index + 1 < len(nums) {
		root.Right = constructMaximumBinaryTree(nums[index + 1 : ])
	}
	return root
}

func findMax(nums []int) (int, int){
	max := math.MinInt
	index := -1
	for i, n := range nums {
		if max < n {
			max = n
			index = i
		}
	}
	return max, index
}