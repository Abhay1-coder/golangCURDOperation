/**
In simple terms, a "WaitGroup" in Go is like a counter that helps you wait for a group of tasks (usually goroutines)
to finish their work before proceeding with other parts of your program. Think of it as a way to synchronize multiple
workers.
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Number of goroutines to wait for
	numGoroutines := 7

	// Add goroutines to the wait group
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done() // Signal that the goroutine is done
			fmt.Printf("Goroutine %d is working...\n", id)

			// Simulate some work
			time.Sleep(5 * time.Second)

			fmt.Printf("Goroutine %d is done!\n", id)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines have finished.")
}

/**
We import the necessary packages, including "sync" for the WaitGroup.

We create a sync.WaitGroup named wg to help us wait for goroutines to finish.

We specify the number of goroutines we want to wait for (numGoroutines), which is set to 3.

Inside a loop, we add each goroutine to the wait group using wg.Add(1).

Each goroutine simulates some work with a sleep (2 seconds) and then calls wg.Done() to signal that it has finished its work.

Finally, we call wg.Wait() to block the main goroutine until all added goroutines have signaled that they are done.

When you run this program, you'll see that the main goroutine waits for all three goroutines to complete their work before printing
"All goroutines have finished." The sync.WaitGroup ensures that the main goroutine doesn't exit prematurely and waits for the concurrent goroutines to finish.
**/

//UNDERSTAND HOW THIS CODE WORK INTERANLLY

/**
Certainly! Let's break down how the provided code works internally, step by step:

1. **Importing Packages:** The program begins by importing necessary packages, including `"fmt"` for printing, `"sync"` for synchronization, and `"time"` for time-related functions.

2. **Initializing a WaitGroup:** The program creates a `sync.WaitGroup` named `wg`. This wait group will help the main goroutine wait for other goroutines to finish their work before proceeding.

3. **Defining the Number of Goroutines:** The variable `numGoroutines` is set to 3, indicating that we want to wait for three goroutines to complete their tasks.

4. **Adding Goroutines to the Wait Group:** In a loop that runs `numGoroutines` times, goroutines are added to the wait group using `wg.Add(1)`. This tells the wait group to expect one more goroutine to signal its completion before the main goroutine can proceed.

5. **Starting Goroutines:** Inside the loop, a new goroutine is created for each iteration using `go func(id int) { ... }(i)`. Each goroutine has its own unique ID (`id`) passed as a parameter.

6. **Goroutine Execution:** Each goroutine executes the following steps:
   - It immediately defers a call to `wg.Done()`. This is important because it ensures that the wait group's counter is decremented when the goroutine finishes.
   - The goroutine prints a message indicating that it's working, including its unique ID (`"Goroutine X is working..."`).
   - It simulates some work by sleeping for 2 seconds using `time.Sleep(2 * time.Second)`.
   - After sleeping, it prints a message indicating that it's done (`"Goroutine X is done!"`).

7. **Waiting for Goroutines:** After starting all the goroutines, the main goroutine reaches `wg.Wait()`. This call blocks the main goroutine until the wait group's counter (which was incremented for each goroutine and decremented when each goroutine called `wg.Done()`) reaches zero. In other words, it waits until all the expected goroutines have signaled that they are done.

8. **Final Output:** Once all the goroutines have finished their work and signaled their completion, the main goroutine proceeds to print "All goroutines have finished."

In summary, the code creates a wait group to coordinate the execution of multiple goroutines. Each goroutine does some work (simulated by sleeping) and signals its completion to the wait group. The main goroutine waits for all goroutines to finish using `wg.Wait()` before printing the final message. This demonstrates a basic example of concurrent execution and synchronization in Go using `sync.WaitGroup`.
**/
