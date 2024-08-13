package leetcode

import (
	"container/heap"
)

// [347] 前 K 个高频元素


func topKFrequentNum(nums []int, k int) []int {
	ans := make([]int, 0, k)
	iHeap := &FHeap{
		nums: make([][2]int, 0),
	}
	heap.Init(iHeap)
	frequency := make(map[int]int)
	for _, n := range nums {
		frequency[n]++
	}
	for n, v := range frequency {
		heap.Push(iHeap, [2]int{n, v})
		if iHeap.Len() > k {
			heap.Pop(iHeap)
		}
	}


	for iHeap.Len() > 0 {
		ans = append(ans, heap.Pop(iHeap).([2]int)[0])
	}
	return ans
}

type FHeap struct{
	nums [][2]int
}

func (h *FHeap) Push(x interface{}) {
	h.nums = append(h.nums, x.([2]int))

}

func (h *FHeap) Pop() interface{} {
	l := len(h.nums) 
	if l == 0 {
		return nil
	}
	x := h.nums[l - 1]
	h.nums = h.nums[: l-1]
	return x
}

func (h *FHeap) Len() int {
	return len(h.nums)
}

func (h *FHeap) Less(i int, j int) bool{
	return h.nums[i][1] < h.nums[j][1]
}

func (h *FHeap) Swap(i int, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

