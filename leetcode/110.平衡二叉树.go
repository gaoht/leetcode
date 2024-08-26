package leetcode

// [110] 平衡二叉树

func isBalanced(root *TreeNode) bool {
	_, b := getDepthR(root)
	return b
}


func getDepthR(n *TreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}
	if n.Left == nil && n.Right == nil {
		return 1, true
	}
	l, lb := 0, true
	if n.Left != nil {
		l, lb = getDepthR(n.Left)
	} 
	r, rb := 0, true
	if n.Right != nil {
		r, rb = getDepthR(n.Right)
	}
	if l > r {
		return 1 + l, (lb && rb) && (l - r == 1)
	} else {
		return 1 + r, (lb && rb) && (l - r == -1 ||  l == r)
	}
}

func isBalanced2(root *TreeNode) bool{
	if root == nil {
		return true
	}
	q := &Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		l := q.Len()
		for i := 0; i < l; i++ {
			n := q.Pop()
			ll, rr := 0, 0 
			if n.Left != nil {
				ll = getDepth(n.Left)
				q.Push(n.Left)
			}
			if n.Right != nil {
				rr = getDepth(n.Right)
				q.Push(n.Right)
			}
			if ll - rr < -1 || ll - rr > 1 {
				return false
			}
		}
	}
	return true
}
func getDepth(n *TreeNode) int {
	d := 0
	if n == nil {
		return d
	}
	q := &Queue{}
	q.Push(n)
	for !q.IsEmpty() {
		l := q.Len()
		for i := 0; i < l; i++ {
			n := q.Pop()
			if n.Left != nil {
				q.Push(n.Left)
			}
			if n.Right != nil {
				q.Push(n.Right)
			}
		}
		d++
	}
	return d
}