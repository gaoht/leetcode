package leetcode

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseList(t *testing.T){
	Convey("反转链表", t, func(){
		head := buildList([]int{1,2,3,4,5})
		So([]int{5,4,3,2,1}, ShouldResemble, listToArray(reverseList(head)))
	})	
}

func TestRemoveElements(t *testing.T){
	Convey("删除元素", t, func(){
		head := buildList([]int{7,7,7,7})
		So([]int{}, ShouldResemble, listToArray(removeElements(head, 7)))
	})	
}