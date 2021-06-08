package library

import (
	"strings"
	"unicode"
)

func ChangeCase(str string) string {
	letters := []rune(str)
	if unicode.IsLower(letters[0]) {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}
}
