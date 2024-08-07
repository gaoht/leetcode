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