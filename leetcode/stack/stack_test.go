package leetcode

import (
	"container/heap"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsValidBrace(t *testing.T) {
	Convey("有效的括号", t, func(){
		So(isValid("()"), ShouldBeTrue)
		So(isValid("()[]{}"), ShouldBeTrue)
		So(isValid("(]"), ShouldBeFalse)
		So(isValid("]"), ShouldBeFalse)
	})
}

func TestRemoveDuplicates(t *testing.T) {
	Convey("删除字符串中的所有相邻重复项", t, func(){
		So(removeDuplicates("abbaca"), ShouldEqual, "ca")
	})
}

func TestMaxSlideWindow(t *testing.T) {
	Convey("滑动窗口最大值", t, func(){
		So(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3), ShouldResemble, []int{3,3,5,5,6,7})
		So(maxSlidingWindow([]int{1}, 1), ShouldResemble, []int{1})
	})
}

func TestTopKFrequent(t *testing.T){
	Convey("前K个高频元素", t, func(){
		iHeap := &FHeap{
			nums: make([][2]int, 0),
		}
		heap.Init(iHeap)
		heap.Push(iHeap, [2]int{1,1})
		heap.Push(iHeap, [2]int{2,2})
		heap.Push(iHeap, [2]int{3,3})

		So(heap.Pop(iHeap), ShouldEqual, [2]int{1,1})
		So(iHeap.Len(), ShouldEqual, 2)
		So(topKFrequentNum([]int{1,1,1,2,2,3}, 2), ShouldResemble, []int{2,1})
		So(topKFrequentNum([]int{1}, 1), ShouldResemble, []int{1})
		So(topKFrequentNum([]int{-1,-1}, 1), ShouldResemble, []int{-1})
	})
}

func TestTopKFrequentWords(t *testing.T){
	Convey("前K个高频单词", t, func(){
		// tHeap := &THeap{
		// 	words: make([]WordCount, 0),
		// }
		// heap.Init(tHeap)
		// heap.Push(tHeap, WordCount{
		// 	"i", 2,
		// })
		// heap.Push(tHeap, WordCount{
		// 	"you", 1,
		// })
		// heap.Push(tHeap, WordCount{
		// 	"love", 2,
		// })
		// So(tHeap.Len(), ShouldEqual, 3)
		// So(heap.Pop(tHeap).(WordCount).word, ShouldEqual, "you")
		// So(heap.Pop(tHeap).(WordCount).word, ShouldEqual, "i")
		// So(heap.Pop(tHeap).(WordCount).word, ShouldEqual, "love")

		So(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2), ShouldResemble, []string{"i", "love"})
		So(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4), ShouldResemble, []string{"the", "is", "sunny", "day"})

	})
}