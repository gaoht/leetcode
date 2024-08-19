package leetcode
// [100] 相同的树

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(q.Right, p.Right)
	} else {
		return false
	}
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := &Queue{}
	queue.Push(p)
	queue.Push(q)
	for !queue.IsEmpty() {
		l := queue.Pop()
		r := queue.Pop()
		if l == nil && r == nil || l != nil && r != nil && l.Val == r.Val {
			if l != nil {
				queue.Push(l.Left)
				queue.Push(r.Left)

				queue.Push(l.Right)
				queue.Push(r.Right)
			}
		} else {
			return false
		}
	}
	return true
	
}

