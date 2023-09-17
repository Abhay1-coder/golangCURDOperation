package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	select {
	case ch <- 42:
		msg := <-ch
		fmt.Println("Sent data into channel", msg)
	default:
		fmt.Println("Channel operation would block, so default case executed")
	}
}
