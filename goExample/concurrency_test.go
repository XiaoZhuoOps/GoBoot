package goExample

import (
	"code.byted.org/gopkg/mockito"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

var (
	once   sync.Once
	oneInt int32
)

func testOnce() {
	once.Do(func() {
		fmt.Println("only once")
	})
}

func TestOnce(t *testing.T) {
	//for i := 0; i < 10; i++ {
	//	once.Do(func() {
	//		fmt.Println("just only once", oneInt)
	//	})
	//}
	mockito.PatchConvey("test once", t, func() {
		Convey("1", func() {
			testOnce()
		})
		Convey("2", func() {
			testOnce()
		})
	})
}

func TestOnceTwo(t *testing.T) {
	mockito.PatchConvey("test once two", t, func() {
		Convey("1", func() {
			testOnce()
		})
		Convey("2", func() {
			testOnce()
		})
	})
}
