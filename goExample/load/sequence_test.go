package load

import (
	"fmt"
	"testing"
)

func init() {
	PrintSeq("c")
}

var (
	b = PrintSeq("b")

	a = PrintSeq("a")
)

const consts = "-consts"

func PrintSeq(who string) string {
	fmt.Println(who + consts)
	return ""
}

func TestSequence(t *testing.T) {
	fmt.Println("")
}
