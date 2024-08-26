package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxDepth(t *testing.T){
	Convey("二叉树的最大深度", t, func(){
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
		So(maxDepth(n), ShouldEqual, 3)

		n = &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		}
		So(maxDepth(n), ShouldEqual, 2)
	})
}



func TestMaxNDepth(t *testing.T){
	Convey("N叉树的最大深度", t, func(){
		n := &Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 2,
				},
				{
					Val: 3,
					Children: []*Node{
						{
							Val: 6,
						},
						{
							Val: 7,
							Children: []*Node{
								{
									Val: 11,
									Children: []*Node{
										{
											Val: 14,
										},
									},
								},
							},
						},
					},
				},
			},
		}
		So(maxNDepth(n), ShouldEqual, 5)
	})
}



func TestMinDepth(t *testing.T){
	Convey("二叉树的最小深度", t, func(){
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
		So(minDepth(n), ShouldEqual, 2)
		n = &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		}
		So(minDepth(n), ShouldEqual, 5)
	})
}

func TestIsBanlanced(t *testing.T) {
	Convey("平衡二叉树", t, func(){
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
		So(isBalanced(n), ShouldBeTrue)
		So(isBalanced2(n), ShouldBeTrue)
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		}
		So(isBalanced(n), ShouldBeFalse)
		So(isBalanced2(n), ShouldBeFalse)

		n = &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		So(isBalanced2(n), ShouldBeTrue)
		So(isBalanced(n), ShouldBeTrue)

		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}
		So(isBalanced2(n), ShouldBeFalse)
		So(isBalanced(n), ShouldBeFalse)

	})
}

func TestBinaryTreePaths(t *testing.T){
	Convey("二叉树的所有路径", t, func(){
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
			},
		}
		So(binaryTreePaths(n), ShouldResemble, []string{"1->2->5","1->3"})
		So(binaryTreePaths2(n), ShouldResemble, []string{"1->2->5","1->3"})

		n = &TreeNode{
			Val: 1,
		}
		So(binaryTreePaths(n), ShouldResemble, []string{"1"})
		So(binaryTreePaths2(n), ShouldResemble, []string{"1"})
	})
}