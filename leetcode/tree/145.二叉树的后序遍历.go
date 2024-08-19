package leetcode
// [145] 二叉树的后序遍历

func postorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return []int{}
	}
	if root.Left != nil {
		nums = append(nums, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nums = append(nums, postorderTraversal(root.Right)...)
	}
	nums = append(nums, root.Val)
	return nums
}

func postorderTraversal2(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return []int{}
	}
	s := &Stack{}
	s.Push(root)
	for !s.IsEmpty() {
		n := s.Pop()
		nums = append(nums, n.Val)
		if n.Left != nil {
			s.Push(n.Left)
		} 
		if n.Right != nil {
			s.Push(n.Right)
		}
	}
	reverseNums(nums)
	return nums
}

func reverseNums(nums []int){
	for left, right := 0, len(nums) - 1; left < right; left, right = left + 1, right - 1{
		nums[left], nums[right] = nums[right], nums[left]
	}
}