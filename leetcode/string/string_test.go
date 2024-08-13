package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestReverseString(t *testing.T) {
	Convey("反转字符串", t, func(){
		bytes := []byte{'h', 'e', 'l', 'l', 'o'}
		reverseString(bytes)
		So(bytes, ShouldResemble, []byte{'o', 'l', 'l', 'e', 'h'})
	})
}


func TestReverseString2(t *testing.T) {
	Convey("反转字符串II", t, func(){
		So(reverseStr("abcdefg", 2), ShouldEqual, "bacdfeg")
		So(reverseStr("abcd", 2), ShouldEqual, "bacd")
	})
}


func TestReverseWords(t *testing.T) {
	Convey("反转字符串的单词", t, func(){
		So(reverseWords("the sky is blue"), ShouldEqual, "blue is sky the")
		So(reverseWords("  hello world  "), ShouldEqual, "world hello")
	})
}

func TestRepeatedSubstringPattern(t *testing.T) {
	Convey("重复的子字符串", t, func(){
		So(repeatedSubstringPattern("abab"), ShouldBeTrue)
		So(repeatedSubstringPattern("aba"), ShouldBeFalse)
		So(repeatedSubstringPattern("abcabcabcabc"), ShouldBeTrue)
		So(repeatedSubstringPattern("aabaaba"), ShouldBeFalse)
		So(repeatedSubstringPattern("abac"), ShouldBeFalse)
	})
}