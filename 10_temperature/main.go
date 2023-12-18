package main

import "fmt"

func GroupTemp(temps []float64, step int) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temps {
		intTemp := int(temp)
		// отбрасываем остаток от деления на step(10)
		key := intTemp - (intTemp % step)
		// записываем результат
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(GroupTemp(temps, 10))
}
