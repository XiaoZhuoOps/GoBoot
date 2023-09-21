package goExample

import (
	"fmt"
	"testing"
)

func a() bool {
	fmt.Println("a")
	return false
}

func b() bool {
	fmt.Println("b")
	return false
}

func c() bool {
	fmt.Println("c")
	return true
}

func TestLogicOpPriority(t *testing.T) {
	// ! > && > ||
	// 从左到右依次执行
	// 如果能判断整个表达式的结果会提前停止
	res := a() && b() || c()
	// a && (b || c)
	// (a && b) || c YES
	fmt.Println(res)
	res = a() || b() && c()
	// (a || b) && c false
	// a || (b && c) false YES
	fmt.Println(res)
}
