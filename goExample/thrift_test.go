package goExample

import (
	"GoBoot/gen-go/constant"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConstant(t *testing.T) {
	convey.Convey("test constant", t, func() {
		fmt.Println(constant.VALID)
	})
}
