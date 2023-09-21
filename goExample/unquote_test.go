package goExample

import (
	. "code.byted.org/gopkg/mockito"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestUnquote(t *testing.T) {
	PatchConvey("", t, func() {
		str := `"[$27]"`
		Convey("1", func() {
			flag := false
			if _, err := strconv.Unquote(str); err == nil {
				flag = true
			}
			So(flag, ShouldBeTrue)
		})
		Convey("2", func() {
			str := `[$27]`
			flag := false
			if _, err := strconv.Unquote(str); err == nil {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
		Convey("3", func() {
			str := `[$27]"`
			flag := false
			if _, err := strconv.Unquote(str); err == nil {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
		Convey("4", func() {
			str := `["$27"]`
			flag := false
			if _, err := strconv.Unquote(str); err == nil {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
	})
}
