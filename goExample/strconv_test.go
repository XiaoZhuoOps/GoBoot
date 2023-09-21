package goExample

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

const regexpCurrencySymbol = `\$|¥|€|£|₹|R\$|C\$|A\$|₩|₽|Rp|₺|SR`
const regexpScientificNotation = `[1-9](\.[0-9]+)?[Ee][-+]?[0-9]+`
const regexpPotentialNumber = `[0-9]+([.,][0-9]+)*`

func TestStrconv(t *testing.T) {
	var a uint64 = 65
	b := 1
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(strconv.FormatUint(a, 10))

	fmt.Println("regexpExactlyCurrencySymbol   ", strconv.Quote(`^(\$|¥|€|£|₹|R\$|C\$|A\$|₩|₽|Rp|₺|SR)$`))
	fmt.Println("regexpStdNumber   ", strconv.Quote("^\\d+(\\.\\d+)?$"))

	fmt.Println("regexpCurrencySymbol ", strconv.Quote(regexpCurrencySymbol))
	fmt.Println("regexpScientificNotation ", strconv.Quote(regexpScientificNotation))
	fmt.Println("regexpPotentialNumber ", strconv.Quote(regexpPotentialNumber))
}
