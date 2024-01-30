package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int, resultCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		resultCh <- value * 2
	}
}

func main() {
	dataCh := make(chan int)
	resultCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(dataCh, &wg)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go consumer(dataCh, resultCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		fmt.Println(result)
	}
}
