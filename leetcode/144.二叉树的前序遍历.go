package leetcode
// [144] 二叉树的前序遍历

func preorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return []int{}
	}
	nums = append(nums, root.Val)
	if root.Left != nil {
		nums = append(nums, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nums = append(nums, preorderTraversal(root.Right)...)
	}
	return nums
}

func preorderTraversal2(root *TreeNode) []int {
	nums := make([]int, 0)
	s := &Stack{}
	if root == nil {
		return []int{}
	}
	s.Push(root)
	for !s.IsEmpty() {
		node := s.Pop() 
		nums = append(nums, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return nums
}

type Stack struct{
	nodes []*TreeNode
}

func (s *Stack) Push(n *TreeNode){
	s.nodes = append(s.nodes, n)
}

func (s *Stack) Pop() *TreeNode{
	if s.IsEmpty() {
		return nil
	}
	n := s.nodes[len(s.nodes) - 1]
	s.nodes = s.nodes[0: len(s.nodes) - 1]
	return n
}
func (s *Stack) IsEmpty() bool{
	return len(s.nodes) == 0
}