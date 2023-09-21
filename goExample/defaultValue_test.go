package goExample

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	Convey("1", t, func() {
		var s string
		So(s, ShouldEqual, "")
		fmt.Println("string=======", s)
		var arr []string
		So(arr, ShouldBeNil)
		//So(arr, ShouldEqual, nil)
		fmt.Println("[]string======", arr)
		fmt.Println(arr == nil)

		var l uint64
		So(l, ShouldEqual, 0)
		fmt.Println("align======", l)
	})
}
