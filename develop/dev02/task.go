package dev

import (
	"fmt"
	"unicode"
)

func unpackString(s string) (string, error) {
	var result string

	for i := 0; i < len(s); i++ {
		char := rune(s[i])

		if unicode.IsDigit(char) {
			return "", fmt.Errorf("некорректная строка")
		}

		result += string(char)

		if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			count := int(rune(s[i+1]) - '0')
			result += repeatChar(char, count-1)
			i++
		}
	}

	return result, nil
}

func repeatChar(char rune, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += string(char)
	}
	return result
}
