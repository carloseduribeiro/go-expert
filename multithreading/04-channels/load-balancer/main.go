package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, data chan int, wg *sync.WaitGroup) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	data := make(chan int)
	workers := 50
	wg := &sync.WaitGroup{}
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, data, wg)
	}
	for i := 0; i < 100; i++ {
		data <- i
	}
	wg.Done()
}
