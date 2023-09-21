package goExample

import set "github.com/deckarep/golang-set"

// valueInspector
// all tokens:
// [ ]
// "
// $
// ,
// .
// Ee
// Digit
// NonNumericChar
type token struct {
	char  rune
	index int
}

type valueInspector struct {
	labels             set.Set
	startIndex         int
	endIndex           int
	nConsecutiveDigits int
	nElement           int
	lastSep            rune
	LeftBrackets       []token
	RightBrackets      []token
	Quotes             []token
	CurrencySymbols    []token
	Commas             []token
	Dots               []token
	Es                 []token // only first one, others belong to NonNumericChars
	Digits             []token
	NonNumericChars    []token
}

func (r *valueInspector) Parse() (labels []string, res string) {
	var ok bool
	if ok, res = r.parseIntOrFloat(); ok {
		return labels, res
	}
	for r.parseSingletonArray() || r.parseQuotedValue() {
	}
	if ok, res = r.parseContainCurrency(r.startIndex, r.endIndex); ok {
		return labels, res
	}
	if ok, res = r.parseScientificNotation(r.startIndex, r.endIndex); ok {
		return labels, res
	}
	return nil, ""
}

func (r *valueInspector) parseIntOrFloat() (bool, string) {
	return false, ""
}

// true
// [ 1 . 2 ]
// [ " 1 , 2 " ]
// [ " 1 2 3 " ]
// false
// [ 1 , 2 ]
// [ " 1 , 2 " , " 1 , 2 " ]
func (r *valueInspector) parseSingletonArray() bool {
	// only one leftBracket && only one rightBracket &&
	// index of the leftBracket == startIndex &&
	// index of the rightBracket == endIndex &&
	// only one element
	// startIndex ++
	// endIndex --
	// addLabel if true
	return false
}

// true
// " 1 , 2 3 "
// " [ " 1 2 3 " ] "
func (r *valueInspector) parseQuotedValue() bool {
	// index of the first quote ==  startIndex &&
	// index of the last quote ==  endIndex
	// startIndex ++
	// endIndex --
	// addLabel if true
	return false
}

// true
// $ 2 , 7 7
// 2 $
// $ 2 e 7

// false
// $ [ 1 ]
// $ " 1 "
func (r *valueInspector) parseContainCurrency(left int, right int) (bool, string) {
	// only one currencySymbol &&
	// index of the currencySymbol == startIndex, startIndex++
	// index of the currencySymbol == endIndex, endIndex--
	// if parseIntOrFloat() || parseScientificNotation() || parseNotStdFormatValue()
	// addLabel if true
	return false, "0"
}

func (r *valueInspector) parseScientificNotation(left int, right int) (bool, string) {

	return false, "0"
}

//func (r *valueInspector) scan(value string) *valueInspector {
//	// TODO remove all space
//	for _, c := range value {
//		switch c {
//		case ',':
//			r.nCommas++
//			r.lastSep = c
//			r.nConsecutiveDigits = 0
//		case '.':
//			r.nDots++
//			r.lastSep = c
//			r.nConsecutiveDigits = 0
//		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
//			r.nConsecutiveDigits++
//			r.nDigits++
//		default:
//			// NOTE: space( ), negative sign(-) are NOT considered numeric at the moment.
//			r.nNonNumericChars++
//		}
//	}
//	return r
//}
//
//func (r *valueInspector) nLastSep() int {
//	if r.lastSep == '.' {
//		return r.nDots
//	} else if r.lastSep == ',' {
//		return r.nCommas
//	} else {
//		return 0
//	}
//}
//
//func (r *valueInspector) theOtherSep() rune {
//	if r.lastSep == '.' && r.nCommas > 0 {
//		return ','
//	} else if r.lastSep == ',' && r.nDots > 0 {
//		return '.'
//	} else {
//		return 0
//	}
//}
