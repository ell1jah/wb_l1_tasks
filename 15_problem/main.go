package main

import (
	"fmt"
)

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// проблемы:
// такое использование глобальной переменной небезопасно
// слайс может отработать некорректно, тк работает с байтами, а символ может быть больше байта
// нет смысла сначала делать строку на 1024 длины, а потом обрезать ее до 100 байт

func createHugeString(len int) string {
	runes := make([]rune, len)

	// заполняем массив рун(символов)
	for i := range runes {
		runes[i] = 'a'
	}

	// конвертируем в строку
	return string(runes)
}

func someFunc() string {
	return createHugeString(50)
}

func main() {
	fmt.Println(someFunc())
}
