package goExample

import (
	"fmt"
	"testing"
)

func TestValueInspector(t *testing.T) {
	//payload := `{"value":["123"]}`
	payload := `{'1214312das`
	for _, str := range payload {
		fmt.Printf("%v ", string(str))
	}
}
