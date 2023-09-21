package goExample

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"testing"
)

func Test_sonic(t *testing.T) {
	str := `{"value":"[1]"}`
	node, err := sonic.GetFromString(str, "value")
	if err != nil {
		fmt.Println(err)
		return
	}
	if node.Type() == 5 {
		nodeList, err := node.ArrayUseNode()
		if err != nil {
			fmt.Println(err)
		}
		n := len(nodeList)
		fmt.Println(n)
		if n == 1 {
			newNode := nodeList[0]
			switch newNode.Type() {
			case ast.V_STRING:
				fmt.Println(newNode.String())
			}
		}
	}
}
