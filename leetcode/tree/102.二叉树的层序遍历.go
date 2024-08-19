package leetcode
// [102] 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
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
		ans = append(ans, levelAns)
	}
	return ans
}


func levelOrder2(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	var traversal func(n *TreeNode, level int)
	traversal = func(n *TreeNode, level int){
		if level == len(ans) {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], n.Val)
		if n.Left != nil {
			traversal(n.Left, level+1)
		}
		if n.Right != nil {
			traversal(n.Right, level+1)
		}
	}
	traversal(root, 0)
	return ans
}


type Queue struct {
	nodes []*TreeNode
}

func (q *Queue) Push(n *TreeNode){
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Pop() *TreeNode {
	n  := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *Queue) IsEmpty () bool{
	return len(q.nodes) == 0
}

func (q *Queue) Len () int{
	return len(q.nodes) 
}
