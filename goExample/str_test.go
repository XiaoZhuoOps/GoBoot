package goExample

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_ParseStr(t *testing.T) {
	//str := "1_5."
	str := "0.000000"
	float, err := String2Float(str, 64)
	if err == nil {
		fmt.Println("float is ", float*2)
	} else {
		fmt.Println("error", err)
	}
}

func String2Float(stringNumber string, bitSize int) (float64, error) {
	float, err := strconv.ParseFloat(stringNumber, bitSize)
	return float, err
}
