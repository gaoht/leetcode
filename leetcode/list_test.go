package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveElements(t *testing.T){
	Convey("测试移除", t, func(){
		l := newList([]int{7,7,7,7})
		l = removeElements(l, 7)	
		So([]int{}, ShouldResemble, listToArray(l))

	})
}

func TestReverseList(t *testing.T){
	Convey("测试反转", t, func(){
		l := newList([]int{1,2,3,4,5})
		l = reverseList(l)
		So([]int{5,4,3,2,1}, ShouldResemble, listToArray(l))
	})
}

func TestSwapPairs(t *testing.T){
	Convey("测试两两交换", t, func(){
		// 输入：head = [1]
		// 输出：[1]
		So([]int{1}, ShouldResemble, listToArray(swapPairs(newList([]int{1}))))
		// 输入：head = []
		// 输出：[]
		So([]int{}, ShouldResemble, listToArray(swapPairs(newList([]int{}))))
		//输入：head = [1,2,3,4]
		// 输出：[2,1,4,3]
		So([]int{2,1,4,3}, ShouldResemble, listToArray(swapPairs(newList([]int{1,2,3,4}))))
	})
}