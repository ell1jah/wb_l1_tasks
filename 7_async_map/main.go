package main

import (
	"fmt"
	"sync"
)

type asyncMap struct {
	body map[int]int
	mu   *sync.RWMutex
}

func NewAsyncMap() *asyncMap {
	return &asyncMap{
		body: make(map[int]int),
		mu:   &sync.RWMutex{},
	}
}

func (am *asyncMap) Set(key, value int) {
	// монопольный доступ
	am.mu.Lock()
	defer am.mu.Unlock()

	am.body[key] = value
}

func (am *asyncMap) Get(key int) (value int, ok bool) {
	// запрещаем писать, читать можно
	am.mu.RLock()
	defer am.mu.RUnlock()

	value, ok = am.body[key]
	return
}

// не относится к реализации, просто для теста
func setMap(am *asyncMap, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		am.Set(i, i)
	}
}

func main() {
	am := NewAsyncMap()
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go setMap(am, wg)
	}

	// ждем горутин
	wg.Wait()

	// шаг 10000 чтоб не захламалять консоль, смотрим что получилось
	for i := 0; i < 10000; i += 1000 {
		fmt.Println(am.Get(i))
	}
}
