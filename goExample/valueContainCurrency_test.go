package goExample

import (
	"regexp"
	"strings"
)

func isCurrencyFormat(s string, currencySymbols []string) bool {
	// 将货币符号数组转换为一个字符串，用“|”分隔
	currencies := strings.Join(currencySymbols, "|")

	// 构建正则表达式
	reg := regexp.MustCompile(`^((` + currencies + `)[0-9]*\.?[0-9]+|[0-9]*\.?[0-9]+(` + currencies + `))$`)

	// 使用正则表达式匹配字符串
	return reg.MatchString(s)
}

//func main() {
//	currencySymbols := []string{"$", "€", "£", "¥"}
//	strings := []string{"$1.2", "20$", "3", "$", "1.2", "20", "3.0$", "abc", "€3.4", "100¥", "5£"}
//
//	for _, s := range strings {
//		fmt.Printf("Does \"%s\" match the format? %v\n", s, isCurrencyFormat(s, currencySymbols))
//	}
//}
