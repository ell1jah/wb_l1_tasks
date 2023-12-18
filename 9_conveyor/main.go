package main

import (
	"fmt"
	"sync"
)

// часть конвеиера возводящая в квадрат
func SquareConveyor(inCh <-chan int, outCh chan<- int) {
	// сообщаем что мы закончили работу
	defer close(outCh)

	// читаем пока не закроют канал
	for num := range inCh {
		outCh <- num * num
	}
}

// часть конвеиера печатающая результат в stdout
func PrintConveyor(inCh <-chan int, wg *sync.WaitGroup) {
	// сообщаем что мы закончили работу
	defer wg.Done()

	// читаем пока не закроют канал
	for num := range inCh {
		fmt.Println(num)
	}
}

// часть конвеиера подающая числа на вход
func InputConveyor(arr []int, outCh chan<- int) {
	// сообщаем что мы закончили работу
	defer close(outCh)

	for _, num := range arr {
		outCh <- num
	}
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	wg := &sync.WaitGroup{}

	arr := []int{1, 2, 3, 4, 5}

	wg.Add(1)
	go InputConveyor(arr, ch1)
	go SquareConveyor(ch1, ch2)
	go PrintConveyor(ch2, wg)

	// ждем завершения PrintConveyor
	wg.Wait()
}
