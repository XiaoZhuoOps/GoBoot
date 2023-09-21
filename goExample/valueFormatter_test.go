package goExample

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
)

import (
	set "github.com/deckarep/golang-set"
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

func TestValueUnformat(t *testing.T) {
	fmt.Println(accounting.UnformatNumber("123,456,789.21", 2, "USD"))
	fmt.Println(accounting.UnformatNumber("3,19 €", 2, "USD"))
	fmt.Println(accounting.UnformatNumber("3,19 €", 2, "EUR"))
}

func TestExtractCurrencySymbol(t *testing.T) {
	s := set.NewSet()
	for _, v := range accounting.LocaleInfo {
		s.Add(strings.TrimSpace(v.ComSymbol))
	}
	fmt.Println(s)
}

func TestValueFormat(t *testing.T) {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))                        // "$123,456,789.21"
	fmt.Println(ac.FormatMoney(12345678))                                // "$12,345,678.00"
	fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))                 // "$25,925,925.67"
	fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3)))                // "-$25,925,925.67"
	fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123)))  // "$123,456,789.21"
	fmt.Println(ac.FormatMoneyDecimal(decimal.New(123456789213123, -6))) // "$123,456,789.21"

	ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	fmt.Println(ac.FormatMoney(4999.99)) // "€4.999,99"

	// Or retrieve currency info from Locale struct
	lc := accounting.LocaleInfo["USD"]
	ac = accounting.Accounting{Symbol: lc.ComSymbol, Precision: 2, Thousand: lc.ThouSep, Decimal: lc.DecSep}
	fmt.Println(ac.FormatMoney(500000)) // "$500,000.00"
	fmt.Println(accounting.UnformatNumber("123,456,789.21", 2, "USD"))

	ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
	fmt.Println(ac.FormatMoney(500000)) // "£ 500,000"

	ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	fmt.Println(ac.FormatMoney(1000000)) // "GBP 1,000,000"
	fmt.Println(ac.FormatMoney(-5000))   // "GBP (5,000)"
	fmt.Println(ac.FormatMoney(0))       // "GBP --"

	ac = *accounting.DefaultAccounting("GBP", 2)
	fmt.Println(ac.FormatMoney(1000000)) // "GBP 1,000,000"
	fmt.Println(ac.FormatMoney(-5000))   // "GBP (5,000)"
	fmt.Println(ac.FormatMoney(0))       // "GBP --"

	ac = *accounting.NewAccounting("GBP", 2, ",", ".", "%s %v", "%s (%v)", "%s --")
	fmt.Println(ac.FormatMoney(1000000)) // "GBP 1,000,000"
	fmt.Println(ac.FormatMoney(-5000))   // "GBP (5,000)"
	fmt.Println(ac.FormatMoney(0))       // "GBP --"
}
