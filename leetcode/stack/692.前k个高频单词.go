package leetcode

import (
	"container/heap"
)

// [692] 前K个高频单词


func topKFrequent(words []string, k int) []string {
	frequency := make(map[string]int)
	tHeap := &THeap{
		words: make([]WordCount, 0),
	}
	for _, w := range words {
		frequency[w]++
	}
	heap.Init(tHeap)
	for w, c := range frequency {
		heap.Push(tHeap, WordCount{w, c})
	}
	ans := make([]string, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(tHeap).(WordCount).word)
	}
	return ans
}
type WordCount struct {
	word string
	count int
}
type THeap struct{
	words []WordCount
}

func (t *THeap) Less(i, j int) bool{
	if t.words[i].count > t.words[j].count {
		return true
	} else if t.words[i].count == t.words[j].count {
		return t.words[i].word < t.words[j].word 
	} else {
		return false
	}
}

func (t *THeap) Swap(i, j int) {
	t.words[i], t.words[j] = t.words[j], t.words[i]
}

func (t *THeap) Len() int {
	return len(t.words)
}

func (t *THeap) Pop() interface{}{
	if t.Len() == 0 {
		return nil
	}
	x := t.words[t.Len() - 1]
	t.words = t.words[0 : t.Len() - 1]
	return x
}

func (t *THeap) Push(x interface{}) {
	t.words = append(t.words, x.(WordCount))
} 