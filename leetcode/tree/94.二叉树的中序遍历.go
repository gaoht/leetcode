package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return []int{}
	}
	if root.Left != nil {
		nums = append(nums, inorderTraversal(root.Left)...)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		nums = append(nums, inorderTraversal(root.Right)...)
	}
	return nums
}

func inorderTraversal2(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return []int{}
	}
	s := &Stack{}
	cur := root
	for cur != nil || !s.IsEmpty(){
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			node := s.Pop()
			nums = append(nums, node.Val)	
			cur = node.Right
		}
	}
	return nums
}

