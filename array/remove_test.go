package array

import (
	"fmt"
	"testing"
)
func TestRemove(t *testing.T){
	a := []int{2,3,2,5}
	Remove(a, 2)
	for _, i := range a {
		fmt.Println(i)
	}
	if a[0] == 2 {
		t.Errorf("removed element exist, %d", 2)
	}
}