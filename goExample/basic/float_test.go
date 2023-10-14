package basic

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_maxFloat(t *testing.T) {
	v, err := strconv.ParseFloat("1.8e400", 64)
	// +Inf float64
	fmt.Println(v, reflect.TypeOf(v))
	// strconv.ParseFloat: parsing "1.8e400": value out of range
	fmt.Println(err)
}
