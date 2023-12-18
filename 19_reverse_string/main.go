package main

import "fmt"

func Reverse(str string) string {
	// кастим к слайсу рун для корректной работы
	srcRunes := []rune(str)
	// результат
	resRunes := make([]rune, len(srcRunes))

	// записываем в обратном порядке
	for i, c := range srcRunes {
		resRunes[len(resRunes)-i-1] = c
	}

	// кастим к строке
	return string(resRunes)
}

func main() {
	str := "главрыба"

	fmt.Println("src str: ", str)
	fmt.Println("result str: ", Reverse(str))
}
