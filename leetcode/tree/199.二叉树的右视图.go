package leetcode


func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := &Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		l := q.Len()
		var n *TreeNode
		for i := 0; i < l; i++ {
			n = q.Pop()
			if n.Right != nil {
				q.Push(n.Right)
			}
			if n.Right == nil && n.Left != nil {
				q.Push(n.Left)
			}
		}
		ans = append(ans, n.Val)
	}
	return ans
}


