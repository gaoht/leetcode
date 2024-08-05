package leetcode

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpiralOrder(t *testing.T){
	Convey("测试螺旋矩阵", t, func(){
		matrix := [][]int{
			{1,2,3,4},
			{5,6,7,8},
			{9,10,11,12},
		}
		nums := spiralOrder(matrix)
		fmt.Println(nums)
		So(nums, ShouldResemble, []int{1,2,3,4,8,12,11,10,9,5,6,7})
	})
}

