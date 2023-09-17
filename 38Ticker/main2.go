package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Use a goroutine to update and display the time
	go func() {
		for {
			select {
			case <-ticker.C:
				// Get the current time
				currentTime := time.Now()
				// Display the time in a specific format
				fmt.Printf("Current time: %s\n", currentTime.Format("2006-01-02 15:04:05"))
			}
		}
	}()

	// Let the program run for 5 seconds to observe the periodic updates
	time.Sleep(5 * time.Second)

	// Stop the ticker when done
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
