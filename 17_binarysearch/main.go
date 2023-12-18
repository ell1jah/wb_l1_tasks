package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	// левая и правая границы поиска
	low, high := 0, len(arr)-1

	// пока границы не сойдутся
	for low <= high {
		// берем серединный элемент
		mid := (low + high) / 2
		midValue := arr[mid]

		// если серединный элемент равен цели, возвращаем индекс
		// если серединный элемент больше цели, смещаем правую границу для поиска в правой части
		// если серединный элемент меньше цели, смещаем правую границу для поиска в левой части
		switch {
		case midValue == target:
			return mid
		case midValue < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	// если цель не найдена
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	index := binarySearch(arr, target)

	fmt.Println("array: ", arr)
	fmt.Println("target: ", target)
	fmt.Println("index: ", index)
}
