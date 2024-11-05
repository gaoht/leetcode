package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestLRU(t *testing.T) {
	Convey("LRU", t, func(){
		lru := Constructor(2)
		lru.Put(1, 1)
		lru.Put(2, 2)
		So(lru.Get(1), ShouldEqual, 1)
		lru.Put(3, 3)
		So(lru.Get(2), ShouldEqual, -1)
		lru.Put(4,4)
		So(lru.Get(1), ShouldEqual, -1)
		So(lru.Get(3), ShouldEqual, 3)
		So(lru.Get(4), ShouldEqual, 4)

	})	
}