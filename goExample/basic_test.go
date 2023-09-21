package goExample

import (
	"fmt"
	"math"
	"testing"
)

func BasicTest() {
	// variables
	var a string = "12"
	const b string = "123"
	fmt.Println(a)
	fmt.Println(b)

	// map
	// Attention: 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，
	// 该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。
	// 这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略。
	vtoi := make(map[int]int)
	vtoi[1] = 2
	r1, r2 := vtoi[1]
	fmt.Print(r1, r2)
	r1, r2 = vtoi[2]
	fmt.Print(r1, r2)

	// range
	for k, v := range vtoi {
		fmt.Println(k, v)
	}

}

func Test_string(t *testing.T) {
	// string
	// 不可变字节序列
	// 内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。
	s := "\"1\""
	fmt.Println(s, "len(s) is", len(s))
	for i, r := range s {
		fmt.Println(i, r)
	}
}

func intTest() {
	// uint int
	var a uint32 = 1
	var b int32 = -1
	fmt.Print(int32(a))
	fmt.Print(uint32(b))
}

func Test_float(t *testing.T) {
	// float
	// 浮点数运算是不精确的，谨慎使用==
	var a float64 = 1111.1
	fmt.Println(a * math.Pow10(5))
}

func Test_type(t *testing.T) {
	//	type
	type benz int32
	type bmw int32
	var car1 benz = 0
	var car2 bmw = 0
	//fmt.Println(car1 == car2)  compile error
	fmt.Println(car1 == benz(car2))
	fmt.Println(car1 == 0)
}

func returnValue(a interface{}) interface{} {
	return a
}
func Test_interface(t *testing.T) {
	//interface{}
	switch returnValue(1).(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	}
}

func Test_slice(t *testing.T) {
	// slice
	aslice := []string{}
	var bslice []string
	var cslice []string = nil
	fmt.Println(aslice, aslice == nil)
	fmt.Println(bslice, bslice == nil)
	fmt.Println(cslice, cslice == nil)

	for _, ele := range bslice {
		println(ele, "Yes")
	}
	println("No")

	//bslice[0] = "1" // will panic index out of bound

	bslice = append(bslice, "1")

	for i, b := range bslice {
		fmt.Println("===========", i, b)
	}

	fmt.Printf("%p", &bslice)
	appendSlice(bslice, "2")
}

func appendSlice(arr []string, ele string) {
	fmt.Printf("%p", &arr)
	arr = append(arr, ele)
}

func Test_map(t *testing.T) {
	// map
	// Attention: 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，
	// 该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。
	// 这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略。
	vtoi := make(map[int]int)
	vtoi[1] = 2
	r1, r2 := vtoi[1]
	fmt.Print(r1, r2)
	r1, r2 = vtoi[2]
	fmt.Print(r1, r2)

	//	nil-map
	var nilMap map[string]bool
	fmt.Println(nilMap["0"])
	nilMap["0"] = true // will panic, assignment to entry in nil map
}
func Test_func(t *testing.T) {
	var f func(string)
	f("string") //panic, invalid memory address or nil pointer dereference
}

type para string

func TestAlign(t *testing.T) {
	var v para = "1"
	m := map[para]string{
		v: "1",
	}
	fmt.Println(m["1"])
	fmt.Println(m["2"])
}
