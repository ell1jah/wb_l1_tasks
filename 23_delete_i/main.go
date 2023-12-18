package main

import "fmt"

func deleteI(arr []int, i int) []int {
	// удаляем i-ый
	return append(arr[:i], arr[i+1:]...)
}

// Если не важен порядок элементов
func deleteILazy(arr []int, i int) []int {
	// ставим на место i-го последний и уменьшаем размер
	arr[i] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]

	return arr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	// создал 2 одинаковых массива для теста

	fmt.Printf("origin: %v\n", arr1)
	fmt.Printf("deleteI i=3: %v\n", deleteI(arr1, 3))
	fmt.Printf("deleteILazy i=3: %v\n", deleteILazy(arr2, 3))
}
