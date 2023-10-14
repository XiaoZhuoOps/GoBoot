package load

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"testing"
)

type CurrencySymbols struct {
	symbols string
}

var (
	instance             CurrencySymbols
	once                 sync.Once
	CurrencySymbolToCode = map[string][]string{
		"$":   {"USD", "USD2"},
		"¥":   {"CNY"},
		"ASD": {"ASD"},
	}
)

func GetIns() CurrencySymbols {
	once.Do(func() {
		var s []string
		for k := range CurrencySymbolToCode {
			s = append(s, regexp.QuoteMeta(k))
		}
		pattern := strings.Join(s, "|")
		instance = CurrencySymbols{symbols: pattern}
	})
	return instance
}

func TestGetIns(t *testing.T) {
	for i := 0; i < 100; i++ {
		go fmt.Println(GetIns().symbols)
	}
}

func TestQuoteMeta(t *testing.T) {
	ss := []string{"$", "$abc", ".\""}
	for _, s := range ss {
		fmt.Println(regexp.QuoteMeta(s))
	}
}
