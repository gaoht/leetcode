package array

import (
	"fmt"
	"testing"
)

func TestGetLeft(t *testing.T){
	a := []int{5,7,7,8,8,10}
	l := getLeft(a, 8)
	r := getRight(a, 8)
	fmt.Println(l, r)

	b := []int{2,2}
	ll := getLeft(b, 3)
	rr := getRight(b, 3)
	fmt.Println(ll, rr, searchRange(b, 3))
}