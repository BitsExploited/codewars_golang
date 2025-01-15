package kata

import (
	"strings"
	"unicode"
)

func ToJadenCase(str string) string {
	list := strings.Split(str, " ")

	for i, j := range list {
		if len(j) > 0 && unicode.IsLower(rune(j[0])) {
			list[i] = strings.ToUpper(string(j[0])) + j[1:]
		}
	}

	return strings.Join(list, " ")
}
