package leetcode
// [589] N 叉树的前序遍历


type Node struct {
	Val int
	Children []*Node
}

func preorderN(root *Node) []int {
	ans := []int{}
	s := &NStack{}
	if root == nil {
		return ans
	}
	s.Push(root)
	for !s.IsEmpty() {
		n := s.Pop()
		ans = append(ans, n.Val)
		for i := len(n.Children) - 1; i >= 0; i--{
			s.Push(n.Children[i])
		}
	}
	return ans
}


func preorderN2(root *Node) []int {
	ans := []int{}
    if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	for _, n := range root.Children {
		ans = append(ans, preorderN(n)...)
	}
	return ans
}


type NStack struct{
	nodes []*Node
}

func (s *NStack) Push(n *Node){
	s.nodes = append(s.nodes, n)
}

func (s *NStack) Pop() *Node{
	if s.IsEmpty() {
		return nil
	}
	n := s.nodes[len(s.nodes) - 1]
	s.nodes = s.nodes[0: len(s.nodes) - 1]
	return n
}
func (s *NStack) IsEmpty() bool{
	return len(s.nodes) == 0
}
