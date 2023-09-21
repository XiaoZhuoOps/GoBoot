package goExample

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = "[$2e7]"

	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println(s.TokenText())
	}
}
