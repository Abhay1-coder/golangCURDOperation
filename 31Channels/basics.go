package main

import (
	"fmt"
)

func basicChannel() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Send a value into the channel
	}()

	val := <-ch // Receive the value from the channel
	fmt.Println(val)
}

func main() {
	basicChannel()
}
