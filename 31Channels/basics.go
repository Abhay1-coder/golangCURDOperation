/*
*
In Go, channels are a powerful and fundamental feature used for communication and synchronization between

	goroutines (concurrent threads of execution). Channels provide a way for goroutines to send and receive data
	 safely and efficiently. They enable you to coordinate the execution of different parts of your program and
	 share data between them. Here's a basic explanation of channels:

Communication: Channels allow goroutines to communicate with each other by sending and receiving data.

	Think of a channel as a conduit or pipe through which data can flow between goroutines.

Synchronization: Channels can be used to synchronize the execution of goroutines. They ensure that one goroutine

	doesn't proceed until another goroutine is ready or has completed its task.

Concurrency: Channels are a key tool for concurrent programming in Go. They enable multiple goroutines to work

	together without conflicts or race conditions, making it easier to write concurrent code.

Blocking: When a goroutine sends data on a channel, it will block (pause) until another goroutine is ready to receive

	the data. Similarly, when a goroutine receives data from a channel, it will block until data is available to receive.

*
*/
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
