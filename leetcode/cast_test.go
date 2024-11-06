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
		So(wordBreakError("applepenapple", []string{"apple", "pen"}), ShouldBeTrue)
		So(wordBreakError("cars", []string{"car", "ca", "rs"}), ShouldBeTrue)
		So(wordBreakError("aaaaaaa", []string{"aaaa", "aaa"}), ShouldBeTrue)
		So(wordBreakError("catsandog", []string{"cats","dog","sand","and","cat"}), ShouldBeFalse)
	})
}