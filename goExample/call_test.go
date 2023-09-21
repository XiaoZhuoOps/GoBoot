package goExample

import (
	"fmt"
	"testing"
)

func add(arr []string) {
	arr[0] = "1"
}

func addMap(m map[string]string) {
	m["1"] = "1"
}
func Test_call(t *testing.T) {
	// slice
	arr := make([]string, 10)
	add(arr)
	fmt.Println(arr)

	//map
	m := map[string]string{}
	addMap(m)
	fmt.Println(m)
}

func addPtr(arrPtr *[]string) {
	*arrPtr = append(*arrPtr, "add element")
}

func Test_slice_call(t *testing.T) {
	// slice传递时尽量用显式引用
	arr := []string{}
	addPtr(&arr)
	fmt.Println(arr)
}

func TestFuncMapCall(t *testing.T) {
	funcA := func() {
		t.Log("funcA")
	}
	m := map[string]func(){
		"func": funcA,
	}
	m["func"]()
	funcA = func() {
		t.Log("funcB")
	}
	m["func"]()
}
