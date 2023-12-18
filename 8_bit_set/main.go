package main

import "fmt"

func SetBit(num int64, i int, toZero bool) int64 {
	setter := int64(1)
	// побитовый сдвиг на нужный бит
	setter <<= i

	var result int64
	if toZero {
		// побитовое И НЕ для установки i-ого бита в ноль
		result = num &^ setter
	} else {
		// побитовое ИЛИ для установки i-ого бита в единицу
		result = num | setter
	}

	return result
}

func main() {
	var num int64
	fmt.Println("num:")
	fmt.Scanln(&num)

	var i int
	fmt.Println("i:")
	fmt.Scanln(&i)
	if i > 63 || i <= 0 {
		fmt.Println("incorrect")
		return
	}

	// если true, сетим бит в ноль, если false, то в единицу
	var toZero bool
	fmt.Println("toZero:")
	fmt.Scanln(&toZero)

	fmt.Println(SetBit(num, i, toZero))
}
