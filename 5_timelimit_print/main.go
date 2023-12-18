package main

import (
	"fmt"
	"time"
)

func Spam(secCnt int) {
	// входной канал
	inCh := make(chan string)

	// горутина читающая канал и пишущая в stdout
	go func(ch <-chan string) {
		for text := range ch {
			fmt.Println(text)
		}
	}(inCh)

	// таймер для остановки через время
	timer := time.NewTimer(time.Duration(secCnt) * time.Second)
	for {
		select {
		// проверка что время еще не истекло
		case <-timer.C:
			// закрываем канал для остановки читателя
			close(inCh)
			return
		default:
			inCh <- fmt.Sprintf("%v", time.Now())
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var n int
	fmt.Println("seconds count:")
	fmt.Scanln(&n)
	Spam(n)
}
