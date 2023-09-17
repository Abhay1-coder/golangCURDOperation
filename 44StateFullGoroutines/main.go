/**
In Go, a "stateful" goroutine is one that retains and manages its internal state between multiple function calls or invocations. It means that the goroutine can remember information or data across different moments in its execution.

Imagine you have a game character in a video game. This character has attributes like health, experience points, and inventory. A stateful goroutine is like this game character because it can keep track of its health, experience, and items over time. It remembers its condition from one moment to the next.

In Go, a stateful goroutine might use variables or data structures to store and update its state. For example, it could maintain a counter that increments with each function call, or it could remember information about ongoing tasks.

Stateful goroutines are valuable when you need to maintain context or data across multiple function calls or when you want to track the progress of a long-running operation. They can remember and modify their own internal data while running concurrently with other goroutines.
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Create a channel to communicate with the stateful goroutine
	statefulChan := make(chan int)

	// Start the stateful goroutine
	go statefulGoroutine(statefulChan, &wg)

	// Interact with the stateful goroutine
	for i := 1; i <= 5; i++ {
		statefulChan <- i        // Send a value to the goroutine
		result := <-statefulChan // Receive the updated value from the goroutine
		fmt.Printf("Updated counter: %d\n", result)
	}

	// Signal the stateful goroutine to exit
	close(statefulChan)

	// Wait for the stateful goroutine to finish
	wg.Wait()
	fmt.Println("Stateful goroutine has finished.")
}

func statefulGoroutine(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var counter int

	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Stateful goroutine is exiting.")
				return
			}
			counter += val
			ch <- counter // Send the updated counter value back
		}
	}
}

/**
We create a stateful goroutine statefulGoroutine that communicates with the main function using a channel statefulChan. This goroutine maintains a counter and increments it with each received value from the channel.

In the main function, we start the stateful goroutine and interact with it by sending values to statefulChan and receiving the updated counter value in response.

We send values 1 to 5 to the goroutine and print the updated counter value after each interaction.

To signal the stateful goroutine to exit gracefully, we close the channel (close(statefulChan)).

We use a sync.WaitGroup (wg) to wait for the stateful goroutine to finish before printing a message indicating its completion.

This example demonstrates how a stateful goroutine can maintain and update its internal state (the counter) across multiple interactions with the main function. The stateful goroutine remembers the counter's value from one invocation to the next.
**/
