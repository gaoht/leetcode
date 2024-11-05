package leetcode

import (
	"fmt"
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

func TestSumOfLeftLeaves(t *testing.T){
	Convey("左叶子之和", t, func(){
		n := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val:  20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		So(sumOfLeftLeaves(n), ShouldEqual, 24)
		n = &TreeNode{
			Val: 1,
		}
		So(sumOfLeftLeaves(n), ShouldEqual, 0)
	})
}

func TestFindBottomLeftValue(t *testing.T){
	Convey("找树左下角的值", t, func(){
		n := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		So(findBottomLeftValue(n), ShouldEqual, 1)
		So(findBottomLeftValueR(n), ShouldEqual, 1)
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 7,
					},
				},
				Right:  &TreeNode{
					Val: 6,
				},
			},
		}
		So(findBottomLeftValue(n), ShouldEqual, 7)
		So(findBottomLeftValueR(n), ShouldEqual, 7)
	})
}

func TestHasPathSum(t *testing.T){
	Convey("路径总和", t, func(){
		n := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		}
		So(hasPathSum(n, 22), ShouldBeTrue)
		So(hasPathSumR(n, 22), ShouldBeTrue)
		n = nil
		So(hasPathSum(n, 0), ShouldBeFalse)
		So(hasPathSumR(n, 0), ShouldBeFalse)
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		So(hasPathSum(n, 5), ShouldBeFalse)
		So(hasPathSumR(n, 5), ShouldBeFalse)
		n = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		So(hasPathSum(n, 1), ShouldBeFalse)
		So(hasPathSumR(n, 1), ShouldBeFalse)
	})
}


func TestPathSum(t *testing.T){
	Convey("路径总和III", t, func(){
		n := &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: -2,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: -3,
				Right: &TreeNode{
					Val: 11,
				},
			},
		}
		So(pathSum(n, 8), ShouldEqual, 3)
	})
}

func TestBuildTree(t *testing.T) {
	Convey("从中序与后续遍历序列构造二叉树", t, func(){
		in := []int{2,3,1}
		post := []int{3,2,1}
		buildTree(in, post)

	})
}

func TestIsBanlanced3(t *testing.T) {
	Convey("平衡二叉树", t, func(){
		n := &TreeNode{
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
		d := getDepth3(n, 0)
		fmt.Printf("%d", d)
	})
	
}
func TestSortedArrayToBST(t *testing.T) {
	Convey("构建平衡搜索数", t, func(){
		nums := []int{-10,-3,0,5,9}
		SortedArrayToBST(nums)
	})
}

func TestFlatten(t *testing.T){
	Convey("Flatten Binary Tree to Linked List", t, func(){
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
				Val: 5,
				Right: &TreeNode{
					Val: 6,
				},
			},
		}
		flatten(n)
	})
}

func Test2DMatrix(t *testing.T){
	Convey("搜索二维矩阵", t, func(){
		matrix := [][]int{
			{1,1},
		}
		So(searchMatrix(matrix, 2), ShouldBeFalse)
		matrix = [][]int{
			{1,3},
		}
		So(searchMatrix(matrix, 3), ShouldBeTrue)
	})
}

func TestFindMin(t *testing.T){
	Convey("寻找旋转排序数组中的最小值", t, func(){
		// x := "abc"

		// fmt.Println(x[1:])
		// So(findMin([]int{3,4,5,1,2}), ShouldEqual, 1)
		permuteUnique([]int{1,1,2})
	})
}

func TestFindItinerary(t *testing.T){
	Convey("重新安排行程", t, func(){
		// x := "abc"
		// fmt.Println(x[1:])
		// So(findMin([]int{3,4,5,1,2}), ShouldEqual, 1)
		// fmt.Println(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}))
		// So(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}), ShouldResemble, []string{"JFK","MUC","LHR","SFO","SJC"})
		So(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}), ShouldResemble, []string{"JFK","ATL","JFK","SFO","ATL","SFO"})
	})
}

func TestSolveNQueens(t *testing.T){
	Convey("N皇后", t, func(){
		So(isValidQueens([]string{"Q"}, 0, 1), ShouldBeTrue)
		So(solveNQueens(1), ShouldResemble, [][]string{{"Q"}})
	})
}

func TestSolveSudoku(t *testing.T){
	Convey("解数独", t, func(){
		// x := "abc"
		// fmt.Println(x[1:])
		// So(findMin([]int{3,4,5,1,2}), ShouldEqual, 1)
		// fmt.Println(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}))
		// So(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}), ShouldResemble, []string{"JFK","MUC","LHR","SFO","SJC"})
		board := [][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}
		// board :=[][]byte{{'.','.','.'}, {'.','.','.'}, {'.','.','.'}}
	
		solveSudoku(board)
		printBoard(board)
		fmt.Println(loops)
		// printBoard(board)
	})
}


func TestCanJump(t *testing.T) {

	Convey("跳跃游戏", t, func(){
		So(canJump([]int{2,3,1,1,4}), ShouldBeTrue)
		So(canJump([]int{3,2,1,0,4}), ShouldBeFalse)
	})
}

func TestMaxProfit(t *testing.T) {
	Convey("买卖股票的最佳时机II", t, func(){
		So(maxProfit([]int{7,1,5,3,6,4}), ShouldEqual, 7)
	})
}


func TestStations(t *testing.T) {
	Convey("加油站", t, func(){
		So(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}), ShouldEqual, 3)
	})
}

func TestReconstructHeight(t *testing.T) {
	Convey("重建身高队列", t, func(){
		reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}})
	})
}

func TestMaxPathSum(t *testing.T) {
	Convey("最大路径和", t, func(){
		a0, a1, a2 := 1, 2, 3
		root := buildBinaryTreeFromArray([]*int{&a0, &a1, &a2,})
		So(maxPathSum(root), ShouldEqual, 6)
		a0, a1, a2, a3, a4, a5, a7 := 1,-2,-3,1,3,-2,-1
		root = buildBinaryTreeFromArray([]*int{&a0, &a1, &a2, &a3, &a4, &a5, nil, &a7})
		So(maxPathSum(root), ShouldEqual, 3)
	})
}