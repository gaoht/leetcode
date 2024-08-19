package leetcode
// [107] 二叉树的层序遍历 II

func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if  root == nil {
		return ans
	}
	q := &Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		l := q.Len()
		levelAns := []int{}
		for i := 0; i < l; i++ {
			n := q.Pop()
			levelAns = append(levelAns, n.Val)
			if n.Left != nil {
				q.Push(n.Left)
			}
			if n.Right != nil {
				q.Push(n.Right)
			}
		}
		ans = append([][]int{levelAns}, ans...)
	}
	return ans
}

