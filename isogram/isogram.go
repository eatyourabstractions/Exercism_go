package isogram

import (
	"unicode"
)


func IsIsogram(phrase string) bool {
	runes := make(map[rune]bool)
	for _, char := range phrase {
		char = unicode.ToLower(char)
		if unicode.IsLetter(char) && runes[char] {
			return false
		}
		runes[char] = true
	}
	return true
}