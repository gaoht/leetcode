package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDuplicates(t *testing.T){
	Convey("移除重复元素", t, func(){
		// 输入：nums = [1,1,1,2,2,3]
		// 输出：5, nums = [1,1,2,2,3]
		nums := []int{1,1,1,2,2,3}
		l := removeDuplicates(nums)
		So(l, ShouldEqual, 5)
		So(nums, ShouldResemble, []int{1,1,2,2,3,3})
	})
}

func TestMoveZeros(t *testing.T){
	Convey("移动0", t, func(){
		nums := []int{0, 1, 0, 3, 12}
		moveZeroes(nums)
		So(nums, ShouldResemble, []int{1,3,12,0,0})

		nums = []int{0}
		moveZeroes(nums)
		So(nums, ShouldResemble, []int{0})
	})
}

func TestBackspaceCompare(t *testing.T) {
	Convey("比较含退格的字符串", t, func(){
		So(backspace([]byte("ab#c")), ShouldResemble, []byte{'a', 'c'})
		So(backspace([]byte("ab##")), ShouldResemble, []byte{})
		So(backspace([]byte("a#c")), ShouldResemble, []byte{'c'})

		So(backspaceCompare("ab#c", "ad#c"), ShouldBeTrue)
		So(backspaceCompare("ab##", "c#d#"), ShouldBeTrue)
		So(backspaceCompare("a#c", "b"), ShouldBeFalse)
	})
}


func TestTotalFruits(t *testing.T) {
	Convey("水果成篮", t, func(){
		So(totalFruit([]int{1,2,1}), ShouldEqual, 3)
		So(totalFruit([]int{0,1,2,2}), ShouldEqual, 3)
		So(totalFruit([]int{1,2,3,2,2}), ShouldEqual, 4)
		So(totalFruit([]int{3,3,3,1,2,1,1,2,3,3,4}), ShouldEqual, 5)
		So(totalFruit([]int{0,1,6,6,4,4,6}), ShouldEqual, 5	)
	})
}

func TestMinWindow(t *testing.T) {
	Convey("最小覆盖子串", t, func(){
		So(minWindow("ADOBECODEBANC", "ABC"), ShouldEqual, "BANC")
		So(minWindow("a", "a"), ShouldEqual, "a")
		So(minWindow("a", "aa"), ShouldEqual, "")
	})
}



func TestGenerateMatrix(t *testing.T){
	Convey("螺旋矩阵", t, func(){
		So(generateMatrix(3), ShouldResemble, [][]int{{1,2,3},{8,9,4},{7,6,5}})
	})
}