package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSegments(t *testing.T){
	Convey("Segments test 1", t, func(){
		sg := NewIntensitySegments()
		sg.Add(10, 30, 1)
		So(sg.ToString(), ShouldEqual, "[[10,1],[30,0]]")
		sg.Add(20, 40, 1)
		So(sg.ToString(), ShouldEqual, "[[10,1],[20,2],[30,1],[40,0]]")
		sg.Add(10, 40, -2)
		So(sg.ToString(), ShouldEqual, "[[10,-1],[20,0],[30,-1],[40,0]]")
	})

	Convey("Segments test 1", t, func(){
		sg := NewIntensitySegments()
		sg.Add(10, 30, 1)
		So(sg.ToString(), ShouldEqual, "[[10,1],[30,0]]")
		sg.Add(30, 40, 1)
		So(sg.ToString(), ShouldEqual, "[[10,1],[30,1],[40,0]]")
	})


	Convey("Segments test 2", t, func(){
		sg := NewIntensitySegments()
		sg.Add(10, 30, 1)
		So(sg.ToString(), ShouldEqual, "[[10,1],[30,0]]")
		sg.Add(20, 40, 1)
		So(sg.ToString(), ShouldEqual, "[[10,1],[20,2],[30,1],[40,0]]")
		sg.Add(10, 40, -1)
		So(sg.ToString(), ShouldEqual, "[[20,1],[30,0]]")
		sg.Add(10, 40, -1)
		So(sg.ToString(), ShouldEqual, "[[10,-1],[20,0],[30,-1],[40,0]]")
	})
}

func TestCaculator(t *testing.T) {
	Convey("计算器", t, func(){
		So(calculate("1 + 1"), ShouldEqual, 2)
		So(calculate("12 + 1"), ShouldEqual, 13)
		So(calculate(" 2-1 + 2 "), ShouldEqual, 3)
		So(calculate("-3"), ShouldEqual, -3)
		So(calculate("(1+(4+5+2)-3)+(6+8)"), ShouldEqual, 23)
		
	})
}

func TestInCreaseDigits(t *testing.T) {
	Convey("单调递增", t, func(){
		So(monotoneIncreasingDigits(332), ShouldEqual, 299)
		
	})
}


func TestBag(t *testing.T){
	Convey("背包", t, func(){
		v := bagSolve(1, 6, []int{2,2,3,1,5,2}, []int{2,3,1,5,4,3})
		So(v, ShouldEqual, 5)
		v = bagSolve2D(1, 6, []int{2,2,3,1,5,2}, []int{2,3,1,5,4,3})
		So(v, ShouldEqual, 5)
		v = bagSolve(3, 6, []int{2,2,3,1,5,2}, []int{2,3,1,5,4,3})
		So(v, ShouldEqual, 8)
		v = bagSolve2D(3, 6, []int{2,2,3,1,5,2}, []int{2,3,1,5,4,3})
		So(v, ShouldEqual, 8)
	})
}


func TestFindTargetSums(t *testing.T){
	Convey("目标和", t, func(){
		So(findTargetSumWays([]int{1,1,1,1,1}, 3), ShouldEqual, 5)
		So(findTargetSumWays([]int{1,0}, 1), ShouldEqual, 2)
	})
}

func TestChange(t *testing.T) {
	Convey("零钱兑换", t, func(){
		So(change(5, []int{1,2,5}), ShouldEqual, 4)
		So(change(5, []int{2,5}), ShouldEqual, 1)
	})
}

func TestWordBreak(t *testing.T) {
	Convey("单词拆分", t, func(){
		So(wordBreak("leetcode", []string{"leet", "code"}), ShouldBeTrue)
		So(wordBreak("applepenapple", []string{"pen", "apple"}), ShouldBeTrue)
		So(wordBreak("cars", []string{"car", "ca", "rs"}), ShouldBeTrue)
		So(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}), ShouldBeTrue)
		So(wordBreak("catsandog", []string{"cats","dog","sand","and","cat"}), ShouldBeFalse)		
		
		So(wordBreakError("leetcode", []string{"leet", "code"}), ShouldBeTrue)
		So(wordBreakError("appleapplepen", []string{"apple", "pen"}), ShouldBeTrue)
		So(wordBreakError("applepenapple", []string{"apple", "pen"}), ShouldBeTrue)
		
		So(wordBreakError("cars", []string{"car", "ca", "rs"}), ShouldBeTrue)
		So(wordBreakError("aaaaaaa", []string{"aaaa", "aaa"}), ShouldBeTrue)
		So(wordBreakError("catsandog", []string{"cats","dog","sand","and","cat"}), ShouldBeFalse)
	})
}

func TestRob(t *testing.T){
	Convey("打家劫舍", t, func(){
		So(rob([]int{2,1,1,2}), ShouldEqual, 4)
		// So(rob2([]int{1,1,1,2}), ShouldEqual, 3)
		So(rob2([]int{1,1,3,6,7,10,7,1,8,5,9,1,4,4,3}), ShouldEqual, 3)
	})
}

func TestLengthOfLIS(t *testing.T){
	Convey("最长递增子序列", t, func(){
		So(lengthOfLIS([]int{0,1,0,3,2,3}), ShouldEqual, 4)
	})
}

func TestNumsDistinct(t *testing.T) {
	Convey("不同的子序列", t, func(){
		So(numDistinct("rabbbiit", "rabbit"), ShouldEqual, 6)
	})
}
func TestMinDistance(t *testing.T) {
	Convey("两个字符串的删除操作", t, func(){
		So(minDistance("leetcode", "etco"), ShouldEqual, 4)
	})
}

func TestLongestPalindromeSubseq(t *testing.T) {
	Convey("最长回文子序列", t, func(){
		So(longestPalindromeSubseq("bbbab"), ShouldEqual, 4)
		So(longestPalindromeSubseq("cbbd"), ShouldEqual, 2)
	})
}

func TestCandy(t *testing.T) {
	Convey("分发糖果", t, func(){
		So(candy([]int{1,3,4,5,2}), ShouldEqual, 11)
	})
}

func TestTrap(t *testing.T){
	Convey("接雨水", t, func ()  {
		So(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}), ShouldEqual, 6)
	})
	Convey("接雨水2", t, func ()  {
		So(trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1}), ShouldEqual, 6)
	})
}


func TestNextGreaterElement(t *testing.T){
	Convey("下一个更大元素", t, func ()  {
		So(nextGreaterElement([]int{4,1,2}, []int{1,3,4,2}), ShouldResemble, []int{-1, 3, -1})
	})
}



