package main

import (
	"fmt"
)

const PoolSize = 8

// сумма квадратов, результат в канал
func squaresSumPart(arr []int, res chan<- int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i] * arr[i]
	}

	res <- sum
}

func SquaresSum(arr []int) int {
	//выходной канал
	ch := make(chan int)

	//размер отдельновычисляемого куска
	partLen := len(arr) / PoolSize
	if len(arr)%PoolSize != 0 {
		partLen++
	}

	// количество кусков
	partCnt := PoolSize

	// запуск горутин считающих кусок от массива
	for i := 0; i < PoolSize; i++ {
		// границы куска
		begin := partLen * i
		end := partLen * (i + 1)
		if begin >= len(arr) {
			partCnt = i
			break
		}
		if end > len(arr) {
			end = len(arr)
		}

		go squaresSumPart(arr[begin:end], ch)
	}

	// сбор результатов
	sum := 0
	for i := 0; i < partCnt; i++ {
		sum += <-ch
	}

	return sum
}

func main() {
	src := make([]int, 5)
	for i := 0; i < len(src); i++ {
		src[i] = i * 2
	}

	fmt.Println("source:", src)
	fmt.Println("squares sum:", SquaresSum(src))
}
