
package leetcode
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Children []*Node
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
