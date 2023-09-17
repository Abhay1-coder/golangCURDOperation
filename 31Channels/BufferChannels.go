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

/**
Buffered channels in Go are a variation of regular channels that allow you to send and receive a specified number
of values without blocking. Unlike unbuffered channels, which block until both a sender and a receiver are ready,
buffered channels provide a buffer (a queue) to hold a limited number of values. Here's a simplified explanation:

Buffered Queue: A buffered channel has a queue that can hold a fixed number of values. Think of it as a waiting
room with a certain number of seats.

Send Without Blocking: When you send a value on a buffered channel, it goes into the queue if there is space available.
If the queue is full, the send operation blocks until there's room.

Receive Without Blocking: When you receive a value from a buffered channel, you can do so as long as there's a value in
the queue. If the queue is empty, the receive operation blocks until there's something to receive.

Buffered channels are useful in scenarios where you want to decouple the timing of send and receive operations or
 when you want to control how many values can be buffered before blocking occurs. They can help balance the workloads
 between sender and receiver goroutines and improve concurrency in your programs.
**/
