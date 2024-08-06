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

func TestRemoveNthFromEnd(t *testing.T) {
	Convey("删除倒数的链表节点", t, func(){
		// 	输入：head = [1,2,3,4,5], n = 2
		// 输出：[1,2,3,5]
		So(listToArray(removeNthFromEnd(newList([]int{1,2,3,4,5}), 2)), ShouldResemble, []int{1,2,3,5})
		// 输入：head = [1], n = 1
		// 输出：[]
		So(listToArray(removeNthFromEnd(newList([]int{1}), 1)), ShouldResemble, []int{})
		// 输入：head = [1,2], n = 1
		// 输出：[1]
		So(listToArray(removeNthFromEnd(newList([]int{1,2}), 1)), ShouldResemble, []int{1})
	})
}

func TestGetIntersectionNode(t *testing.T){
	Convey("相交链表", t, func(){
		// listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
		a := newList([]int{4})
		b := newList([]int{5,6})
		c := newList([]int{1,8,4,5})
		a.Next = c
		b.Next.Next = c
		So(getIntersectionNode(a, b), ShouldEqual, c)
		// listA = [1,9,1,2,4], listB = [3,2,4]
		a = newList([]int{1,9,1})
		b = newList([]int{3})
		c = newList([]int{2,4})
		a.Next.Next.Next = c
		b.Next = c
		So(getIntersectionNode(a, b), ShouldEqual, c)
		// listA = [2,6,4], listB = [1,5]
		a = newList([]int{2,6,4})
		b = newList([]int{1,5})
		So(getIntersectionNode(a, b), ShouldBeNil)
	})
}

func TestHasCircle(t *testing.T){
	Convey("链表是否有环", t, func(){
		// head = [3,2,0,-4], pos = 1
		head := newList([]int{3,2,0,-4})
		head.Next.Next.Next.Next = head.Next
		So(hasCycle(head), ShouldBeTrue)
		head = newList([]int{3,2,0,-4})
		So(hasCycle(head), ShouldBeFalse)
	})
}

func TestDetectCircle(t *testing.T){
	Convey("链表环检测", t, func(){
		// head = [3,2,0,-4], pos = 1
		head := newList([]int{3,2,0,-4})
		head.Next.Next.Next.Next = head.Next
		So(hasCycle(head), ShouldBeTrue)
		So(detectCycle(head).Val, ShouldEqual, 2)
	})
}