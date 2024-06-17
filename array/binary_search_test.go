package array

import (
	"testing"
)

func TestBinarySearch(t *testing.T){
	a := []int{-1, 0, 1, 2, 3, 4}
	if(BinarySearch(a, -1) != 0) {
		t.Errorf("binary search, target=%d, index should be %d", -1, 0)
	}
	if(BinarySearch(a, 2) != 3) {
		t.Errorf("binary search, target=%d, index should be %d", 2, 3)
	}
	if(BinarySearch(a, 4) != 5) {
		t.Errorf("binary search, target=%d, index should be %d", 4, 5)
	}
	if(BinarySearch(a, 5) != -1) {
		t.Errorf("binary search, target=%d, index should be %d", 5, -1)
	}
}