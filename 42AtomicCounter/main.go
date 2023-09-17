/**
In Go, an atomic counter is a special kind of counter that ensures safe and synchronized access to a shared variable, especially in concurrent or parallel programming. It guarantees that operations on the counter, such as incrementing or decrementing, happen in an orderly and predictable manner, without conflicts or data races.

Let's break it down in simpler terms:

Imagine you have a scoreboard for a game that multiple players can update simultaneously. Each player might want to increase their score by 1 point whenever they make a move. Without an atomic counter, if two players try to update their scores at the exact same time, there could be a mix-up, and one player's score might not be incremented correctly.

An atomic counter is like a special mechanism that ensures that even if multiple players try to update their scores at the same time, there won't be any mix-up or conflict. It's as if each player takes a number, waits their turn, updates their score, and then goes back to playing. This ensures that the scores are always accurate and don't get messed up due to simultaneous updates.

In Go, atomic counters are used to coordinate and control access to shared data in a way that avoids problems like race conditions. They are an essential tool for safe concurrent programming, ensuring that multiple goroutines can work together without stepping on each other's toes when accessing shared variables.
**/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var player1Score int64 // Player 1's score
	var player2Score int64 // Player 2's score

	// Create a WaitGroup to wait for both players to finish
	var wg sync.WaitGroup

	// Function to simulate a player updating their score
	updateScore := func(playerName string, scorePtr *int64) {
		defer wg.Done()

		for i := 0; i < 5; i++ { // Simulate 5 updates per player
			// Simulate some time passing between score updates
			time.Sleep(time.Millisecond * 100)

			// Atomic operation to increment the score
			atomic.AddInt64(scorePtr, 1)
			fmt.Printf("%s scores a point! Total: %d\n", playerName, atomic.LoadInt64(scorePtr))
		}
	}

	// Add 2 to the WaitGroup (one for each player)
	wg.Add(2)

	// Start two goroutines, one for each player
	go updateScore("Player 1", &player1Score)
	go updateScore("Player 2", &player2Score)

	// Wait for both players to finish
	wg.Wait()

	// Print the final scores
	fmt.Printf("Final Scores - Player 1: %d, Player 2: %d\n", atomic.LoadInt64(&player1Score), atomic.LoadInt64(&player2Score))
}

/**
In this code:

We have two players, Player 1 and Player 2, each with their own score represented by player1Score and player2Score, respectively. These scores are of type int64 to allow atomic operations.

We define an updateScore function that simulates a player updating their score. Each player "scores a point" five times with a slight delay between updates.

We use the sync.WaitGroup (wg) to wait for both players to finish updating their scores.

Inside the main function, we start two goroutines, one for each player, using go. These goroutines simulate the players updating their scores concurrently.

We wait for both players to finish their updates using wg.Wait().

Finally, we print the final scores of both players.

With the use of atomic operations (atomic.AddInt64 and atomic.LoadInt64), we ensure that the scores of both players are updated safely and concurrently without data races.
**/
