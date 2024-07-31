package tree

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewMaxHeap()
	h.Insert(3)
	h.Insert(5)
	h.Insert(7)
	h.Insert(9)
	h.Insert(10)
	h.Insert(6)
	h.Insert(66)
	for i, n := range h.nums {
		if i == 0 {
			continue
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	
	for m, b := h.ExtractMax(); b;  m, b = h.ExtractMax(){
		fmt.Printf("%d ", m)
	}
}

func TestMinHeap(t *testing.T){
	h := NewMinHeap()
	h.Insert(3)
	h.Insert(5)
	h.Insert(7)
	h.Insert(9)
	h.Insert(10)
	h.Insert(6)
	h.Insert(66)
	for i, n := range h.nums {
		if i == 0 {
			continue
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	
	for m, b := h.ExtractMin(); b;  m, b = h.ExtractMin(){
		fmt.Printf("%d ", m)
	}
}