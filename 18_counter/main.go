package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// можно было бы сделать через мьютексы, но это бы работало гораздо медленее
// плюс в 7 задании я реализовывал с помощью мьютексов, тут было бы тоже самое

type counter struct {
	// атомарный инт
	atomic.Int64
}

func NewCounter() *counter {
	return &counter{}
}

func (c *counter) Inc() {
	c.Add(1)
}

func (c *counter) Get() int64 {
	return c.Load()
}

// не относится к реализации, просто для теста
func incTest(cnt *counter, wg *sync.WaitGroup, iterCnt int) {
	defer wg.Done()

	for i := 0; i < iterCnt; i++ {
		cnt.Inc()
	}
}

func main() {
	cnt := NewCounter()
	wg := &sync.WaitGroup{}

	// запускаем несколько горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incTest(cnt, wg, 1000)
	}

	// ждем горутин
	wg.Wait()

	// шаг 10000 чтоб не захламалять консоль, смотрим что получилось
	fmt.Println("result value:", cnt.Get())
}
