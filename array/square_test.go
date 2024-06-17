package array

import (
	"fmt"
	"testing"
)

// 977
func TestSquare(t *testing.T){
	// []{-4, 0, 1, 2, 3}
	a := []int{-4, 0, 1, 2, 3}
	ret := Square(a) 
	for _, i := range ret{
		fmt.Println(i)
	}
}