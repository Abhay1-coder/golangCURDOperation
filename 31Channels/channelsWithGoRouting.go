package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := id * 2
	ch <- result
}

func channelWithGoroutines() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch) // Close the channel when all workers are done
	}()

	for result := range ch {
		fmt.Println(result)
	}
}

func main() {
	channelWithGoroutines()
}
