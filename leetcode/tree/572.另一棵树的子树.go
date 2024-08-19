package leetcode

// [572] 另一棵树的子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}
	q := &Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		n := q.Pop()
		if checkIsSameTree(n, subRoot) {
			return true
		} 
		if n.Left != nil {
			q.Push(n.Left)	
		}
		if n.Right != nil {
			q.Push(n.Right)
		}
	}
	return false
}

func isSubtreeR(r *TreeNode, s *TreeNode) bool {
	if r == nil {
		return false
	}
	if isSameTreeR(r, s) {
		return true
	} else {
		if r.Left != nil {
			 if (isSubtreeR(r.Left, s)){
				return true
			 }
		}
		if r.Right != nil {
			if (isSubtreeR(r.Right, s)){
				return true
			 }
		}
	}
	return false
}
func checkIsSameTree(l *TreeNode, r *TreeNode) bool {
	q := &Queue{}	
	q.Push(l)
	q.Push(r)

	for !q.IsEmpty() {
		left := q.Pop()
		right := q.Pop()
		if left == nil && right == nil{
			continue
		} else if left != nil && right != nil && left.Val == right.Val{
			q.Push(left.Left)
			q.Push(right.Left)
			q.Push(left.Right)
			q.Push(right.Right)
		} else {
			return false
		}
	}
	return true
}

func isSameTreeR(l *TreeNode, r *TreeNode) bool { 
	if l == nil && r == nil {
		return true
	}
	if l != nil && r != nil && r.Val == l.Val {
		return isSameTreeR(l.Left, r.Left) && isSameTreeR(l.Right, r.Right)
	}
	return false
}