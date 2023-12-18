package main

import (
	"fmt"
	"strings"
)

func ReverseWords(input string) string {
	// делим строку на слова
	words := strings.Fields(input)

	// переставляем слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// соединяем обратно в одну строку
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"

	fmt.Println("src str: ", str)
	fmt.Println("result str: ", ReverseWords(str))
}
