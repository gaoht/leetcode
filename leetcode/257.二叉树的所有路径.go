package leetcode

import (
	"strconv"
	"strings"
)

type Ans struct {
	nums []string
}

// [257] 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	ans := &Ans{nums: []string{}}
	getTreePathR(root, []string{}, ans)
	return ans.nums
}


func getTreePathR(n *TreeNode, nums []string, ans *Ans){
	if n == nil {
		return
	}
	if n.Left == nil && n.Right == nil {
		nums = append(nums, strconv.Itoa(n.Val))
		ans.nums = append(ans.nums, strings.Join(nums, "->"))
	}
	nums = append(nums, strconv.Itoa(n.Val))
	if n.Left != nil {
		getTreePathR(n.Left, nums, ans)
	}
	if n.Right != nil {
		getTreePathR(n.Right, nums, ans)
	}
}


// [257] 二叉树的所有路径
func binaryTreePaths2(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	s := &Stack{}
	p := &Stack{}
	ans := []string{}
	s.Push(root)
	p.Push([]string{strconv.Itoa(root.Val)})
	for !s.IsEmpty() {
		n := s.Pop().(*TreeNode)
		path := p.Pop().([]string)
		if n.Left == nil && n.Right == nil {
			ans = append(ans, strings.Join(path, "->"))
		}
		
		if n.Right != nil {
			t := []string{}
			t  = append(t, path...)
			s.Push(n.Right)
			t = append(t, strconv.Itoa(n.Right.Val))
			p.Push(t)

		}
		if n.Left != nil {
			t := []string{}
		t  = append(t, path...)
			s.Push(n.Left)
			t = append(t, strconv.Itoa(n.Left.Val))
			p.Push(t)
		}
	}
	return ans
}


type Stack struct {
	values []interface{}
}

func (s *Stack) Size() int {
	return len(s.values)
}



func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Push(o interface{}) {
	s.values = append(s.values, o)
}

func (s *Stack) Pop() interface{} {
	o := s.values[len(s.values) - 1]
	s.values = s.values[:len(s.values) - 1]
	return o
}