package goExample

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// 测试代码：
func Test_Fib(t *testing.T) {
	var (
		in       = 7  //测试需要输入的参数
		expected = 13 //期望函数返回的结果
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
