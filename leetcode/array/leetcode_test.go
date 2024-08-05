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

func TestMinSubArrayLen(t *testing.T){
	Convey("测试最小子数组", t, func(){
		// target = 7, nums = [2,3,1,2,4,3]  2
		// target = 4, nums = [1,4,4] 1
		// target = 11, nums = [1,1,1,1,1,1,1,1]
		So(minSubArrayLen(7, []int{2,3,1,2,4,3}), ShouldEqual, 2)
		So(minSubArrayLen(4, []int{1,4,4}), ShouldEqual, 1)
		So(minSubArrayLen(11, []int{1,1,1,1,1,1,1,1}), ShouldEqual, 0)
	})
}
