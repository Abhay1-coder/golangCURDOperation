package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func(a int, b int) {
		time.Sleep(1 * time.Second)
		c1 <- a + b

	}(4, 8)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout-1")
	}

	go func(a int, b int) {
		time.Sleep(4 * time.Second)
		c2 <- a * b
	}(4, 8)

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout-2")
	}
}
