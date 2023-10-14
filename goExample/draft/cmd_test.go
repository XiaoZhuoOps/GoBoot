package draft

import (
	"fmt"
	"testing"
)

type Cmd interface {
	Execute()
}

type GetCmd struct {
	Key string
	Val string
	Err error
}

func (g *GetCmd) Execute() {
	// external
}

func TestDivideByZero(t *testing.T) {
	var exchangeRate float64
	fmt.Println(0 == exchangeRate)
	fmt.Println(1 / exchangeRate)
}
