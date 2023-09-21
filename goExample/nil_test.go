package goExample

import (
	"fmt"
	"testing"
)

type bmw struct {
	a int
	b string
	c []string
	// 组成
	// 顺序
}

func (b bmw) toString() {
	fmt.Println(b.a, b.b, b.c)
}

func Test_zeroAndNil(t *testing.T) {
	var a int
	//a = nil compile error
	fmt.Println(a)
	var b []string
	fmt.Println(b)
	fmt.Println(b == nil)

	var car bmw
	fmt.Println(car)
}
