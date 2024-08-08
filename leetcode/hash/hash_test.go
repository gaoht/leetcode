package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsAnagram(t *testing.T){
	Convey("是否异位词", t, func(){
		So(isAnagram("anagram", "nagaram"), ShouldBeTrue)
	})
}

func TestGroupAnagrams(t *testing.T) {
	Convey("字母异位词分组", t, func(){
		So(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), ShouldResemble, [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		})
		So(groupAnagrams([]string{"", ""}), ShouldResemble, [][]string{
			{"", ""},
		})
		So(groupAnagrams([]string{"a"}), ShouldResemble, [][]string{
			{"a"},
		})
	})
}

func TestFindAnagrams(t *testing.T) {
	Convey("找到字符串中所有字母异位词", t, func(){
		So(findAnagrams("abab", "ab"), ShouldResemble, []int{0, 1, 2})
		So(findAnagrams("cbaebabacd", "abc"), ShouldResemble, []int{0, 6})
	})
}

func TestIsHappy(t *testing.T) {
	Convey("快乐数", t, func(){
		So(isHappy(19), ShouldBeTrue)
		So(isHappy(2), ShouldBeFalse)
		So(isHappy(1), ShouldBeTrue)
	})
}

func TestTwoSum(t *testing.T) {
	Convey("两数之和", t, func(){
		So(twoSum([]int{2,7,11,15}, 9), ShouldResemble, []int{0,1})
		So(twoSum([]int{3,2,4}, 6), ShouldResemble, []int{1,2})
		So(twoSum([]int{3,3}, 6), ShouldResemble, []int{0,1})
		So(twoSum([]int{-1,-2,-3,-4,-5}, -8), ShouldResemble, []int{2,4})
	})
}

func TestFourSumCount(t *testing.T) {
	Convey("四数相加", t, func(){
		So(fourSumCount([]int{1,2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}), ShouldEqual, 2)
		So(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}), ShouldEqual, 1)
	})
}

func TestThreeSum(t *testing.T) {
	Convey("三数相加", t, func(){
		So(threeSum([]int{-1,0,1,2,-1,-4}), ShouldResemble, [][]int{{-1,-1,2},{-1,0,1}})
		So(threeSum([]int{0,1,1}), ShouldResemble, [][]int{})
		So(threeSum([]int{0,0,0}), ShouldResemble, [][]int{{0,0,0}})
	})
}

func TestFourSum(t *testing.T) {
	Convey("四数相加", t, func(){
		So(fourSum([]int{1,0,-1,0,-2,2}, 0), ShouldResemble, [][]int{{-2,-1,1,2},{-2,0,0,2}, {-1,0,0,1}})
		So(fourSum([]int{2,2,2,2}, 8), ShouldResemble, [][]int{{2,2,2,2}})
		So(fourSum([]int{1,-2,-5,-4,-3,3,3,5}, -11), ShouldResemble, [][]int{{-5,-4,-3,1}})
	})
}