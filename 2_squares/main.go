package main

import (
	"fmt"
)

const PoolSize = 8

type Task struct {
	ID  int
	Num int
}

type Response struct {
	ID     int
	Result int
}

func Squares(arr []int) []int {
	// входной и выходной каналы
	inCh := make(chan Task)
	outCh := make(chan Response)

	// пул воркеров
	for i := 0; i < PoolSize; i++ {
		go func(in <-chan Task, out chan<- Response) {
			for task := range in {
				result := task.Num * task.Num
				out <- Response{task.ID, result}
			}
		}(inCh, outCh)
	}

	// горутина отправляющая задачи воркерам
	go func(ch chan<- Task, arr []int) {
		for i := 0; i < len(arr); i++ {
			ch <- Task{i, arr[i]}
		}
		close(ch)
	}(inCh, arr)

	squares := make([]int, len(arr))

	// сбор результатов
	for i := 0; i < len(arr); i++ {
		res := <-outCh
		squares[res.ID] = res.Result
	}

	return squares
}

func main() {
	src := make([]int, 1)
	for i := 0; i < len(src); i++ {
		src[i] = i * 2
	}

	fmt.Println("source:", src)
	fmt.Println("squares:", Squares(src))
}
