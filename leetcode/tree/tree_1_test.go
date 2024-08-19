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


func TestLevelOrder(t *testing.T){
	Convey("二叉树层序遍历", t, func(){
		n := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		So(levelOrder(n), ShouldResemble, [][]int{{3}, {9, 20}, {15, 7}})
		So(levelOrder2(n), ShouldResemble, [][]int{{3}, {9, 20}, {15, 7}})
		n = &TreeNode{
			Val: 1,
		}
		So(levelOrder(n), ShouldResemble, [][]int{{1}})
		So(levelOrder2(n), ShouldResemble, [][]int{{1}})
		So(levelOrder(nil), ShouldResemble, [][]int{})
		So(levelOrder2(nil), ShouldResemble, [][]int{})
	})
}


func TestLevelOrderBottom(t *testing.T){
	Convey("二叉树层序遍历II", t, func(){
		n := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		So(levelOrderBottom(n), ShouldResemble, [][]int{{15, 7},  {9, 20}, {3}})
		n = &TreeNode{
			Val: 1,
		}
		So(levelOrderBottom(n), ShouldResemble, [][]int{{1}})
		So(levelOrderBottom(nil), ShouldResemble, [][]int{})
	})
}


func TestRightSideView(t *testing.T){
	Convey("二叉树的右视图", t, func(){
		n := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
			
		}
		So(rightSideView(n), ShouldResemble, []int{1, 3, 4})

		n = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
			},
		}
		So(rightSideView(n), ShouldResemble, []int{1, 3})
		So(rightSideView(nil), ShouldResemble, []int{})
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		So(rightSideView(n), ShouldResemble, []int{1, 2})
	})
}

func TestInverTree(t *testing.T){
	Convey("翻转二叉树", t, func(){
		n := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		}
		// So(invertTree(n), ShouldResemble, []int{1, 3, 4})

		n = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
			},
		}
		So(rightSideView(n), ShouldResemble, []int{1, 3})
		So(rightSideView(nil), ShouldResemble, []int{})
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		So(rightSideView(n), ShouldResemble, []int{1, 2})
	})
}


func TestNTreeTranversal(t *testing.T){
	Convey("N叉树前序遍历", t, func(){
		n := &Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 3,
					Children: []*Node {
						 { 
							Val: 5,
						},
						 {
							Val: 6,
						},
					},
				},
				{
					Val: 2,
				},
				{
					Val: 4,
				},
			},
		}
		So(preorderN(n), ShouldResemble, []int{1, 3, 5, 6, 2, 4})
		So(preorderN2(n), ShouldResemble, []int{1, 3, 5, 6, 2, 4})
	})
}

func TestPostOrderN(t *testing.T){
	Convey("N叉树后序遍历", t, func(){
		n := &Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 3,
					Children: []*Node {
						 { 
							Val: 5,
						},
						 {
							Val: 6,
						},
					},
				},
				{
					Val: 2,
				},
				{
					Val: 4,
				},
			},
		}
		So(postorderN(n), ShouldResemble, []int{5, 6, 3, 2, 4, 1})
	})
}


func TestIsSymmetric(t *testing.T){
	Convey("对称二叉树", t, func(){
		n := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		}
		So(isSymmetric(n), ShouldBeTrue)
		So(isSymmetric2(n), ShouldBeTrue)
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		}
		So(isSymmetric(n), ShouldBeFalse)
		So(isSymmetric2(n), ShouldBeFalse)
	})
}

func TestIsSame(t *testing.T){
	Convey("相同的树", t, func(){
		p := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		q := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		So(isSameTree(q, p), ShouldBeTrue)
	})
}


func TestIsSubTree(t *testing.T) {
	Convey("另一个颗树的子树", t, func(){
		n1 := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		}
		n2 := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}
		So(isSubtree(n1, n2), ShouldBeTrue)
		So(isSubtree(n1, n2), ShouldBeTrue)


		n1 = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				
			},
		}
		n2 = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Left:  &TreeNode{
						Val: 1,
					},
					Right:  &TreeNode{
						Val: 2,
					},
				},
			},
		}
		So(isSubtree(n1, n2), ShouldBeTrue)

		n1 = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
								Right: &TreeNode{
									Val: 1,
									Right: &TreeNode{
										Val: 1,
										Right: &TreeNode{
											Val: 1,
											Right: &TreeNode{
												Val: 1,
												Left: &TreeNode{
													Val: 1,
												},
												Right: &TreeNode{
													Val: 2,
												},
											},
										},
									},
								},
							},
						},
					},
				},
				
			},
		}
		n2 = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Left:  &TreeNode{
						Val: 1,
					},
					Right:  &TreeNode{
						Val: 2,
					},
				},
			},
		}
		So(isSubtreeR(n1, n2), ShouldBeTrue)
	})
}