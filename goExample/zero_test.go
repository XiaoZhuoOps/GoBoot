package goExample

import (
	"code.byted.org/gopkg/mockito"
	"strconv"
	"testing"
)

func TestZero(t *testing.T) {
	mockito.PatchConvey("", t, func() {
		s, err := strconv.ParseFloat("0.0000000000", 64)
		t.Log(s == 0.0, err)
		s, err = strconv.ParseFloat("0.0000000001", 64)
		t.Log(s == 0.0, err)
		s, err = strconv.ParseFloat("0.00", 64)
		t.Log(s == 0.0, err)
	})
}
