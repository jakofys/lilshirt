package errors

import (
	"strings"
	"unicode"
)

func tosnakecase(init string) string {
	str := ""
	init = strings.TrimSpace(init)
	for i, letter := range init {
		switch {
		case unicode.IsSpace(letter), unicode.IsSymbol(letter), unicode.IsMark(letter), unicode.IsPunct(letter):
			str += "_"
		case unicode.IsUpper(letter) && i > 0:
			if unicode.IsUpper(rune(init[i-1])) && unicode.IsUpper(rune(init[i+1])) {
				str += string(unicode.ToLower(letter))
			} else {
				str += "_" + string(unicode.ToLower(letter))
			}
		default:
			str += string(unicode.ToLower(letter))
		}
	}
	return str
}
