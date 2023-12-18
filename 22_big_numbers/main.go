package main

import (
	"fmt"
	"math/big"
)

func main() {
	// результат
	res := big.NewInt(0)

	// создаем числа
	num1, _ := big.NewInt(0).SetString("239480284023840930893229317519381095713474", 10)
	num2, _ := big.NewInt(0).SetString("19379382748397192759814758145138949789134", 10)

	fmt.Printf("sum: %v\n", res.Add(num1, num2))
	fmt.Printf("sub: %v\n", res.Sub(num1, num2))
	fmt.Printf("mul: %v\n", res.Mul(num1, num2))
	fmt.Printf("div: %v\n", res.Div(num1, num2))
}
