package main

import (
	"fmt"
	"sync"
	"time"
)

func Spam(poolSize int) {
	// входной канал
	inCh := make(chan string)
	//WaitGroup
	wg := &sync.WaitGroup{}

	// пул воркеров
	for i := 0; i < poolSize; i++ {
		// дабавляем в вейт группу
		wg.Add(1)
		go func(in <-chan string, wg *sync.WaitGroup) {
			// вычитаем из вейт группы
			defer wg.Done()
			for text := range in {
				fmt.Println(text)
			}
		}(inCh, wg)
	}

	// горутина дающая текст воркерам
	go func(ch chan<- string) {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%v", time.Now())
			time.Sleep(time.Second)
		}
		close(ch)
	}(inCh)

	wg.Wait()
	// ждем пока читающие канал горутины закончат свое существование, а они
	// работают пока не закроется канал, а его закроет писатель, когда напишет
	// все что ему надо
}

func main() {
	var n int
	fmt.Println("pool size:")
	fmt.Scanln(&n)
	Spam(n)
}
