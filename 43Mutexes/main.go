/**
In Go, a mutex (short for "mutual exclusion") is like a lock that helps protect shared resources, variables, or critical sections of code from being accessed or modified simultaneously by multiple goroutines. It ensures that only one goroutine can access the protected area at any given time, preventing conflicts and data races.

Think of a mutex as a key to a room that contains a valuable item. Only one person can hold the key and enter the room at a time. When they are done, they give the key back, allowing someone else to access the room. This way, everyone can use the valuable item without interfering with each other.

In Go, you use a mutex to protect shared data like this:

1. A goroutine that wants to access the protected data requests the mutex lock (the key) using `mutex.Lock()`. If no other goroutine holds the lock, it gets the lock and can access the data.

2. If another goroutine already holds the lock, the requesting goroutine waits until the lock is released by the first goroutine.

3. After finishing the critical section (the protected code or data), the goroutine releases the lock using `mutex.Unlock()`, allowing other goroutines to access the protected area.

Mutexes ensure that concurrent access to shared resources is coordinated and orderly, preventing conflicts and ensuring data integrity. They are essential for safe concurrent programming in Go.
**/

//Example : I understand your analogy, and I'll provide a Go example that simulates two people trying to eat an apple,
// with a mutex acting as the "key" to access the apple. Only one person can eat the apple at a time:

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex // Mutex acts as the key to access the apple
	var appleAvailable = true

	// Person 1
	go func() {
		for {
			mutex.Lock() // Person 1 tries to get the key (lock) //gate ka key liya
			if appleAvailable {
				fmt.Println("Person 1 is eating the apple.")
				appleAvailable = false      // The apple is now being eaten
				mutex.Unlock()              // Person 1 releases the key (unlock) //gate ko khola
				time.Sleep(2 * time.Second) //2 second apple khaya
				mutex.Lock()                // Person 1 tries to get the key (lock) to put the key back (unlock) //gate ko wapis se lock kiya
				appleAvailable = true       // Person 1 has finished eating, so the apple is available again
				mutex.Unlock()              // Person 1 puts the key back (unlock) //key ko wapas rakh diya
				fmt.Println("Person 1 has finished eating.")
			} else {
				mutex.Unlock() // The apple is not available, so Person 1 puts the key back (unlock)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Person 2
	go func() {
		for {
			mutex.Lock() // Person 2 tries to get the key (lock)
			if appleAvailable {
				fmt.Println("Person 2 is eating the apple.")
				appleAvailable = false // The apple is now being eaten
				mutex.Unlock()         // Person 2 releases the key (unlock)
				time.Sleep(2 * time.Second)
				mutex.Lock()          // Person 2 tries to get the key (lock) to put the key back (unlock)
				appleAvailable = true // Person 2 has finished eating, so the apple is available again
				mutex.Unlock()        // Person 2 puts the key back (unlock)
				fmt.Println("Person 2 has finished eating.")
			} else {
				mutex.Unlock() // The apple is not available, so Person 2 puts the key back (unlock)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Keep the program running
	select {}
}

/**In this code:

We use a sync.Mutex named mutex to represent the "key" for accessing the apple. When a person wants to eat, they lock the mutex (get the key), and when they're done, they unlock the mutex (return the key).

Each person (Person 1 and Person 2) attempts to eat the apple in a loop.

If the apple is available (appleAvailable is true), the person eats it, sets appleAvailable to false, and then releases the mutex (returns the key). After eating, they put the key back by locking the mutex, marking the apple as available, and then unlocking it again.

If the apple is not available, the person puts the key back (unlocks the mutex) and waits for a short time before trying again.

This example demonstrates how a mutex can ensure exclusive access to a shared resource, like the apple in this case, and prevent two people from eating it simultaneously.**/
