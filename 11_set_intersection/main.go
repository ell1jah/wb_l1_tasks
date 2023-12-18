package main

import "fmt"

func Intersect(set1, set2 map[int]bool) map[int]bool {
	if len(set2) < len(set1) {
		// меняем местами чтоб итерироваться по мапе меньшей длины
		set1, set2 = set2, set1
	}

	result := make(map[int]bool, len(set1))

	for key := range set1 {
		// если значение в обоих множествах, добавляем в результат
		if set2[key] {
			result[key] = true
		}
	}

	return result
}

func main() {
	set1 := map[int]bool{1: true, 2: true, 3: true, 8: true, 10: true, 11: true, 14: true}
	set2 := map[int]bool{2: true, 4: true, 5: true, 8: true, 11: true, 13: true}

	fmt.Println("set1: ", set1)
	fmt.Println("set2: ", set2)
	fmt.Println("intersection: ", Intersect(set1, set2))
}
