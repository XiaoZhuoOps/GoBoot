package goExample

import (
	"code.byted.org/aweme-go/ajson"
	. "code.byted.org/gopkg/mockito"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSonicArray(t *testing.T) {
	PatchConvey("", t, func() {
		str := "[$27]"
		Convey("1", func() {
			flag := false
			if arr, err := ajson.GetFromString(str).Array(); err == nil && len(arr) == 1 {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
		Convey("2", func() {
			str = "[\"$27\"]"
			flag := false
			if arr, err := ajson.GetFromString(str).Array(); err == nil && len(arr) == 1 {
				flag = true
			}
			So(flag, ShouldBeTrue)
		})
		Convey("3", func() {
			str = "[2]"
			flag := false
			if arr, err := ajson.GetFromString(str).Array(); err == nil && len(arr) == 1 {
				flag = true
			}
			So(flag, ShouldBeTrue)
		})
		Convey("4", func() {
			str = "[2,5,7]"
			flag := false
			if arr, err := ajson.GetFromString(str).Array(); err == nil && len(arr) == 1 {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
		Convey("5", func() {
			str = "undefined[1].1"
			flag := false
			if arr, err := ajson.GetFromString(str).Array(); err == nil && len(arr) == 1 {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
		Convey("6", func() {
			str = "\"[1]\""
			flag := false
			if arr, err := ajson.GetFromString(str).Array(); err == nil && len(arr) == 1 {
				flag = true
			}
			So(flag, ShouldBeFalse)
		})
	})
}
