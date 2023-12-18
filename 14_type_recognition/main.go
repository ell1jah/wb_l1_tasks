package main

import (
	"fmt"
	"reflect"
)

func Recognize(something interface{}) string {
	// получаем тип переменной с помощью рефлекта
	sType := reflect.ValueOf(something).Type()

	// проверяем "вид типа" переменной
	switch sType.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	}

	return "unknown"
}

func main() {
	str := "string"
	i := 1
	b := false
	ch := make(chan int)

	fmt.Printf("string is: %s\n", Recognize(str))
	fmt.Printf("int is: %s\n", Recognize(i))
	fmt.Printf("bool is: %s\n", Recognize(b))
	fmt.Printf("chan is: %s\n", Recognize(ch))
}
