package leetcode

// [590] N 叉树的后序遍历
func postorderN(root *Node) []int {
    if root == nil {
		return []int{}
	}
	ans := []int{}
	s := &NStack{}
	s.Push(root)
	for !s.IsEmpty() {
		n := s.Pop()
		ans = append(ans, n.Val)
		for _, c := range n.Children {
			s.Push(c)
		}
	}
	reverseNums(ans)
	return ans
}


