package goExample

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"testing"
)

var testSet = set.NewSetFromSlice([]interface{}{})

type typeA struct {
	a int
}

func TestSet(t *testing.T) {
	// type related
	var ele1 int8 = 0
	testSet.Add(ele1)

	var ele2 int32 = 0
	fmt.Println(testSet.Contains(ele2))

	// deep equals
	t1 := &typeA{a: 1}
	t2 := &typeA{a: 1}
	testSet.Add(t1)
	fmt.Println(testSet.Contains(t2))
}
