package leetcode
func inorderTraversal(root *TreeNode) []int {
	s := &Stack{}
	cur := root
	ans := []int{}
	for cur != nil || !s.IsEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			n := s.Pop().(*TreeNode)
			ans = append(ans, n.Val)
			if n.Right != nil {
				cur = n.Right
			}
		}
	}
	return ans
}

