package main

import (
	"fmt"
	"log"
)

import (
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

// const regexp_currency_symbol = `[¥$€£₹₽₣₤₱₩₪₫₴₵₸₹₺₼₽₾]`
const regexp_currency_symbol = `\$|¥|€|£|₹|R\$|C\$|A\$|₩|₽|Rp|₺|SR`
const regexp_scientific_notation = `[1-9](\.[0-9]+)?[Ee][-+]?[0-9]+`
const regexp_potential_number = `[0-9,.]+`

const regexp_float = "^\\d+(\\.\\d+)?$"
const regexp_int = "^\\d+$"

var Literals []string       // The tokens representing literal strings
var Keywords []string       // The keyword tokens
var Tokens []string         // All of the tokens (including literals and keywords)
var TokenIds map[string]int // A map from the token names to their int ids
var Lexer *lexmachine.Lexer // The lexer object. Use this to construct a Scanner

// Called at package initialization. Creates the lexer and populates token lists.
func init() {
	initTokens()
	var err error
	Lexer, err = initLexer()
	if err != nil {
		panic(err)
	}
}

func initTokens() {
	Literals = []string{
		"LeftBracket",  // [
		"RightBracket", // ]
		"Quotation",    // "
	}
	Keywords = []string{
		"ScientificNotation", // 2.3e7
		"PotentialNumber",    // 123 1,23 1.12 1,1,1
	}
	Tokens = []string{
		"CurrencySymbol",      // $
		"OtherNonNumericChar", // others, contain { E/e
	}
	Tokens = append(Tokens, Keywords...)
	Tokens = append(Tokens, Literals...)
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

// Creates the lexer object and compiles the NFA.
func initLexer() (*lexmachine.Lexer, error) {
	lexer := lexmachine.NewLexer()

	lexer.Add([]byte(`\[`), token("LeftBracket"))
	lexer.Add([]byte(`\]`), token("RightBracket"))
	lexer.Add([]byte(`\"`), token("Quotation"))
	lexer.Add([]byte(regexp_scientific_notation), token("ScientificNotation"))
	lexer.Add([]byte(regexp_potential_number), token("PotentialNumber"))
	lexer.Add([]byte(regexp_currency_symbol), token("CurrencySymbol"))
	lexer.Add([]byte(`.`), token("OtherNonNumericChar"))

	//for _, lit := range Literals {
	//	r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
	//	lexer.Add([]byte(r), token(lit))
	//}
	//for _, name := range Keywords {
	//	lexer.Add([]byte(strings.ToLower(name)), token(name))
	//}
	//
	//lexer.Add([]byte(`//[^\n]*\n?`), token("COMMENT"))
	//lexer.Add([]byte(`/\*([^*]|\r|\n|(\*+([^*/]|\r|\n)))*\*+/`), token("COMMENT"))
	//lexer.Add([]byte(`([a-z]|[A-Z]|[0-9]|_)+`), token("ID"))
	//lexer.Add([]byte(`[0-9]*\.[0-9]+`), token("ID"))
	//lexer.Add([]byte(`"([^\\"]|(\\.))*"`),
	//	func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
	//		x, _ := token("ID")(scan, match)
	//		t := x.(*lexmachine.Token)
	//		v := t.Value.(string)
	//		t.Value = v[1 : len(v)-1]
	//		return t, nil
	//	})
	//lexer.Add([]byte("( |\t|\n|\r)+"), skip)
	//lexer.Add([]byte(`\<`),
	//	func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
	//		str := make([]byte, 0, 10)
	//		str = append(str, match.Bytes...)
	//		brackets := 1
	//		match.EndLine = match.StartLine
	//		match.EndColumn = match.StartColumn
	//		for tc := scan.TC; tc < len(scan.Text); tc++ {
	//			str = append(str, scan.Text[tc])
	//			match.EndColumn += 1
	//			if scan.Text[tc] == '\n' {
	//				match.EndLine += 1
	//			}
	//			if scan.Text[tc] == '<' {
	//				brackets += 1
	//			} else if scan.Text[tc] == '>' {
	//				brackets -= 1
	//			}
	//			if brackets == 0 {
	//				match.TC = scan.TC
	//				scan.TC = tc + 1
	//				match.Bytes = str
	//				x, _ := token("ID")(scan, match)
	//				t := x.(*lexmachine.Token)
	//				v := t.Value.(string)
	//				t.Value = v[1 : len(v)-1]
	//				return t, nil
	//			}
	//		}
	//		return nil,
	//			fmt.Errorf("unclosed HTML literal starting at %d, (%d, %d)",
	//				match.TC, match.StartLine, match.StartColumn)
	//	},
	//)

	err := lexer.Compile()
	if err != nil {
		return nil, err
	}
	return lexer, nil
}

// a lexmachine.Action function which skips the match.
func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// a lexmachine.Action function with constructs a Token of the given token type by
// the token type's name.
func token(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}

func main() {
	//test("N/A")
	//test("[Object object]")
	//test("$1,99")
	//test("1.99 $")
	//test("2.3e-7")
	//test("undefined[1].price")
	//test("[1]")
	//test("[\"1,99\"]")
	//
	//test("20e7")
	//test("[$27]")
	//test("[\"$27\"]")
	//test("$1,99.909")
	//test("$+1,99.909")
	//test("$+1,99.909")
	//test("1,99.909")
	//test("1e99e909")
	//test("9,99 R$")
}

func test(value string) {
	s, err := Lexer.Scanner([]byte(value))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("======================", value, "======================")
	fmt.Println("Type    | Lexeme     | Position")
	fmt.Println("--------+------------+------------")
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if _, is := err.(*machines.UnconsumedInput); is {
			// to skip bad token do:
			// s.TC = ui.FailTC
			log.Fatal(err) // however, we will just fail the program
		} else if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lexmachine.Token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n",
			Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
}
