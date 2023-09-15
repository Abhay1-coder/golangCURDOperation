package main

import (
	"fmt"
)

func bufferedChannel() {
	ch := make(chan int, 2) // Create a buffered channel with a capacity of 2

	ch <- 1
	ch <- 2
	// ch <- 3 // Uncommenting this line would result in a deadlock

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	bufferedChannel()
}
