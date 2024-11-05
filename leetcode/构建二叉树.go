package leetcode

import "math"

func buildBinaryTreeFromArray(nums []*int) *TreeNode{
	q := []*TreeNode{nil}
	i := 0
	depth := 0
	var root *TreeNode
	for i < len(nums) {
		nodes := int(math.Pow(2, float64(depth)))
		for j := 0; j < nodes && i < len(nums); i, j = i + 1, j + 1 {
			parent := q[0]
			q = q[1:]
			var node *TreeNode
			if nums[i] != nil {
				node = &TreeNode{
					Val: *nums[i],
				}
			}
			if i == 0 {
				root = node
			}
			if parent != nil {
				if j % 2 == 0 {
					parent.Left = node
				} else {
					parent.Right = node
				}
			}
			q = append(q, node)
			q = append(q, node)
		}
		depth++
	}
	return root
}