package leetcode
// [101] 对称二叉树

func isSymmetric2(root *TreeNode) bool {
	return check(root, root)
}


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := &Queue{}
	q.Push(root)
	q.Push(root)
	for !q.IsEmpty() {
		left := q.Pop()
		right := q.Pop()
		if left == right || left !=nil && right != nil && left.Val == right.Val {
			if left != nil {
				q.Push(left.Left)
				q.Push(right.Right)
				q.Push(left.Right)
				q.Push(right.Left)
			}
		} else {
			return false
		}
	}
	return true
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}