package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInOrderTraversal(t *testing.T){
	Convey("中序遍历", t, func(){
		node := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		So(inorderTraversal(node), ShouldResemble, []int{1,3,2})
		So(inorderTraversal2(node), ShouldResemble, []int{1,3,2})
		node = nil
		So(inorderTraversal(node), ShouldResemble, []int{})
		So(inorderTraversal2(node), ShouldResemble, []int{})
		node = &TreeNode{
			Val: 1,
		}
		So(inorderTraversal(node), ShouldResemble, []int{1})
		So(inorderTraversal2(node), ShouldResemble, []int{1})
	})
}
func TestPreOrderTraversal(t *testing.T){
	Convey("前序遍历", t, func(){
		node := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		So(preorderTraversal(node), ShouldResemble, []int{1,2,3})
		So(preorderTraversal2(node), ShouldResemble, []int{1,2,3})
		node = nil
		So(preorderTraversal(node), ShouldResemble, []int{})
		So(preorderTraversal2(node), ShouldResemble, []int{})
		node = &TreeNode{
			Val: 1,
		}
		So(preorderTraversal(node), ShouldResemble, []int{1})
		So(preorderTraversal2(node), ShouldResemble, []int{1})
	})
}

func TestPostOrderTraversal(t *testing.T){
	Convey("后序遍历", t, func(){
		node := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		So(postorderTraversal(node), ShouldResemble, []int{3,2,1})
		So(postorderTraversal2(node), ShouldResemble, []int{3,2,1})
		node = nil
		So(postorderTraversal(node), ShouldResemble, []int{})
		So(postorderTraversal2(node), ShouldResemble, []int{})
		node = &TreeNode{
			Val: 1,
		}
		So(postorderTraversal(node), ShouldResemble, []int{1})
		So(postorderTraversal2(node), ShouldResemble, []int{1})
	})
}