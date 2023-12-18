package main

import (
	"context"
	"fmt"
	"time"
)

func cancelChan(cancelCh chan bool) {
	for {
		select {
		// проверка что время еще не истекло
		case <-cancelCh:
			return
		default:
			fmt.Printf("cancelChan goroutine: %v\n", time.Now())
			time.Sleep(time.Second)
		}
	}
}

func closeChan(closeCh chan bool) {
	for {
		select {
		// проверка что время еще не истекло
		case <-closeCh:
			return
		default:
			fmt.Printf("closeChan goroutine: %v\n", time.Now())
			time.Sleep(time.Second)
		}
	}
}

func contextCancel(ctx context.Context) {
	for {
		select {
		// проверка что время еще не истекло
		case <-ctx.Done():
			return
		default:
			fmt.Printf("contextCancel goroutine:  %v\n", time.Now())
			time.Sleep(time.Second)
		}
	}
}

func contextTimeout(ctx context.Context) {
	for {
		select {
		// проверка что время еще не истекло
		case <-ctx.Done():
			return
		default:
			fmt.Printf("contextTimeout goroutine:  %v\n", time.Now())
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var n int
	fmt.Println("seconds count:")
	fmt.Scanln(&n)

	// 1 способ: кинуть значение в канал
	cancelCh := make(chan bool)
	go cancelChan(cancelCh)

	// 2 способ: закрыть канал
	closeCh := make(chan bool)
	go closeChan(closeCh)

	// 3 способ: контекст с функцией завершения
	ctxCancel, ctxCancelFunc := context.WithCancel(context.Background())
	go contextCancel(ctxCancel)

	// 4 способ: контекст с таймаутом
	ctxTimeout, _ := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	go contextTimeout(ctxTimeout)

	// ждем
	time.Sleep(time.Duration(n) * time.Second)

	// завершаем горутины
	cancelCh <- true
	close(closeCh)
	ctxCancelFunc()

	// чтоб программа сразу не завершалась и можно было убедится что горутины остановились
	fmt.Scanln()
}
