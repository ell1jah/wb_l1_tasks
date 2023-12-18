package main

import "fmt"

func Quicksort(arr []int) {
	quicksortRec(arr, 0, len(arr)-1)
}

func quicksortRec(arr []int, low, high int) {
	// условие выхода, если отрезок состоит из 1 элемента
	if low < high {
		// определяем опорный элемент и ставим числа меньшие или равные ему слева
		// а большие справа
		partitionIndex := partition(arr, low, high)

		// рекурсивно вызываем сортировку для левой и правой части массива относительно опорного эл.
		quicksortRec(arr, low, partitionIndex-1)
		quicksortRec(arr, partitionIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// определяем опорного крайним правым значением
	pivot := arr[high]
	// правая граница левой части
	i := low - 1

	for j := low; j < high; j++ {
		// если эл меньше или равен опорному, сдвигаем правую границу левой части и ставим туда этот элемент
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// ставим опорный элемент на границе правой и левой части
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Unsorted:", arr)

	Quicksort(arr)
	fmt.Println("Sorted:", arr)
}
