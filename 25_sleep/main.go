package main

import (
	"fmt"
	"time"
)

func mySleepTimer(dur time.Duration) {
	// устанавливаем таймер
	timer := time.NewTimer(dur)

	// ждем пока чето прилетит в канал
	<-timer.C
}

func mySleepAfter(dur time.Duration) {
	// ждем
	<-time.After(dur)
}

func main() {
	fmt.Printf("time before sleep: %v\n", time.Now())

	mySleepTimer(time.Second)

	fmt.Printf("time after second mySleepTimer: %v\n", time.Now())

	mySleepAfter(time.Second)

	fmt.Printf("time after second mySleepAfter: %v\n", time.Now())
}
