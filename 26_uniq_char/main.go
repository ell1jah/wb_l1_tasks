package main

import (
	"fmt"
	"unicode"
)

func isCharsUniq(str string) bool {
	// мапа учета символов
	chars := make(map[rune]bool, len(str))

	for _, char := range str {
		// в нижний регистр
		lower := unicode.ToLower(char)
		// чекаем что такого символа еще не было
		if chars[lower] {
			return false
		}

		// добавляем символ в мапу
		chars[lower] = true
	}

	return true
}

func main() {
	fmt.Printf("string: %s, isCharsUniq: %t\n", "abcd", isCharsUniq("abcd"))
	fmt.Printf("string: %s, isCharsUniq: %t\n", "abCdefAaf", isCharsUniq("abCdefAaf"))
	fmt.Printf("string: %s, isCharsUniq: %t\n", "aabcd", isCharsUniq("aabcd"))
}
